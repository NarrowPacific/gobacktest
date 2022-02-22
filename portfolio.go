package gobacktest

import (
	"math"
	"time"
)

// PortfolioHandler is the combined interface building block for a portfolio.
type PortfolioHandler interface {
	OnSignaler
	OnFiller
	Investor
	Updater
	Casher
	Valuer
	Reseter
}

// OnSignaler is an interface for the OnSignal method
type OnSignaler interface {
	OnSignal(SignalEvent, DataHandler) (*Order, error)
}

// OnFiller is an interface for the OnFill method
type OnFiller interface {
	OnFill(FillEvent, DataHandler, *ExecutionHandler) (*Fill, *SettledTrade, error)
}

// Investor is an interface to check if a portfolio has a position of a symbol
type Investor interface {
	IsInvested(string) (Position, bool)
	IsLong(string) (Position, bool)
	IsShort(string) (Position, bool)
}

// Updater handles the updating of the portfolio on data events
type Updater interface {
	Update(DataEvent)
}

// Casher handles basic portolio info
type Casher interface {
	InitialCash() float64
	SetInitialCash(float64)
	Cash() float64
	SetCash(float64)
}

// Valuer returns the values of the portfolio
type Valuer interface {
	Value() float64
}

// Booker defines methods for handling the order book of the portfolio
type Booker interface {
	OrderBook() ([]OrderEvent, bool)
	OrdersBySymbol(symbol string) ([]OrderEvent, bool)
}

// Portfolio represent a simple portfolio struct.
type Portfolio struct {
	initialCash  float64
	cash         float64
	holdings     map[string]Position
	orderBook    []OrderEvent
	transactions []FillEvent
	sizeManager  SizeHandler
	riskManager  RiskHandler
}

// NewPortfolio creates a default portfolio with sensible defaults ready for use.
func NewPortfolio() *Portfolio {
	return &Portfolio{
		initialCash: 100000,
		sizeManager: &Size{DefaultSize: 100, DefaultValue: 1000},
		riskManager: &Risk{},
	}
}

// SizeManager return the size manager of the portfolio.
func (p Portfolio) SizeManager() SizeHandler {
	return p.sizeManager
}

// SetSizeManager sets the size manager to be used with the portfolio.
func (p *Portfolio) SetSizeManager(size SizeHandler) {
	p.sizeManager = size
}

// RiskManager returns the risk manager of the portfolio.
func (p Portfolio) RiskManager() RiskHandler {
	return p.riskManager
}

// SetRiskManager sets the risk manager to be used with the portfolio.
func (p *Portfolio) SetRiskManager(risk RiskHandler) {
	p.riskManager = risk
}

// Reset the portfolio into a clean state with set initial cash.
func (p *Portfolio) Reset() error {
	p.cash = 0
	p.holdings = nil
	p.transactions = nil
	return nil
}

// OnSignal handles an incomming signal event
func (p *Portfolio) OnSignal(signal SignalEvent, data DataHandler) (*Order, error) {
	// fmt.Printf("Portfolio receives Signal: %#v \n", signal)

	initialOrder := initializeOrder(signal.Time(), signal.Symbol(), signal.Direction())

	// fetch latest known price for the symbol
	latest := data.Latest(signal.Symbol())

	sizedOrder, err := p.sizeManager.SizeOrder(initialOrder, latest, p)
	if err != nil {
	}

	order, err := p.riskManager.EvaluateOrder(sizedOrder, latest, p.holdings)
	if err != nil {
	}

	return order, nil
}

func initializeOrder(time time.Time, symbol string, direction Direction) *Order {
	// set order type
	orderType := MarketOrder // default Market, should be set by risk manager
	var limit float64

	return &Order{
		Event: Event{
			timestamp: time,
			symbol:    symbol,
		},
		direction: direction,
		// Qty should be set by PositionSizer
		orderType:  orderType,
		limitPrice: limit,
	}
}

// OnFill handles an incomming fill event
func (p *Portfolio) OnFill(fill FillEvent, data DataHandler, exchange *ExecutionHandler) (*Fill, *SettledTrade, error) {
	// Check for nil map, else initialise the map
	if p.holdings == nil {
		p.holdings = make(map[string]Position)
	}

	var f *Fill
	var settledTrade, exitSettledTrade *SettledTrade
	var err error
	// Check if this fill is reverting current position then exit current position first
	if _, ok := p.holdings[fill.Symbol()]; ok {
		_, isShortPos := p.IsShort(fill.Symbol())
		_, isLongPos := p.IsLong(fill.Symbol())
		if (fill.Signal() == ENTRY_BOT && isShortPos) || (fill.Signal() == ENTRY_SLD && isLongPos) {
			if exchange != nil {
				_, exitSettledTrade, err = p.exitPosition(fill.Symbol(), data, *exchange)
				if err != nil {
					return nil, nil, err
				}
			}
		}
	}

	f, settledTrade, err = p.onFill(fill, data)

	// return either onFill or exit Position if present
	if settledTrade == nil {
		settledTrade = exitSettledTrade
	}

	return f, settledTrade, err
}

// OnFill handles an incomming fill event
func (p *Portfolio) onFill(fill FillEvent, data DataHandler) (*Fill, *SettledTrade, error) {
	// Check for nil map, else initialise the map
	if p.holdings == nil {
		p.holdings = make(map[string]Position)
	}

	// check if portfolio has already a holding of the symbol from this fill
	if pos, ok := p.holdings[fill.Symbol()]; ok {
		// update existing Position
		pos.Update(fill)
		p.holdings[fill.Symbol()] = pos
	} else {
		// create new position
		pos := Position{}
		pos.Create(fill)
		p.holdings[fill.Symbol()] = pos
	}

	// update cash
	if fill.Direction() == BOT {
		p.cash = p.cash - fill.NetValue()
	} else {
		// direction is "SLD"
		p.cash = p.cash + fill.NetValue()
	}

	// add fill to transactions
	p.transactions = append(p.transactions, fill)

	f := fill.(*Fill)

	// Check if settled trade
	var s *SettledTrade
	if pos, ok := p.holdings[fill.Symbol()]; ok {
		if pos.qty == 0 {
			s = &SettledTrade{
				Orders:          pos.transactions,
				Qty:             pos.qtyBOT,
				Profit:          pos.Profit(),
				ProfitPercent:   pos.ProfitPercent(),
				DrawDown:        pos.Drawdown(),
				DrawDownPercent: pos.DrawdownPercent(),
				RunUp:           pos.RunUp(),
				RunUpPercent:    pos.RunUpPercent(),
			}

			// Remove position which is settled
			delete(p.holdings, fill.Symbol())
		}
	}

	return f, s, nil
}

func (p *Portfolio) exitPosition(symbol string, data DataHandler, exchange ExecutionHandler) (*Fill, *SettledTrade, error) {
	if pos, ok := p.holdings[symbol]; ok {
		order := initializeOrder(data.Latest(symbol).Time(), symbol, pos.direction.GetOpposite())
		order.SetQty(int64(math.Abs(float64(pos.qty))))
		order.SetSignal(pos.direction.GetOpposite())
		exitFill, err := exchange.OnOrder(order, data)
		if err != nil {
			return nil, nil, err
		}

		return p.onFill(exitFill, data)
	}

	return nil, nil, nil
}

// IsInvested checks if the portfolio has an open position on the given symbol
func (p Portfolio) IsInvested(symbol string) (pos Position, ok bool) {
	pos, ok = p.holdings[symbol]
	if ok && (pos.qty != 0) {
		return pos, true
	}
	return pos, false
}

// IsLong checks if the portfolio has an open long position on the given symbol
func (p Portfolio) IsLong(symbol string) (pos Position, ok bool) {
	pos, ok = p.holdings[symbol]
	if ok && (pos.qty > 0) {
		return pos, true
	}
	return pos, false
}

// IsShort checks if the portfolio has an open short position on the given symbol
func (p Portfolio) IsShort(symbol string) (pos Position, ok bool) {
	pos, ok = p.holdings[symbol]
	if ok && (pos.qty < 0) {
		return pos, true
	}
	return pos, false
}

// Update updates the holding on a data event
func (p *Portfolio) Update(d DataEvent) {
	if pos, ok := p.IsInvested(d.Symbol()); ok {
		pos.UpdateValue(d)
		p.holdings[d.Symbol()] = pos
	}
}

// SetInitialCash sets the initial cash value of the portfolio
func (p *Portfolio) SetInitialCash(initial float64) {
	p.initialCash = initial
}

// InitialCash returns the initial cash value of the portfolio
func (p Portfolio) InitialCash() float64 {
	return p.initialCash
}

// SetCash sets the current cash value of the portfolio
func (p *Portfolio) SetCash(cash float64) {
	p.cash = cash
}

// Cash returns the current cash value of the portfolio
func (p Portfolio) Cash() float64 {
	return p.cash
}

// Value return the current total value of the portfolio
func (p Portfolio) Value() float64 {
	var holdingValue float64
	for _, pos := range p.holdings {

		holdingValue += pos.marketValue
	}

	value := p.cash + holdingValue
	return value
}

// Holdings returns the holdings of the portfolio
func (p Portfolio) Holdings() map[string]Position {
	return p.holdings
}

// OrderBook returns the order book of the portfolio
func (p Portfolio) OrderBook() ([]OrderEvent, bool) {
	if len(p.orderBook) == 0 {
		return p.orderBook, false
	}

	return p.orderBook, true
}

// OrdersBySymbol returns the order of a specific symbol from the order book.
func (p Portfolio) OrdersBySymbol(symbol string) ([]OrderEvent, bool) {
	var orders = []OrderEvent{}

	for _, order := range p.orderBook {
		if order.Symbol() == symbol {
			orders = append(orders, order)
		}
	}

	if len(orders) == 0 {
		return orders, false
	}

	return orders, true
}
