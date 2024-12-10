package rb

type Array[T String | Int] []T
type CountArrayArg[T String | Int] any

func (a Array[T]) Count(args ...CountArrayArg[T]) Int {
	if len(args) == 0 {
		return Int(len(a))
	}

	arg := args[0]
	switch needle := arg.(type) {
	case nil:
		return Int(len(a))
	case string:
		return a.Count(String(needle))
	case int:
		return a.Count(Int(needle))
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
		return 0
	}
}
