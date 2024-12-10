package rb

// Array is a generic type to emulate Ruby-like arrays.
// Supports elements of type String, Integer, or Boolean.
type Array[T String | Integer | Boolean] []T

// CountArrayArg defines the valid argument types for the Count method.
// It can be an element of the Array's type or a predicate function that returns a Boolean.
type CountArrayArg[T String | Integer | Boolean] any

// Count returns the count of elements in the Array based on the provided argument:
// - If no arguments are given, it returns the total number of elements in the Array.
// - If an argument of type T is provided, it counts elements equal to the argument.
// - If a function (predicate) is provided, it counts elements for which the function returns true.
//
// Example:
// a := Array[String]{"a", "b", "a"}
// a.Count()                  // -> 3
// a.Count(String("a"))       // -> 2
// a.Count(func(s String) bool { return s == "b" }) // -> 1
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
	case bool:
		return a.Count(Boolean(needle))
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
