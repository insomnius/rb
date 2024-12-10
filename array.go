package rb

type Array[T comparable] []T

type CountArrayArg[T comparable] interface {
	any | func(T) bool
}

func (a Array[T]) Count(arg CountArrayArg[T]) Int {
	switch needle := arg.(type) {
	case nil:
		return Int(len(a))
	case T:
		tot := 0

		for _, v := range a {
			if v == needle {
				tot++
			}
		}

		return Int(tot)
	case func(T) bool:
		tot := 0
		for _, v := range a {
			if needle(v) {
				tot++
			}
		}
		return Int(tot)
	default:
		return Int(len(a))
	}
}
