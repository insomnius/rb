package rb

type Integer int

func (i Integer) IsOdd() Boolean {
	return i%2 == 1
}
