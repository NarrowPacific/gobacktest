package gobacktest

// Direction defines which direction a signal indicates
type Direction int

// different types of order directions
const (
	// Buy
	BOT Direction = iota // 0
	// Entry Buy
	ENTRY_BOT
	// Sell
	SLD
	// Entry Sell
	ENTRY_SLD
	// Hold
	HLD
	// Exit
	EXT
)

func (d Direction) String() string {
	switch d {
	case BOT, ENTRY_BOT:
		return "BUY"
	case SLD, ENTRY_SLD:
		return "SELL"
	case HLD:
		return "HOLD"
	case EXT:
		return "EXIT"
	default:
		return "UNKNOWN"
	}
}

func (d Direction) GetOpposite() Direction {
	switch d {
	case BOT:
		return SLD
	case SLD:
		return BOT
	default:
		return d
	}
}

// Signal declares a basic signal event
type Signal struct {
	Event
	direction Direction // long, short, exit or hold
}

// Direction returns the Direction of a Signal
func (s Signal) Direction() Direction {
	return s.direction
}

// SetDirection sets the Directions field of a Signal
func (s *Signal) SetDirection(dir Direction) {
	s.direction = dir
}
