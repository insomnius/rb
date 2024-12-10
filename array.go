package rb

type Array[T String | Integer] []T
type CountArrayArg[T String | Integer] any

func (a Array[T]) Count(args ...CountArrayArg[T]) Integer {
	if len(args) == 0 {
		return Integer(len(a))
	}

	arg := args[0]
	switch needle := arg.(type) {
	case nil:
		return Integer(len(a))
	case string:
		return a.Count(String(needle))
	case int:
		return a.Count(Integer(needle))
	case T:
		tot := 0

		for _, v := range a {
			if v == needle {
				tot++
			}
		}

		return Integer(tot)
	case func(T) bool:
		tot := 0
		for _, v := range a {
			if needle(v) {
				tot++
			}
		}
		return Integer(tot)
	default:
		return 0
	}
}
