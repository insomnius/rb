package rb

import (
	"fmt"
)

// Range represents a range of values, similar to Ruby's Range class.
type Range[T Integer | Float] struct {
	Begin T
	End   T
	Exclusive bool // true for exclusive range (..), false for inclusive range (...)
}

// NewRange creates a new inclusive Range.
// Example: NewRange(Integer(1), Integer(5)) -> 1..5
func NewRange[T Integer | Float](begin, end T) Range[T] {
	return Range[T]{Begin: begin, End: end, Exclusive: false}
}

// NewExclusiveRange creates a new exclusive Range.
// Example: NewExclusiveRange(Integer(1), Integer(5)) -> 1...5
func NewExclusiveRange[T Integer | Float](begin, end T) Range[T] {
	return Range[T]{Begin: begin, End: end, Exclusive: true}
}

// Each executes the given function for each value in the Range.
// Example: NewRange(Integer(1), Integer(3)).Each(func(i Integer) { fmt.Println(i) })
func (r Range[T]) Each(fn func(T)) {
	if r.Begin <= r.End {
		for i := r.Begin; i <= r.End; i++ {
			if r.Exclusive && i == r.End {
				break
			}
			fn(i)
		}
	} else {
		for i := r.Begin; i >= r.End; i-- {
			if r.Exclusive && i == r.End {
				break
			}
			fn(i)
		}
	}
}

// EachWithIndex executes the given function for each value in the Range with its index.
// Example: NewRange(Integer(1), Integer(3)).EachWithIndex(func(i, idx Integer) { fmt.Printf("%d: %d\n", idx, i) })
func (r Range[T]) EachWithIndex(fn func(T, Integer)) {
	idx := Integer(0)
	if r.Begin <= r.End {
		for i := r.Begin; i <= r.End; i++ {
			if r.Exclusive && i == r.End {
				break
			}
			fn(i, idx)
			idx++
		}
	} else {
		for i := r.Begin; i >= r.End; i-- {
			if r.Exclusive && i == r.End {
				break
			}
			fn(i, idx)
			idx++
		}
	}
}

// Include checks if the given value is included in the Range.
// Example: NewRange(Integer(1), Integer(5)).Include(Integer(3)) -> true
func (r Range[T]) Include(value T) Boolean {
	if r.Begin <= r.End {
		if r.Exclusive {
			return Boolean(value >= r.Begin && value < r.End)
		}
		return Boolean(value >= r.Begin && value <= r.End)
	} else {
		if r.Exclusive {
			return Boolean(value <= r.Begin && value > r.End)
		}
		return Boolean(value <= r.Begin && value >= r.End)
	}
}

// Cover checks if the given value is covered by the Range (alias for Include).
// Example: NewRange(Integer(1), Integer(5)).Cover(Integer(3)) -> true
func (r Range[T]) Cover(value T) Boolean {
	return r.Include(value)
}

// Step executes the given function for each value in the Range, incrementing by the given step.
// Example: NewRange(Integer(0), Integer(10)).Step(Integer(2), func(i Integer) { fmt.Println(i) })
func (r Range[T]) Step(step T, fn func(T)) {
	if step == 0 {
		return
	}
	
	if r.Begin <= r.End {
		for i := r.Begin; i <= r.End; i += step {
			if r.Exclusive && i >= r.End {
				break
			}
			fn(i)
		}
	} else {
		for i := r.Begin; i >= r.End; i -= step {
			if r.Exclusive && i <= r.End {
				break
			}
			fn(i)
		}
	}
}

// ToArray converts the Range to an Array.
// Example: NewRange(Integer(1), Integer(3)).ToArray() -> [1, 2, 3]
func (r Range[T]) ToArray() Array[T] {
	var result []T
	
	if r.Begin <= r.End {
		for i := r.Begin; i <= r.End; i++ {
			if r.Exclusive && i == r.End {
				break
			}
			result = append(result, i)
		}
	} else {
		for i := r.Begin; i >= r.End; i-- {
			if r.Exclusive && i == r.End {
				break
			}
			result = append(result, i)
		}
	}
	
	return Array[T](result)
}

// Size returns the number of values in the Range.
// Example: NewRange(Integer(1), Integer(5)).Size() -> 5
func (r Range[T]) Size() Integer {
	if r.Begin <= r.End {
		if r.Exclusive {
			return Integer(r.End - r.Begin)
		}
		return Integer(r.End - r.Begin + 1)
	} else {
		if r.Exclusive {
			return Integer(r.Begin - r.End)
		}
		return Integer(r.Begin - r.End + 1)
	}
}

// Length is an alias for Size.
func (r Range[T]) Length() Integer {
	return r.Size()
}

// IsEmpty checks if the Range is empty.
// Example: NewRange(Integer(5), Integer(1)).IsEmpty() -> true
func (r Range[T]) IsEmpty() Boolean {
	if r.Begin <= r.End {
		return Boolean(r.Begin > r.End || (r.Exclusive && r.Begin == r.End))
	}
	return Boolean(r.Begin < r.End || (r.Exclusive && r.Begin == r.End))
}

// Min returns the minimum value in the Range.
// Example: NewRange(Integer(1), Integer(5)).Min() -> 1
func (r Range[T]) Min() T {
	if r.Begin <= r.End {
		return r.Begin
	}
	return r.End
}

// Max returns the maximum value in the Range.
// Example: NewRange(Integer(1), Integer(5)).Max() -> 5
func (r Range[T]) Max() T {
	if r.Begin <= r.End {
		return r.End
	}
	return r.Begin
}

// First returns the first value in the Range.
// Example: NewRange(Integer(1), Integer(5)).First() -> 1
func (r Range[T]) First() T {
	return r.Begin
}

// Last returns the last value in the Range.
// Example: NewRange(Integer(1), Integer(5)).Last() -> 5
func (r Range[T]) Last() T {
	if r.Exclusive {
		if r.Begin <= r.End {
			return r.End - 1
		}
		return r.End + 1
	}
	return r.End
}

// Begin returns the beginning value of the Range.
// Example: NewRange(Integer(1), Integer(5)).Begin() -> 1
func (r Range[T]) BeginValue() T {
	return r.Begin
}

// End returns the end value of the Range.
// Example: NewRange(Integer(1), Integer(5)).End() -> 5
func (r Range[T]) EndValue() T {
	return r.End
}

// IsExclusive checks if the Range is exclusive.
// Example: NewExclusiveRange(Integer(1), Integer(5)).IsExclusive() -> true
func (r Range[T]) IsExclusive() Boolean {
	return Boolean(r.Exclusive)
}

// IsInclusive checks if the Range is inclusive.
// Example: NewRange(Integer(1), Integer(5)).IsInclusive() -> true
func (r Range[T]) IsInclusive() Boolean {
	return Boolean(!r.Exclusive)
}

// Reverse returns a new Range with begin and end values swapped.
// Example: NewRange(Integer(1), Integer(5)).Reverse() -> 5..1
func (r Range[T]) Reverse() Range[T] {
	return Range[T]{
		Begin:     r.End,
		End:       r.Begin,
		Exclusive: r.Exclusive,
	}
}

// Overlap checks if this Range overlaps with another Range.
// Example: NewRange(Integer(1), Integer(5)).Overlap(NewRange(Integer(3), Integer(7))) -> true
func (r Range[T]) Overlap(other Range[T]) Boolean {
	if r.Begin <= r.End && other.Begin <= other.End {
		return Boolean(r.Begin <= other.End && other.Begin <= r.End)
	} else if r.Begin > r.End && other.Begin > other.End {
		return Boolean(r.End <= other.Begin && other.End <= r.Begin)
	}
	return false
}

// Contains checks if this Range completely contains another Range.
// Example: NewRange(Integer(1), Integer(10)).Contains(NewRange(Integer(3), Integer(7))) -> true
func (r Range[T]) Contains(other Range[T]) Boolean {
	if r.Begin <= r.End && other.Begin <= other.End {
		return Boolean(r.Begin <= other.Begin && r.End >= other.End)
	} else if r.Begin > r.End && other.Begin > other.End {
		return Boolean(r.Begin >= other.Begin && r.End <= other.End)
	}
	return false
}

// ToS converts the Range to a String representation.
// Example: NewRange(Integer(1), Integer(5)).ToS() -> "1..5"
func (r Range[T]) ToS() String {
	if r.Exclusive {
		return String(fmt.Sprintf("%v...%v", r.Begin, r.End))
	}
	return String(fmt.Sprintf("%v..%v", r.Begin, r.End))
}

// ToStr is an alias for ToS.
func (r Range[T]) ToStr() String {
	return r.ToS()
}

// Clone returns a copy of the Range.
// Example: NewRange(Integer(1), Integer(5)).Clone()
func (r Range[T]) Clone() Range[T] {
	return Range[T]{
		Begin:     r.Begin,
		End:       r.End,
		Exclusive: r.Exclusive,
	}
}