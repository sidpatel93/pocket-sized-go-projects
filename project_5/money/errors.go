package money

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	// ErrInvalidDecimal is returned if the decimal is malformed.
	ErrInvalidDecimal = Error("unable to convert the decimal")

	// ErrTooLarge is returned if the quantity is too large
	// this would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")
)
