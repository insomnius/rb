package rb

// Integer is a custom integer type to emulate Ruby-like behavior.
type Integer int

// IsOdd checks if the Integer is odd.
// Example: Integer(3).IsOdd() -> true
// Example: Integer(4).IsOdd() -> false
func (i Integer) IsOdd() Boolean {
	return i%2 == 1
}

func (i Integer) IsEven() Boolean {
	return i%2 == 0
}
