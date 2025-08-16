package rb

import (
	"math"
	"testing"
)

func TestFloat_IsPositive(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(3.14), Boolean(true)},
		{Float(0.001), Boolean(true)},
		{Float(0.0), Boolean(false)},
		{Float(-3.14), Boolean(false)},
		{Float(-0.001), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsPositive()
		if result != test.expected {
			t.Errorf("IsPositive() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestFloat_IsNegative(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(-3.14), Boolean(true)},
		{Float(-0.001), Boolean(true)},
		{Float(0.0), Boolean(false)},
		{Float(3.14), Boolean(false)},
		{Float(0.001), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsNegative()
		if result != test.expected {
			t.Errorf("IsNegative() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestFloat_IsZero(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(0.0), Boolean(true)},
		{Float(3.14), Boolean(false)},
		{Float(-3.14), Boolean(false)},
		{Float(0.001), Boolean(false)},
		{Float(-0.001), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsZero()
		if result != test.expected {
			t.Errorf("IsZero() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestFloat_IsFinite(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(3.14), Boolean(true)},
		{Float(0.0), Boolean(true)},
		{Float(-3.14), Boolean(true)},
		{Float(math.Inf(1)), Boolean(false)},
		{Float(math.Inf(-1)), Boolean(false)},
		{Float(math.NaN()), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsFinite()
		if result != test.expected {
			t.Errorf("IsFinite() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestFloat_IsInfinite(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(math.Inf(1)), Boolean(true)},
		{Float(math.Inf(-1)), Boolean(true)},
		{Float(3.14), Boolean(false)},
		{Float(0.0), Boolean(false)},
		{Float(-3.14), Boolean(false)},
		{Float(math.NaN()), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsInfinite()
		if result != test.expected {
			t.Errorf("IsInfinite() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestFloat_IsNaN(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(math.NaN()), Boolean(true)},
		{Float(3.14), Boolean(false)},
		{Float(0.0), Boolean(false)},
		{Float(-3.14), Boolean(false)},
		{Float(math.Inf(1)), Boolean(false)},
		{Float(math.Inf(-1)), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsNaN()
		if result != test.expected {
			t.Errorf("IsNaN() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestFloat_Ceil(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(3.14), Float(4.0)},
		{Float(3.0), Float(3.0)},
		{Float(-3.14), Float(-3.0)},
		{Float(-3.0), Float(-3.0)},
		{Float(0.0), Float(0.0)},
	}

	for _, test := range tests {
		result := test.input.Ceil()
		if result != test.expected {
			t.Errorf("Ceil() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Floor(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(3.14), Float(3.0)},
		{Float(3.0), Float(3.0)},
		{Float(-3.14), Float(-4.0)},
		{Float(-3.0), Float(-3.0)},
		{Float(0.0), Float(0.0)},
	}

	for _, test := range tests {
		result := test.input.Floor()
		if result != test.expected {
			t.Errorf("Floor() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Round(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(3.14), Float(3.0)},
		{Float(3.5), Float(4.0)},
		{Float(3.6), Float(4.0)},
		{Float(3.0), Float(3.0)},
		{Float(-3.14), Float(-3.0)},
		{Float(-3.5), Float(-4.0)},
		{Float(-3.6), Float(-4.0)},
		{Float(0.0), Float(0.0)},
	}

	for _, test := range tests {
		result := test.input.Round()
		if result != test.expected {
			t.Errorf("Round() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Abs(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(3.14), Float(3.14)},
		{Float(-3.14), Float(3.14)},
		{Float(0.0), Float(0.0)},
		{Float(100.5), Float(100.5)},
		{Float(-100.5), Float(100.5)},
	}

	for _, test := range tests {
		result := test.input.Abs()
		if result != test.expected {
			t.Errorf("Abs() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_ToI(t *testing.T) {
	tests := []struct {
		input    Float
		expected Integer
	}{
		{Float(3.14), Integer(3)},
		{Float(3.99), Integer(3)},
		{Float(0.0), Integer(0)},
		{Float(-3.14), Integer(-3)},
		{Float(-3.99), Integer(-3)},
	}

	for _, test := range tests {
		result := test.input.ToI()
		if result != test.expected {
			t.Errorf("ToI() for %f expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestFloat_ToS(t *testing.T) {
	tests := []struct {
		input    Float
		expected String
	}{
		{Float(3.14), String("3.14")},
		{Float(0.0), String("0")},
		{Float(-3.14), String("-3.14")},
		{Float(100.0), String("100")},
	}

	for _, test := range tests {
		result := test.input.ToS()
		if result != test.expected {
			t.Errorf("ToS() for %f expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestFloat_ToStr(t *testing.T) {
	input := Float(3.14)
	result := input.ToStr()
	
	if result != input.ToS() {
		t.Errorf("ToStr() should return same as ToS()")
	}
}

func TestFloat_Power(t *testing.T) {
	tests := []struct {
		base     Float
		exponent Float
		expected Float
	}{
		{Float(2.0), Float(3.0), Float(8.0)},
		{Float(5.0), Float(2.0), Float(25.0)},
		{Float(10.0), Float(0.0), Float(1.0)},
		{Float(0.0), Float(5.0), Float(0.0)},
		{Float(1.0), Float(100.0), Float(1.0)},
		{Float(2.5), Float(2.0), Float(6.25)},
	}

	for _, test := range tests {
		result := test.base.Power(test.exponent)
		if result != test.expected {
			t.Errorf("Power() for %f^%f expected %f, got %f", test.base, test.exponent, test.expected, result)
		}
	}
}

func TestFloat_Sqrt(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(16.0), Float(4.0)},
		{Float(25.0), Float(5.0)},
		{Float(100.0), Float(10.0)},
		{Float(0.0), Float(0.0)},
		{Float(1.0), Float(1.0)},
		{Float(2.0), Float(1.4142135623730951)},
	}

	for _, test := range tests {
		result := test.input.Sqrt()
		if result != test.expected {
			t.Errorf("Sqrt() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Sin(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(0.0), Float(0.0)},
		{Float(math.Pi / 2), Float(1.0)},
		{Float(math.Pi), Float(0.0)},
		{Float(3 * math.Pi / 2), Float(-1.0)},
		{Float(2 * math.Pi), Float(0.0)},
	}

	for _, test := range tests {
		result := test.input.Sin()
		if result != test.expected {
			t.Errorf("Sin() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Cos(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(0.0), Float(1.0)},
		{Float(math.Pi / 2), Float(0.0)},
		{Float(math.Pi), Float(-1.0)},
		{Float(3 * math.Pi / 2), Float(0.0)},
		{Float(2 * math.Pi), Float(1.0)},
	}

	for _, test := range tests {
		result := test.input.Cos()
		if result != test.expected {
			t.Errorf("Cos() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Tan(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(0.0), Float(0.0)},
		{Float(math.Pi / 4), Float(1.0)},
		{Float(math.Pi), Float(0.0)},
		{Float(3 * math.Pi / 4), Float(-1.0)},
	}

	for _, test := range tests {
		result := test.input.Tan()
		if result != test.expected {
			t.Errorf("Tan() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Log(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(math.E), Float(1.0)},
		{Float(1.0), Float(0.0)},
		{Float(math.E * math.E), Float(2.0)},
	}

	for _, test := range tests {
		result := test.input.Log()
		if result != test.expected {
			t.Errorf("Log() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Log10(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(100.0), Float(2.0)},
		{Float(1.0), Float(0.0)},
		{Float(1000.0), Float(3.0)},
		{Float(0.1), Float(-1.0)},
	}

	for _, test := range tests {
		result := test.input.Log10()
		if result != test.expected {
			t.Errorf("Log10() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Exp(t *testing.T) {
	tests := []struct {
		input    Float
		expected Float
	}{
		{Float(0.0), Float(1.0)},
		{Float(1.0), Float(math.E)},
		{Float(2.0), Float(math.E * math.E)},
		{Float(-1.0), Float(1.0 / math.E)},
	}

	for _, test := range tests {
		result := test.input.Exp()
		if result != test.expected {
			t.Errorf("Exp() for %f expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestFloat_Min(t *testing.T) {
	tests := []struct {
		a        Float
		b        Float
		expected Float
	}{
		{Float(5.0), Float(3.0), Float(3.0)},
		{Float(3.0), Float(5.0), Float(3.0)},
		{Float(0.0), Float(10.0), Float(0.0)},
		{Float(-5.0), Float(5.0), Float(-5.0)},
		{Float(5.0), Float(5.0), Float(5.0)},
	}

	for _, test := range tests {
		result := test.a.Min(test.b)
		if result != test.expected {
			t.Errorf("Min() for %f and %f expected %f, got %f", test.a, test.b, test.expected, result)
		}
	}
}

func TestFloat_Max(t *testing.T) {
	tests := []struct {
		a        Float
		b        Float
		expected Float
	}{
		{Float(5.0), Float(3.0), Float(5.0)},
		{Float(3.0), Float(5.0), Float(5.0)},
		{Float(0.0), Float(10.0), Float(10.0)},
		{Float(-5.0), Float(5.0), Float(5.0)},
		{Float(5.0), Float(5.0), Float(5.0)},
	}

	for _, test := range tests {
		result := test.a.Max(test.b)
		if result != test.expected {
			t.Errorf("Max() for %f and %f expected %f, got %f", test.a, test.b, test.expected, result)
		}
	}
}

func TestFloat_Clamp(t *testing.T) {
	tests := []struct {
		input    Float
		min      Float
		max      Float
		expected Float
	}{
		{Float(5.0), Float(0.0), Float(10.0), Float(5.0)},
		{Float(15.0), Float(0.0), Float(10.0), Float(10.0)},
		{Float(-5.0), Float(0.0), Float(10.0), Float(0.0)},
		{Float(0.0), Float(0.0), Float(10.0), Float(0.0)},
		{Float(10.0), Float(0.0), Float(10.0), Float(10.0)},
	}

	for _, test := range tests {
		result := test.input.Clamp(test.min, test.max)
		if result != test.expected {
			t.Errorf("Clamp() for %f between %f and %f expected %f, got %f", test.input, test.min, test.max, test.expected, result)
		}
	}
}

func TestFloat_Between(t *testing.T) {
	tests := []struct {
		input    Float
		min      Float
		max      Float
		expected Boolean
	}{
		{Float(5.0), Float(0.0), Float(10.0), Boolean(true)},
		{Float(0.0), Float(0.0), Float(10.0), Boolean(true)},
		{Float(10.0), Float(0.0), Float(10.0), Boolean(true)},
		{Float(-1.0), Float(0.0), Float(10.0), Boolean(false)},
		{Float(11.0), Float(0.0), Float(10.0), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.Between(test.min, test.max)
		if result != test.expected {
			t.Errorf("Between() for %f between %f and %f expected %t, got %t", test.input, test.min, test.max, test.expected, result)
		}
	}
}

func TestFloat_IsInteger(t *testing.T) {
	tests := []struct {
		input    Float
		expected Boolean
	}{
		{Float(3.0), Boolean(true)},
		{Float(0.0), Boolean(true)},
		{Float(-5.0), Boolean(true)},
		{Float(3.14), Boolean(false)},
		{Float(-3.14), Boolean(false)},
		{Float(0.001), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsInteger()
		if result != test.expected {
			t.Errorf("IsInteger() for %f expected %t, got %t", test.input, test.expected, result)
		}
	}
}