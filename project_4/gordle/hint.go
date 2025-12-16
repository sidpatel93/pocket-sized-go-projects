package gordle

type hint byte

const (
	absentCharacter hint = iota
	wrongPositionCharacter
	correctCharacter
)

// hint implements the Stringer interface.
func (h hint) String() string {
	switch h {
	case absentCharacter:
		return "◻️" // grey square
	case wrongPositionCharacter:
		return "🟡" // yellow circle
	case correctCharacter:
		return "💚" // green heart
	default:
		// This should never happen.
		return "💔" // red broken heart
	}
}
