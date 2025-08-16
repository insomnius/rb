package rb

import (
	"testing"
)

func TestBoolean_ToS(t *testing.T) {
	tests := []struct {
		input    Boolean
		expected String
	}{
		{Boolean(true), String("true")},
		{Boolean(false), String("false")},
	}

	for _, test := range tests {
		result := test.input.ToS()
		if result != test.expected {
			t.Errorf("ToS() for %t expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestBoolean_ToStr(t *testing.T) {
	input := Boolean(true)
	result := input.ToStr()
	
	if result != input.ToS() {
		t.Errorf("ToStr() should return same as ToS()")
	}
}

func TestBoolean_And(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(true)},
		{Boolean(true), Boolean(false), Boolean(false)},
		{Boolean(false), Boolean(true), Boolean(false)},
		{Boolean(false), Boolean(false), Boolean(false)},
	}

	for _, test := range tests {
		result := test.a.And(test.b)
		if result != test.expected {
			t.Errorf("And() for %t AND %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_Or(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(true)},
		{Boolean(true), Boolean(false), Boolean(true)},
		{Boolean(false), Boolean(true), Boolean(true)},
		{Boolean(false), Boolean(false), Boolean(false)},
	}

	for _, test := range tests {
		result := test.a.Or(test.b)
		if result != test.expected {
			t.Errorf("Or() for %t OR %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_Not(t *testing.T) {
	tests := []struct {
		input    Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(false)},
		{Boolean(false), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.Not()
		if result != test.expected {
			t.Errorf("Not() for %t expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestBoolean_Xor(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(false)},
		{Boolean(true), Boolean(false), Boolean(true)},
		{Boolean(false), Boolean(true), Boolean(true)},
		{Boolean(false), Boolean(false), Boolean(false)},
	}

	for _, test := range tests {
		result := test.a.Xor(test.b)
		if result != test.expected {
			t.Errorf("Xor() for %t XOR %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_Nand(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(false)},
		{Boolean(true), Boolean(false), Boolean(true)},
		{Boolean(false), Boolean(true), Boolean(true)},
		{Boolean(false), Boolean(false), Boolean(true)},
	}

	for _, test := range tests {
		result := test.a.Nand(test.b)
		if result != test.expected {
			t.Errorf("Nand() for %t NAND %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_Nor(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(false)},
		{Boolean(true), Boolean(false), Boolean(false)},
		{Boolean(false), Boolean(true), Boolean(false)},
		{Boolean(false), Boolean(false), Boolean(true)},
	}

	for _, test := range tests {
		result := test.a.Nor(test.b)
		if result != test.expected {
			t.Errorf("Nor() for %t NOR %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_Xnor(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(true)},
		{Boolean(true), Boolean(false), Boolean(false)},
		{Boolean(false), Boolean(true), Boolean(false)},
		{Boolean(false), Boolean(false), Boolean(true)},
	}

	for _, test := range tests {
		result := test.a.Xnor(test.b)
		if result != test.expected {
			t.Errorf("Xnor() for %t XNOR %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_Implies(t *testing.T) {
	tests := []struct {
		a        Boolean
		b        Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true), Boolean(true)},
		{Boolean(true), Boolean(false), Boolean(false)},
		{Boolean(false), Boolean(true), Boolean(true)},
		{Boolean(false), Boolean(false), Boolean(true)},
	}

	for _, test := range tests {
		result := test.a.Implies(test.b)
		if result != test.expected {
			t.Errorf("Implies() for %t implies %t expected %t, got %t", test.a, test.b, test.expected, result)
		}
	}
}

func TestBoolean_ToI(t *testing.T) {
	tests := []struct {
		input    Boolean
		expected Integer
	}{
		{Boolean(true), Integer(1)},
		{Boolean(false), Integer(0)},
	}

	for _, test := range tests {
		result := test.input.ToI()
		if result != test.expected {
			t.Errorf("ToI() for %t expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestBoolean_ToF(t *testing.T) {
	tests := []struct {
		input    Boolean
		expected Float
	}{
		{Boolean(true), Float(1.0)},
		{Boolean(false), Float(0.0)},
	}

	for _, test := range tests {
		result := test.input.ToF()
		if result != test.expected {
			t.Errorf("ToF() for %t expected %f, got %f", test.input, test.expected, result)
		}
	}
}

func TestBoolean_IsTrue(t *testing.T) {
	tests := []struct {
		input    Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(true)},
		{Boolean(false), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsTrue()
		if result != test.expected {
			t.Errorf("IsTrue() for %t expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestBoolean_IsFalse(t *testing.T) {
	tests := []struct {
		input    Boolean
		expected Boolean
	}{
		{Boolean(true), Boolean(false)},
		{Boolean(false), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.IsFalse()
		if result != test.expected {
			t.Errorf("IsFalse() for %t expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestBoolean_IfTrue(t *testing.T) {
	executed := false
	Boolean(true).IfTrue(func() {
		executed = true
	})
	
	if !executed {
		t.Error("IfTrue() should execute function when boolean is true")
	}
	
	executed = false
	Boolean(false).IfTrue(func() {
		executed = true
	})
	
	if executed {
		t.Error("IfTrue() should not execute function when boolean is false")
	}
}

func TestBoolean_IfFalse(t *testing.T) {
	executed := false
	Boolean(false).IfFalse(func() {
		executed = true
	})
	
	if !executed {
		t.Error("IfFalse() should execute function when boolean is false")
	}
	
	executed = false
	Boolean(true).IfFalse(func() {
		executed = true
	})
	
	if executed {
		t.Error("IfFalse() should not execute function when boolean is true")
	}
}

func TestBoolean_If(t *testing.T) {
	trueExecuted := false
	falseExecuted := false
	
	Boolean(true).If(
		func() { trueExecuted = true },
		func() { falseExecuted = true },
	)
	
	if !trueExecuted || falseExecuted {
		t.Error("If() should execute true function when boolean is true")
	}
	
	trueExecuted = false
	falseExecuted = false
	
	Boolean(false).If(
		func() { trueExecuted = true },
		func() { falseExecuted = true },
	)
	
	if trueExecuted || !falseExecuted {
		t.Error("If() should execute false function when boolean is false")
	}
}

func TestBoolean_Ternary(t *testing.T) {
	tests := []struct {
		input    Boolean
		ifTrue   any
		ifFalse  any
		expected any
	}{
		{Boolean(true), "yes", "no", "yes"},
		{Boolean(false), "yes", "no", "no"},
		{Boolean(true), 42, 0, 42},
		{Boolean(false), 42, 0, 0},
		{Boolean(true), true, false, true},
		{Boolean(false), true, false, false},
	}

	for _, test := range tests {
		result := test.input.Ternary(test.ifTrue, test.ifFalse)
		if result != test.expected {
			t.Errorf("Ternary() for %t with %v and %v expected %v, got %v", test.input, test.ifTrue, test.ifFalse, test.expected, result)
		}
	}
}