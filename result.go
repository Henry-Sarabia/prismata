package prismata

// Result represents the outcome of the match.
type Result int

const (
	// P1 denotes a Player 1 win.
	P1 Result = 0
	// P2 deontes a Player 2 win.
	P2 Result = 1
	// Draw denotes a draw.
	Draw Result = 2
)

func (r Result) String() string {
	switch r {
	case 0:
		return "P1"
	case 1:
		return "P2"
	case 2:
		return "Draw"
	default:
		return "Unknown"
	}
}
