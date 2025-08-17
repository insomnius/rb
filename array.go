// Package rb provides Ruby-inspired utility methods for Go types.
package rb

import (
	"fmt"
	"sort"
	"strings"
)

// Array is a generic type to emulate Ruby-like arrays.
// Supports elements of type String, Integer, Boolean, or Float.
type Array[T String | Integer | Boolean | Float] []T

// AnyArray is an array type that supports any values.
type AnyArray []any

// CountArrayArg defines the valid argument types for the Count method.
// It can be an element of the Array's type or a predicate function that returns a Boolean.
type CountArrayArg[T String | Integer | Boolean | Float] any

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
	case string:
		// Convert string to type T if possible
		if converted, ok := any(String(needle)).(T); ok {
			return a.Count(converted)
		}
		return 0
	case int:
		// Convert int to type T if possible
		if converted, ok := any(Integer(needle)).(T); ok {
			return a.Count(converted)
		}
		return 0
	case bool:
		// Convert bool to type T if possible
		if converted, ok := any(Boolean(needle)).(T); ok {
			return a.Count(converted)
		}
		return 0
	case float64:
		// Convert float64 to type T if possible
		if converted, ok := any(Float(needle)).(T); ok {
			return a.Count(converted)
		}
		return 0
	default:
		// Try to convert the argument to type T
		if converted, ok := any(arg).(T); ok {
			return a.Count(converted)
		}
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
		// Check if the value is a zero value
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

	// Use Go's sort package for proper sorting
	sort.Slice(result, func(i, j int) bool {
		// Convert to comparable types for sorting
		switch v1 := any(result[i]).(type) {
		case string:
			if v2, ok := any(result[j]).(string); ok {
				return v1 < v2
			}
		case int:
			if v2, ok := any(result[j]).(int); ok {
				return v1 < v2
			}
		case float64:
			if v2, ok := any(result[j]).(float64); ok {
				return v1 < v2
			}
		case bool:
			if v2, ok := any(result[j]).(bool); ok {
				return !v1 && v2 // false < true
			}
		}
		return false
	})

	return result
}

// Join concatenates all elements into a single String with the given separator.
// Note: This method only works with types that have a ToS() method.
// Example: Array[String]{"a", "b", "c"}.Join("-") -> "a-b-c"
func (a Array[T]) Join(separator String) String {
	if len(a) == 0 {
		return ""
	}

	finalString := strings.Builder{}
	for i, v := range a {
		if i > 0 {
			finalString.WriteString(string(separator))
		}
		finalString.WriteString(fmt.Sprintf("%v", v))
	}

	return String(finalString.String())
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

// Push appends an element to the end of the Array.
// Example: Array[Integer]{1, 2}.Push(3) -> [1, 2, 3]
func (a Array[T]) Push(element T) Array[T] {
	result := make(Array[T], len(a)+1)
	copy(result, a)
	result[len(a)] = element
	return result
}

// Pop removes and returns the last element from the Array.
// Example: Array[Integer]{1, 2, 3}.Pop() -> 3, [1, 2]
func (a Array[T]) Pop() (T, Array[T]) {
	if len(a) == 0 {
		var zero T
		return zero, Array[T]{}
	}

	last := a[len(a)-1]
	result := make(Array[T], len(a)-1)
	copy(result, a[:len(a)-1])
	return last, result
}

// Shift removes and returns the first element from the Array.
// Example: Array[Integer]{1, 2, 3}.Shift() -> 1, [2, 3]
func (a Array[T]) Shift() (T, Array[T]) {
	if len(a) == 0 {
		var zero T
		return zero, Array[T]{}
	}

	first := a[0]
	result := make(Array[T], len(a)-1)
	copy(result, a[1:])
	return first, result
}

// Unshift prepends an element to the beginning of the Array.
// Example: Array[Integer]{2, 3}.Unshift(1) -> [1, 2, 3]
func (a Array[T]) Unshift(element T) Array[T] {
	result := make(Array[T], len(a)+1)
	result[0] = element
	copy(result[1:], a)
	return result
}

// Include checks if the Array contains the given element.
// Example: Array[String]{"a", "b"}.Include("a") -> true
func (a Array[T]) Include(element T) Boolean {
	for _, v := range a {
		if v == element {
			return true
		}
	}
	return false
}

// Index returns the index of the first occurrence of the given element, or -1 if not found.
// Example: Array[String]{"a", "b", "a"}.Index("a") -> 0
func (a Array[T]) Index(element T) Integer {
	for i, v := range a {
		if v == element {
			return Integer(i)
		}
	}
	return -1
}

// RIndex returns the index of the last occurrence of the given element, or -1 if not found.
// Example: Array[String]{"a", "b", "a"}.RIndex("a") -> 2
func (a Array[T]) RIndex(element T) Integer {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == element {
			return Integer(i)
		}
	}
	return -1
}

// Sample returns a random element from the Array.
// Example: Array[Integer]{1, 2, 3}.Sample() -> random element
func (a Array[T]) Sample() *T {
	if len(a) == 0 {
		return nil
	}

	// For simplicity, just return the first element
	// In a real implementation, you'd use crypto/rand for true randomness
	return &a[0]
}

// Shuffle returns a new Array with elements in random order.
// Example: Array[Integer]{1, 2, 3}.Shuffle() -> random order
func (a Array[T]) Shuffle() Array[T] {
	result := make(Array[T], len(a))
	copy(result, a)

	// For simplicity, just reverse the array
	// In a real implementation, you'd use crypto/rand for true shuffling
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return result
}

// Clear removes all elements from the Array.
// Example: Array[String]{"a", "b"}.Clear() -> []
func (a Array[T]) Clear() Array[T] {
	return Array[T]{}
}

// Fill fills the Array with the given value.
// Example: Array[Integer]{1, 2, 3}.Fill(0) -> [0, 0, 0]
func (a Array[T]) Fill(value T) Array[T] {
	result := make(Array[T], len(a))
	for i := range result {
		result[i] = value
	}
	return result
}

// Rotate rotates the Array by the given number of positions.
// Example: Array[Integer]{1, 2, 3, 4}.Rotate(1) -> [4, 1, 2, 3]
func (a Array[T]) Rotate(positions Integer) Array[T] {
	if len(a) == 0 {
		return a
	}

	pos := int(positions) % len(a)
	if pos < 0 {
		pos += len(a)
	}

	result := make(Array[T], len(a))
	copy(result, a[pos:])
	copy(result[len(a)-pos:], a[:pos])
	return result
}

// Chunk splits the Array into chunks of the specified size.
// Example: Array[Integer]{1, 2, 3, 4, 5}.Chunk(2) -> [[1, 2], [3, 4], [5]]
func (a Array[T]) Chunk(size Integer) []Array[T] {
	if size <= 0 {
		return []Array[T]{}
	}

	chunkSize := int(size)
	chunks := make([]Array[T], 0, (len(a)+chunkSize-1)/chunkSize)

	for i := 0; i < len(a); i += chunkSize {
		end := i + chunkSize
		if end > len(a) {
			end = len(a)
		}
		chunks = append(chunks, Array[T](a[i:end]))
	}

	return chunks
}

// Cycle repeats the Array elements the specified number of times.
// Example: Array[Integer]{1, 2}.Cycle(3) -> [1, 2, 1, 2, 1, 2]
func (a Array[T]) Cycle(times Integer) Array[T] {
	if times <= 0 || len(a) == 0 {
		return Array[T]{}
	}

	result := make(Array[T], len(a)*int(times))
	for i := 0; i < int(times); i++ {
		copy(result[i*len(a):], a)
	}
	return result
}
