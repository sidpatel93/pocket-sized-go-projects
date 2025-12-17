package gordle

import "strings"

// hint represents the status of a character in a guess
type hint byte

// feedback is the list of hints for the given word
type feedback []hint

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

func (f feedback) String() string {
	result := strings.Builder{}
	for _, h := range f {
		result.WriteString(h.String())
	}
	return result.String()
}

// Equal determines equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}
	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}
	return true
}
