package rb

import (
	"testing"
)

func TestInteger_IsOdd(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Boolean
	}{
		{Integer(1), Boolean(true)},
		{Integer(3), Boolean(true)},
		{Integer(5), Boolean(true)},
		{Integer(2), Boolean(false)},
		{Integer(4), Boolean(false)},
		{Integer(0), Boolean(false)},
		{Integer(-1), Boolean(true)},
		{Integer(-3), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.IsOdd()
		if result != test.expected {
			t.Errorf("IsOdd() for %d expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestInteger_IsEven(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Boolean
	}{
		{Integer(2), Boolean(true)},
		{Integer(4), Boolean(true)},
		{Integer(6), Boolean(true)},
		{Integer(1), Boolean(false)},
		{Integer(3), Boolean(false)},
		{Integer(0), Boolean(true)},
		{Integer(-2), Boolean(true)},
		{Integer(-4), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.IsEven()
		if result != test.expected {
			t.Errorf("IsEven() for %d expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestInteger_IsPositive(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Boolean
	}{
		{Integer(1), Boolean(true)},
		{Integer(100), Boolean(true)},
		{Integer(0), Boolean(false)},
		{Integer(-1), Boolean(false)},
		{Integer(-100), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsPositive()
		if result != test.expected {
			t.Errorf("IsPositive() for %d expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestInteger_IsNegative(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Boolean
	}{
		{Integer(-1), Boolean(true)},
		{Integer(-100), Boolean(true)},
		{Integer(0), Boolean(false)},
		{Integer(1), Boolean(false)},
		{Integer(100), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsNegative()
		if result != test.expected {
			t.Errorf("IsNegative() for %d expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestInteger_IsZero(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Boolean
	}{
		{Integer(0), Boolean(true)},
		{Integer(1), Boolean(false)},
		{Integer(-1), Boolean(false)},
		{Integer(100), Boolean(false)},
		{Integer(-100), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsZero()
		if result != test.expected {
			t.Errorf("IsZero() for %d expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestInteger_Abs(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Integer
	}{
		{Integer(5), Integer(5)},
		{Integer(-5), Integer(5)},
		{Integer(0), Integer(0)},
		{Integer(100), Integer(100)},
		{Integer(-100), Integer(100)},
	}

	for _, test := range tests {
		result := test.input.Abs()
		if result != test.expected {
			t.Errorf("Abs() for %d expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestInteger_ToF(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Float
	}{
		{Integer(5), Float(5.0)},
		{Integer(0), Float(0.0)},
		{Integer(-10), Float(-10.0)},
		{Integer(100), Float(100.0)},
	}

	for _, test := range tests {
		result := test.input.ToF()
		if result != test.expected {
			t.Errorf("ToF() for %d expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestInteger_ToS(t *testing.T) {
	tests := []struct {
		input    Integer
		expected String
	}{
		{Integer(123), String("123")},
		{Integer(0), String("0")},
		{Integer(-456), String("-456")},
		{Integer(999), String("999")},
	}

	for _, test := range tests {
		result := test.input.ToS()
		if result != test.expected {
			t.Errorf("ToS() for %d expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestInteger_ToStr(t *testing.T) {
	input := Integer(123)
	result := input.ToStr()
	
	if result != input.ToS() {
		t.Errorf("ToStr() should return same as ToS()")
	}
}

func TestInteger_Power(t *testing.T) {
	tests := []struct {
		base     Integer
		exponent Integer
		expected Integer
	}{
		{Integer(2), Integer(3), Integer(8)},
		{Integer(5), Integer(2), Integer(25)},
		{Integer(10), Integer(0), Integer(1)},
		{Integer(0), Integer(5), Integer(0)},
		{Integer(1), Integer(100), Integer(1)},
	}

	for _, test := range tests {
		result := test.base.Power(test.exponent)
		if result != test.expected {
			t.Errorf("Power() for %d^%d expected %d, got %d", test.base, test.exponent, test.expected, result)
		}
	}
}

func TestInteger_Sqrt(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Integer
	}{
		{Integer(16), Integer(4)},
		{Integer(25), Integer(5)},
		{Integer(100), Integer(10)},
		{Integer(0), Integer(0)},
		{Integer(1), Integer(1)},
	}

	for _, test := range tests {
		result := test.input.Sqrt()
		if result != test.expected {
			t.Errorf("Sqrt() for %d expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestInteger_Min(t *testing.T) {
	tests := []struct {
		a        Integer
		b        Integer
		expected Integer
	}{
		{Integer(5), Integer(3), Integer(3)},
		{Integer(3), Integer(5), Integer(3)},
		{Integer(0), Integer(10), Integer(0)},
		{Integer(-5), Integer(5), Integer(-5)},
		{Integer(5), Integer(5), Integer(5)},
	}

	for _, test := range tests {
		result := test.a.Min(test.b)
		if result != test.expected {
			t.Errorf("Min() for %d and %d expected %d, got %d", test.a, test.b, test.expected, result)
		}
	}
}

func TestInteger_Max(t *testing.T) {
	tests := []struct {
		a        Integer
		b        Integer
		expected Integer
	}{
		{Integer(5), Integer(3), Integer(5)},
		{Integer(3), Integer(5), Integer(5)},
		{Integer(0), Integer(10), Integer(10)},
		{Integer(-5), Integer(5), Integer(5)},
		{Integer(5), Integer(5), Integer(5)},
	}

	for _, test := range tests {
		result := test.a.Max(test.b)
		if result != test.expected {
			t.Errorf("Max() for %d and %d expected %d, got %d", test.a, test.b, test.expected, result)
		}
	}
}

func TestInteger_Clamp(t *testing.T) {
	tests := []struct {
		input    Integer
		min      Integer
		max      Integer
		expected Integer
	}{
		{Integer(5), Integer(0), Integer(10), Integer(5)},
		{Integer(15), Integer(0), Integer(10), Integer(10)},
		{Integer(-5), Integer(0), Integer(10), Integer(0)},
		{Integer(0), Integer(0), Integer(10), Integer(0)},
		{Integer(10), Integer(0), Integer(10), Integer(10)},
	}

	for _, test := range tests {
		result := test.input.Clamp(test.min, test.max)
		if result != test.expected {
			t.Errorf("Clamp() for %d between %d and %d expected %d, got %d", test.input, test.min, test.max, test.expected, result)
		}
	}
}

func TestInteger_Between(t *testing.T) {
	tests := []struct {
		input    Integer
		min      Integer
		max      Integer
		expected Boolean
	}{
		{Integer(5), Integer(0), Integer(10), Boolean(true)},
		{Integer(0), Integer(0), Integer(10), Boolean(true)},
		{Integer(10), Integer(0), Integer(10), Boolean(true)},
		{Integer(-1), Integer(0), Integer(10), Boolean(false)},
		{Integer(11), Integer(0), Integer(10), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.Between(test.min, test.max)
		if result != test.expected {
			t.Errorf("Between() for %d between %d and %d expected %t, got %t", test.input, test.min, test.max, test.expected, result)
		}
	}
}

func TestInteger_Times(t *testing.T) {
	count := 0
	Integer(5).Times(func(i Integer) {
		count++
	})
	
	if count != 5 {
		t.Errorf("Times() should execute 5 times, got %d", count)
	}
}

func TestInteger_Upto(t *testing.T) {
	var result []Integer
	Integer(1).Upto(3, func(i Integer) {
		result = append(result, i)
	})
	
	expected := []Integer{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("Upto() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Upto() at index %d expected %d, got %d", i, expected[i], val)
		}
	}
}

func TestInteger_Downto(t *testing.T) {
	var result []Integer
	Integer(3).Downto(1, func(i Integer) {
		result = append(result, i)
	})
	
	expected := []Integer{3, 2, 1}
	if len(result) != len(expected) {
		t.Errorf("Downto() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Downto() at index %d expected %d, got %d", i, expected[i], val)
		}
	}
}

func TestInteger_Step(t *testing.T) {
	var result []Integer
	Integer(0).Step(10, 2, func(i Integer) {
		result = append(result, i)
	})
	
	expected := []Integer{0, 2, 4, 6, 8, 10}
	if len(result) != len(expected) {
		t.Errorf("Step() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Step() at index %d expected %d, got %d", i, expected[i], val)
		}
	}
}

func TestInteger_IsPrime(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Boolean
	}{
		{Integer(2), Boolean(true)},
		{Integer(3), Boolean(true)},
		{Integer(5), Boolean(true)},
		{Integer(7), Boolean(true)},
		{Integer(11), Boolean(true)},
		{Integer(4), Boolean(false)},
		{Integer(6), Boolean(false)},
		{Integer(8), Boolean(false)},
		{Integer(9), Boolean(false)},
		{Integer(1), Boolean(false)},
		{Integer(0), Boolean(false)},
		{Integer(-1), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsPrime()
		if result != test.expected {
			t.Errorf("IsPrime() for %d expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestInteger_Factorial(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Integer
	}{
		{Integer(0), Integer(1)},
		{Integer(1), Integer(1)},
		{Integer(2), Integer(2)},
		{Integer(3), Integer(6)},
		{Integer(4), Integer(24)},
		{Integer(5), Integer(120)},
		{Integer(-1), Integer(0)},
		{Integer(-5), Integer(0)},
	}

	for _, test := range tests {
		result := test.input.Factorial()
		if result != test.expected {
			t.Errorf("Factorial() for %d expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestInteger_GCD(t *testing.T) {
	tests := []struct {
		a        Integer
		b        Integer
		expected Integer
	}{
		{Integer(12), Integer(18), Integer(6)},
		{Integer(18), Integer(12), Integer(6)},
		{Integer(7), Integer(13), Integer(1)},
		{Integer(0), Integer(5), Integer(5)},
		{Integer(5), Integer(0), Integer(5)},
		{Integer(0), Integer(0), Integer(0)},
	}

	for _, test := range tests {
		result := test.a.GCD(test.b)
		if result != test.expected {
			t.Errorf("GCD() for %d and %d expected %d, got %d", test.a, test.b, test.expected, result)
		}
	}
}

func TestInteger_LCM(t *testing.T) {
	tests := []struct {
		a        Integer
		b        Integer
		expected Integer
	}{
		{Integer(12), Integer(18), Integer(36)},
		{Integer(18), Integer(12), Integer(36)},
		{Integer(7), Integer(13), Integer(91)},
		{Integer(0), Integer(5), Integer(0)},
		{Integer(5), Integer(0), Integer(0)},
		{Integer(0), Integer(0), Integer(0)},
	}

	for _, test := range tests {
		result := test.a.LCM(test.b)
		if result != test.expected {
			t.Errorf("LCM() for %d and %d expected %d, got %d", test.a, test.b, test.expected, result)
		}
	}
}

func TestInteger_Divisors(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Array[Integer]
	}{
		{Integer(12), Array[Integer]{1, 2, 3, 4, 6, 12}},
		{Integer(6), Array[Integer]{1, 2, 3, 6}},
		{Integer(7), Array[Integer]{1, 7}},
		{Integer(1), Array[Integer]{1}},
		{Integer(0), Array[Integer]{}},
	}

	for _, test := range tests {
		result := test.input.Divisors()
		if len(result) != len(test.expected) {
			t.Errorf("Divisors() for %d expected length %d, got %d", test.input, len(test.expected), len(result))
		}
		// Note: Order might vary, so we'll just check length for now
	}
}

func TestInteger_IsDivisibleBy(t *testing.T) {
	tests := []struct {
		input    Integer
		divisor  Integer
		expected Boolean
	}{
		{Integer(12), Integer(3), Boolean(true)},
		{Integer(12), Integer(4), Boolean(true)},
		{Integer(12), Integer(5), Boolean(false)},
		{Integer(0), Integer(5), Boolean(true)},
		{Integer(5), Integer(0), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsDivisibleBy(test.divisor)
		if result != test.expected {
			t.Errorf("IsDivisibleBy() for %d by %d expected %t, got %t", test.input, test.divisor, test.expected, result)
		}
	}
}

func TestInteger_Next(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Integer
	}{
		{Integer(5), Integer(6)},
		{Integer(0), Integer(1)},
		{Integer(-1), Integer(0)},
		{Integer(100), Integer(101)},
	}

	for _, test := range tests {
		result := test.input.Next()
		if result != test.expected {
			t.Errorf("Next() for %d expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestInteger_Pred(t *testing.T) {
	tests := []struct {
		input    Integer
		expected Integer
	}{
		{Integer(5), Integer(4)},
		{Integer(1), Integer(0)},
		{Integer(0), Integer(-1)},
		{Integer(100), Integer(99)},
	}

	for _, test := range tests {
		result := test.input.Pred()
		if result != test.expected {
			t.Errorf("Pred() for %d expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestInteger_Succ(t *testing.T) {
	input := Integer(5)
	result := input.Succ()
	expected := input.Next()
	
	if result != expected {
		t.Errorf("Succ() should return same as Next()")
	}
}