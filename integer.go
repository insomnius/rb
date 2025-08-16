package rb

import (
	"fmt"
	"math"
)

// Integer is a custom integer type to emulate Ruby-like behavior.
type Integer int

// IsOdd checks if the Integer is odd.
// Example: Integer(3).IsOdd() -> true
// Example: Integer(4).IsOdd() -> false
func (i Integer) IsOdd() Boolean {
	return !i.IsEven()
}

func (i Integer) IsEven() Boolean {
	return Boolean(i%2 == 0)
}

// IsPositive checks if the Integer is positive (greater than 0).
// Example: Integer(5).IsPositive() -> true
func (i Integer) IsPositive() Boolean {
	return Boolean(i > 0)
}

// IsNegative checks if the Integer is negative (less than 0).
// Example: Integer(-5).IsNegative() -> true
func (i Integer) IsNegative() Boolean {
	return Boolean(i < 0)
}

// IsZero checks if the Integer is zero.
// Example: Integer(0).IsZero() -> true
func (i Integer) IsZero() Boolean {
	return Boolean(i == 0)
}

// Abs returns the absolute value of the Integer.
// Example: Integer(-5).Abs() -> 5
func (i Integer) Abs() Integer {
	if i < 0 {
		return -i
	}
	return i
}

// ToF converts the Integer to a Float.
// Example: Integer(5).ToF() -> 5.0
func (i Integer) ToF() Float {
	return Float(i)
}

// ToS converts the Integer to a String.
// Example: Integer(123).ToS() -> "123"
func (i Integer) ToS() String {
	return String(fmt.Sprintf("%d", i))
}

// ToStr is an alias for ToS.
func (i Integer) ToStr() String {
	return i.ToS()
}

// Power raises the Integer to the given power.
// Example: Integer(2).Power(3) -> 8
func (i Integer) Power(power Integer) Integer {
	return Integer(math.Pow(float64(i), float64(power)))
}

// Sqrt returns the square root of the Integer.
// Example: Integer(16).Sqrt() -> 4
func (i Integer) Sqrt() Integer {
	return Integer(math.Sqrt(float64(i)))
}

// Min returns the smaller of two Integers.
// Example: Integer(3).Min(5) -> 3
func (i Integer) Min(other Integer) Integer {
	if i < other {
		return i
	}
	return other
}

// Max returns the larger of two Integers.
// Example: Integer(3).Max(5) -> 5
func (i Integer) Max(other Integer) Integer {
	if i > other {
		return i
	}
	return other
}

// Clamp clamps the Integer between min and max values.
// Example: Integer(5).Clamp(0, 3) -> 3
func (i Integer) Clamp(min, max Integer) Integer {
	if i < min {
		return min
	}
	if i > max {
		return max
	}
	return i
}

// Between checks if the Integer is between min and max (inclusive).
// Example: Integer(3).Between(1, 5) -> true
func (i Integer) Between(min, max Integer) Boolean {
	return Boolean(i >= min && i <= max)
}

// Times executes the given function n times.
// Example: Integer(3).Times(func(i Integer) { fmt.Println(i) })
func (i Integer) Times(fn func(Integer)) {
	for j := Integer(0); j < i; j++ {
		fn(j)
	}
}

// Upto executes the given function for each integer from i up to max (inclusive).
// Example: Integer(1).Upto(3, func(i Integer) { fmt.Println(i) })
func (i Integer) Upto(max Integer, fn func(Integer)) {
	for j := i; j <= max; j++ {
		fn(j)
	}
}

// Downto executes the given function for each integer from i down to min (inclusive).
// Example: Integer(3).Downto(1, func(i Integer) { fmt.Println(i) })
func (i Integer) Downto(min Integer, fn func(Integer)) {
	for j := i; j >= min; j-- {
		fn(j)
	}
}

// Step executes the given function for each integer from i to max, incrementing by step.
// Example: Integer(0).Step(10, 2, func(i Integer) { fmt.Println(i) })
func (i Integer) Step(max, step Integer, fn func(Integer)) {
	for j := i; j <= max; j += step {
		fn(j)
	}
}

// IsPrime checks if the Integer is a prime number.
// Example: Integer(7).IsPrime() -> true
func (i Integer) IsPrime() Boolean {
	if i < 2 {
		return false
	}
	if i == 2 {
		return true
	}
	if i%2 == 0 {
		return false
	}

	limit := Integer(math.Sqrt(float64(i)))
	for j := Integer(3); j <= limit; j += 2 {
		if i%j == 0 {
			return false
		}
	}
	return true
}

// Factorial returns the factorial of the Integer.
// Example: Integer(5).Factorial() -> 120
func (i Integer) Factorial() Integer {
	if i < 0 {
		return 0
	}
	if i <= 1 {
		return 1
	}

	result := Integer(1)
	for j := Integer(2); j <= i; j++ {
		result *= j
	}
	return result
}

// GCD returns the greatest common divisor of two Integers.
// Example: Integer(12).GCD(18) -> 6
func (i Integer) GCD(other Integer) Integer {
	a, b := i, other
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM returns the least common multiple of two Integers.
// Example: Integer(12).LCM(18) -> 36
func (i Integer) LCM(other Integer) Integer {
	if i == 0 || other == 0 {
		return 0
	}
	return (i * other) / i.GCD(other)
}

// Divisors returns an Array of all positive divisors of the Integer.
// Example: Integer(12).Divisors() -> [1, 2, 3, 4, 6, 12]
func (i Integer) Divisors() Array[Integer] {
	if i == 0 {
		return Array[Integer]{}
	}

	abs := i.Abs()
	divisors := make([]Integer, 0)

	for j := Integer(1); j <= abs; j++ {
		if abs%j == 0 {
			divisors = append(divisors, j)
		}
	}

	result := make(Array[Integer], len(divisors))
	copy(result, divisors)
	return result
}

// IsDivisibleBy checks if the Integer is divisible by the given divisor.
// Example: Integer(12).IsDivisibleBy(3) -> true
func (i Integer) IsDivisibleBy(divisor Integer) Boolean {
	if divisor == 0 {
		return false
	}
	return Boolean(i%divisor == 0)
}

// Next returns the next Integer (i + 1).
// Example: Integer(5).Next() -> 6
func (i Integer) Next() Integer {
	return i + 1
}

// Pred returns the previous Integer (i - 1).
// Example: Integer(5).Pred() -> 4
func (i Integer) Pred() Integer {
	return i - 1
}

// Succ is an alias for Next.
func (i Integer) Succ() Integer {
	return i.Next()
}
