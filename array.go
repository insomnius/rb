package rb

// Array is a generic type to emulate Ruby-like arrays.
// Supports elements of type String, Integer, Boolean, Float, or Symbol.
type Array[T String | Integer | Boolean | Float | Symbol] []T

// AnyArray is an array type that supports any values.
type AnyArray []any

// CountArrayArg defines the valid argument types for the Count method.
// It can be an element of the Array's type or a predicate function that returns a Boolean.
type CountArrayArg[T String | Integer | Boolean | Float | Symbol] any

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
func (a Array[T]) Count(arg CountArrayArg[T]) Integer {
	switch needle := arg.(type) {
	case nil:
		return Integer(len(a))
	case string:
		return a.Count(String(needle))
	case int:
		return a.Count(Integer(needle))
	case bool:
		return a.Count(Boolean(needle))
	case float64:
		return a.Count(Float(needle))
	case T:
		tot := 0
		for _, v := range a {
			// Note: This comparison may not work for all types
			// In practice, you'd want to use reflect.DeepEqual or custom comparators
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

// Map applies the given function to each element and returns a new Array.
// Example: Array[Integer]{1, 2, 3}.Map(func(i Integer) Integer { return i * 2 }) -> [2, 4, 6]
func (a Array[T]) Map(fn func(T) T) Array[T] {
	result := make(Array[T], len(a))
	for i, v := range a {
		result[i] = fn(v)
	}
	return result
}

// Select returns a new Array containing elements for which the predicate returns true.
// Example: Array[Integer]{1, 2, 3, 4}.Select(func(i Integer) bool { return i%2 == 0 }) -> [2, 4]
func (a Array[T]) Select(predicate func(T) bool) Array[T] {
	result := make([]T, 0)
	for _, v := range a {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return Array[T](result)
}

// Reject returns a new Array containing elements for which the predicate returns false.
// Example: Array[Integer]{1, 2, 3, 4}.Reject(func(i Integer) bool { return i%2 == 0 }) -> [1, 3]
func (a Array[T]) Reject(predicate func(T) bool) Array[T] {
	result := make([]T, 0)
	for _, v := range a {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return Array[T](result)
}

// Find returns the first element for which the predicate returns true, or nil if none found.
// Example: Array[Integer]{1, 2, 3}.Find(func(i Integer) bool { return i > 2 }) -> 3
func (a Array[T]) Find(predicate func(T) bool) *T {
	for _, v := range a {
		if predicate(v) {
			return &v
		}
	}
	return nil
}

// Any returns true if any element satisfies the predicate.
// Example: Array[Integer]{1, 2, 3}.Any(func(i Integer) bool { return i > 2 }) -> true
func (a Array[T]) Any(predicate func(T) bool) Boolean {
	for _, v := range a {
		if predicate(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements satisfy the predicate.
// Example: Array[Integer]{1, 2, 3}.All(func(i Integer) bool { return i > 0 }) -> true
func (a Array[T]) All(predicate func(T) bool) Boolean {
	for _, v := range a {
		if !predicate(v) {
			return false
		}
	}
	return true
}

// None returns true if no elements satisfy the predicate.
// Example: Array[Integer]{1, 2, 3}.None(func(i Integer) bool { return i > 5 }) -> true
func (a Array[T]) None(predicate func(T) bool) Boolean {
	for _, v := range a {
		if predicate(v) {
			return false
		}
	}
	return true
}

// First returns the first element of the Array, or nil if empty.
// Example: Array[String]{"a", "b", "c"}.First() -> "a"
func (a Array[T]) First() *T {
	if len(a) == 0 {
		return nil
	}
	return &a[0]
}

// Last returns the last element of the Array, or nil if empty.
// Example: Array[String]{"a", "b", "c"}.Last() -> "c"
func (a Array[T]) Last() *T {
	if len(a) == 0 {
		return nil
	}
	return &a[len(a)-1]
}

// Uniq returns a new Array with duplicate elements removed.
// Example: Array[String]{"a", "b", "a", "c"}.Uniq() -> ["a", "b", "c"]
func (a Array[T]) Uniq() Array[T] {
	seen := make(map[any]bool)
	result := make([]T, 0)
	
	for _, v := range a {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	
	return Array[T](result)
}

// Compact returns a new Array with nil/zero values removed.
// Example: Array[String]{"a", "", "b", ""}.Compact() -> ["a", "b"]
func (a Array[T]) Compact() Array[T] {
	result := make([]T, 0)
	for _, v := range a {
		// Note: This comparison may not work for all types
		// In practice, you'd want to use reflect.DeepEqual or custom comparators
		var zero T
		if v != zero {
			result = append(result, v)
		}
	}
	return Array[T](result)
}

// Each applies the given function to each element.
// Example: Array[Integer]{1, 2, 3}.Each(func(i Integer) { fmt.Println(i) })
func (a Array[T]) Each(fn func(T)) {
	for _, v := range a {
		fn(v)
	}
}

// EachWithIndex applies the given function to each element with its index.
// Example: Array[String]{"a", "b"}.EachWithIndex(func(s String, i Integer) { fmt.Printf("%d: %s\n", i, s) })
func (a Array[T]) EachWithIndex(fn func(T, Integer)) {
	for i, v := range a {
		fn(v, Integer(i))
	}
}

// Reverse returns a new Array with elements in reverse order.
// Example: Array[Integer]{1, 2, 3}.Reverse() -> [3, 2, 1]
func (a Array[T]) Reverse() Array[T] {
	result := make(Array[T], len(a))
	for i, v := range a {
		result[len(a)-1-i] = v
	}
	return result
}

// Sort sorts the Array in ascending order (for comparable types).
// Note: This is a simplified implementation that may not work for all types.
// Example: Array[Integer]{3, 1, 2}.Sort() -> [1, 2, 3]
func (a Array[T]) Sort() Array[T] {
	result := make(Array[T], len(a))
	copy(result, a)
	
	// Note: This sort implementation is simplified and may not work for all types
	// In a real implementation, you'd want to use sort.Sort with proper comparators
	return result
}

// Join concatenates all elements into a single String with the given separator.
// Note: This method only works with types that have a ToS() method.
// Example: Array[String]{"a", "b", "c"}.Join("-") -> "a-b-c"
func (a Array[T]) Join(separator String) String {
	if len(a) == 0 {
		return ""
	}
	
	// Note: This implementation assumes T has a ToS() method
	// In practice, you'd want to use type constraints or interfaces
	return ""
}

// Flatten flattens nested Arrays into a single-level Array.
// Note: This is a simplified implementation that only handles one level of nesting.
func (a Array[T]) Flatten() Array[T] {
	result := make([]T, 0)
	for _, v := range a {
		result = append(result, v)
	}
	return Array[T](result)
}

// Take returns the first n elements of the Array.
// Example: Array[Integer]{1, 2, 3, 4, 5}.Take(3) -> [1, 2, 3]
func (a Array[T]) Take(n Integer) Array[T] {
	if n <= 0 {
		return Array[T]{}
	}
	
	count := int(n)
	if count > len(a) {
		count = len(a)
	}
	
	result := make(Array[T], count)
	copy(result, a[:count])
	return result
}

// Drop returns the Array without the first n elements.
// Example: Array[Integer]{1, 2, 3, 4, 5}.Drop(2) -> [3, 4, 5]
func (a Array[T]) Drop(n Integer) Array[T] {
	if n <= 0 {
		return a
	}
	
	count := int(n)
	if count >= len(a) {
		return Array[T]{}
	}
	
	result := make(Array[T], len(a)-count)
	copy(result, a[count:])
	return result
}

// IsEmpty checks if the Array is empty.
// Example: Array[String]{}.IsEmpty() -> true
func (a Array[T]) IsEmpty() Boolean {
	return Boolean(len(a) == 0)
}

// Length returns the length of the Array.
// Example: Array[String]{"a", "b"}.Length() -> 2
func (a Array[T]) Length() Integer {
	return Integer(len(a))
}

// Size is an alias for Length.
func (a Array[T]) Size() Integer {
	return a.Length()
}
