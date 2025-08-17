// Package rb provides Ruby-inspired utility methods for Go types.
package rb

import (
	"fmt"
	"math"
)

// Float is a custom float64 type to emulate Ruby-like behavior.
type Float float64

// IsPositive checks if the Float is positive (greater than 0).
// Example: Float(3.14).IsPositive() -> true
func (f Float) IsPositive() Boolean {
	return Boolean(f > 0)
}

// IsNegative checks if the Float is negative (less than 0).
// Example: Float(-3.14).IsNegative() -> true
func (f Float) IsNegative() Boolean {
	return Boolean(f < 0)
}

// IsZero checks if the Float is zero.
// Example: Float(0.0).IsZero() -> true
func (f Float) IsZero() Boolean {
	return Boolean(f == 0)
}

// IsFinite checks if the Float is finite (not Inf or NaN).
// Example: Float(3.14).IsFinite() -> true
func (f Float) IsFinite() Boolean {
	return Boolean(!math.IsInf(float64(f), 0) && !math.IsNaN(float64(f)))
}

// IsInfinite checks if the Float is infinite.
// Example: Float(math.Inf(1)).IsInfinite() -> true
func (f Float) IsInfinite() Boolean {
	return Boolean(math.IsInf(float64(f), 0))
}

// IsNaN checks if the Float is NaN (Not a Number).
// Example: Float(math.NaN()).IsNaN() -> true
func (f Float) IsNaN() Boolean {
	return Boolean(math.IsNaN(float64(f)))
}

// Ceil returns the smallest integer greater than or equal to the Float.
// Example: Float(3.14).Ceil() -> 4.0
func (f Float) Ceil() Float {
	return Float(math.Ceil(float64(f)))
}

// Floor returns the largest integer less than or equal to the Float.
// Example: Float(3.14).Floor() -> 3.0
func (f Float) Floor() Float {
	return Float(math.Floor(float64(f)))
}

// Round returns the Float rounded to the nearest integer.
// Example: Float(3.5).Round() -> 4.0
func (f Float) Round() Float {
	return Float(math.Round(float64(f)))
}

// Abs returns the absolute value of the Float.
// Example: Float(-3.14).Abs() -> 3.14
func (f Float) Abs() Float {
	return Float(math.Abs(float64(f)))
}

// ToI converts the Float to an Integer by truncating.
// Example: Float(3.14).ToI() -> 3
func (f Float) ToI() Integer {
	return Integer(f)
}

// ToS converts the Float to a String.
// Example: Float(3.14).ToS() -> "3.14"
func (f Float) ToS() String {
	return String(fmt.Sprintf("%g", f))
}

// ToStr is an alias for ToS.
func (f Float) ToStr() String {
	return f.ToS()
}

// Power raises the Float to the given power.
// Example: Float(2.0).Power(3.0) -> 8.0
func (f Float) Power(power Float) Float {
	return Float(math.Pow(float64(f), float64(power)))
}

// Sqrt returns the square root of the Float.
// Example: Float(16.0).Sqrt() -> 4.0
func (f Float) Sqrt() Float {
	return Float(math.Sqrt(float64(f)))
}

// Sin returns the sine of the Float (in radians).
// Example: Float(math.Pi/2).Sin() -> 1.0
func (f Float) Sin() Float {
	return Float(math.Sin(float64(f)))
}

// Cos returns the cosine of the Float (in radians).
// Example: Float(0).Cos() -> 1.0
func (f Float) Cos() Float {
	return Float(math.Cos(float64(f)))
}

// Tan returns the tangent of the Float (in radians).
// Example: Float(0).Tan() -> 0.0
func (f Float) Tan() Float {
	return Float(math.Tan(float64(f)))
}

// Log returns the natural logarithm of the Float.
// Example: Float(math.E).Log() -> 1.0
func (f Float) Log() Float {
	return Float(math.Log(float64(f)))
}

// Log10 returns the base-10 logarithm of the Float.
// Example: Float(100.0).Log10() -> 2.0
func (f Float) Log10() Float {
	return Float(math.Log10(float64(f)))
}

// Exp returns e raised to the power of the Float.
// Example: Float(1.0).Exp() -> math.E
func (f Float) Exp() Float {
	return Float(math.Exp(float64(f)))
}

// Min returns the smaller of two Floats.
// Example: Float(3.14).Min(2.71) -> 2.71
func (f Float) Min(other Float) Float {
	if f < other {
		return f
	}
	return other
}

// Max returns the larger of two Floats.
// Example: Float(3.14).Max(2.71) -> 3.14
func (f Float) Max(other Float) Float {
	if f > other {
		return f
	}
	return other
}

// Clamp clamps the Float value between minVal and maxVal.
func (f Float) Clamp(minVal, maxVal Float) Float {
	if f < minVal {
		return minVal
	}
	if f > maxVal {
		return maxVal
	}
	return f
}

// Between checks if the Float value is between minVal and maxVal (inclusive).
func (f Float) Between(minVal, maxVal Float) Boolean {
	return Boolean(f >= minVal && f <= maxVal)
}

// IsInteger checks if the Float represents an integer value.
// Example: Float(3.0).IsInteger() -> true
func (f Float) IsInteger() Boolean {
	return Boolean(f == Float(int(f)))
}
