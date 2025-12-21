package money

import (
	"strconv"
	"strings"
)

// Decimal can represent a floating-point number with a fixed precision.
// example:
// 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Decimal struct {
	// subunits is the amount of subunits.
	// Multiply it by the precision to get the real value
	subunits int64
	// Number of "subunits" in a unit, expressed as a power of 10.
	precision byte
}

const maxDecimal = 1e12

func ParseDecimmal(s string) (Decimal, error) {
	// split string on the decimal point
	before, after, _ := strings.Cut(s, ".")
	subunit, err := strconv.ParseInt(before+after, 10, 64)
	if err != nil {
		return Decimal{}, ErrInvalidDecimal
	}
	if subunit > maxDecimal {
		return Decimal{}, ErrTooLarge
	}
	precision := byte(len(after))
	return Decimal{
		subunits:  subunit,
		precision: precision,
	}, nil
}
