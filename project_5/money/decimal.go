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

func ParseDecimmal(s string) (Decimal, error) {
	// split string on the decimal point
	before, after, found := strings.Cut(s, ".")
	var subunits int64
	if !found {
		// parse before as integer
		subunits, _ = strconv.ParseInt(before, 10, 64)
	}
	subunits, err := strconv.ParseInt(before+after, 10, 64)
	if err != nil {
		return Decimal{}, err
	}
	precision := byte(len(after))
	return Decimal{
		subunits:  subunits,
		precision: precision,
	}, nil
}
