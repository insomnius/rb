package rb

import (
	"testing"
)

func TestRange_NewRange(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	
	if range1.Begin != 1 || range1.End != 5 || range1.Exclusive {
		t.Errorf("NewRange() failed to create correct range: begin=%d, end=%d, exclusive=%t", range1.Begin, range1.End, range1.Exclusive)
	}
}

func TestRange_NewExclusiveRange(t *testing.T) {
	range1 := NewExclusiveRange[Integer](Integer(1), Integer(5))
	
	if range1.Begin != 1 || range1.End != 5 || !range1.Exclusive {
		t.Errorf("NewExclusiveRange() failed to create correct range: begin=%d, end=%d, exclusive=%t", range1.Begin, range1.End, range1.Exclusive)
	}
}

func TestRange_Each(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(3))
	var result []Integer
	
	range1.Each(func(i Integer) {
		result = append(result, i)
	})
	
	expected := []Integer{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("Each() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Each() at index %d expected %d, got %d", i, expected[i], val)
		}
	}
}

func TestRange_EachWithIndex(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(3))
	var result []Integer
	var indices []Integer
	
	range1.EachWithIndex(func(i Integer, idx Integer) {
		result = append(result, i)
		indices = append(indices, idx)
	})
	
	expected := []Integer{1, 2, 3}
	expectedIndices := []Integer{0, 1, 2}
	
	if len(result) != len(expected) {
		t.Errorf("EachWithIndex() expected length %d, got %d", len(expected), len(result))
	}
	
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("EachWithIndex() at index %d expected value %d, got %d", i, expected[i], val)
		}
		if indices[i] != expectedIndices[i] {
			t.Errorf("EachWithIndex() at index %d expected index %d, got %d", i, expectedIndices[i], indices[i])
		}
	}
}

func TestRange_Include(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	
	tests := []struct {
		value    Integer
		expected Boolean
	}{
		{Integer(1), Boolean(true)},
		{Integer(3), Boolean(true)},
		{Integer(5), Boolean(true)},
		{Integer(0), Boolean(false)},
		{Integer(6), Boolean(false)},
	}

	for _, test := range tests {
		result := range1.Include(test.value)
		if result != test.expected {
			t.Errorf("Include() for value %d expected %t, got %t", test.value, test.expected, result)
		}
	}
}

func TestRange_Cover(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	
	tests := []struct {
		value    Integer
		expected Boolean
	}{
		{Integer(1), Boolean(true)},
		{Integer(3), Boolean(true)},
		{Integer(5), Boolean(true)},
		{Integer(0), Boolean(false)},
		{Integer(6), Boolean(false)},
	}

	for _, test := range tests {
		result := range1.Cover(test.value)
		if result != test.expected {
			t.Errorf("Cover() for value %d expected %t, got %t", test.value, test.expected, result)
		}
	}
}

func TestRange_Step(t *testing.T) {
	range1 := NewRange[Integer](Integer(0), Integer(10))
	var result []Integer
	
	range1.Step(Integer(2), func(i Integer) {
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

func TestRange_ToArray(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(3))
	result := range1.ToArray()
	
	expected := Array[Integer]{1, 2, 3}
	if len(result) != len(expected) {
		t.Errorf("ToArray() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("ToArray() at index %d expected %d, got %d", i, expected[i], val)
		}
	}
}

func TestRange_Size(t *testing.T) {
	tests := []struct {
		range1   Range[Integer]
		expected Integer
	}{
		{NewRange[Integer](Integer(1), Integer(5)), Integer(5)},
		{NewRange[Integer](Integer(0), Integer(0)), Integer(1)},
		{NewRange[Integer](Integer(10), Integer(15)), Integer(6)},
		{NewExclusiveRange[Integer](Integer(1), Integer(5)), Integer(4)},
	}

	for _, test := range tests {
		result := test.range1.Size()
		if result != test.expected {
			t.Errorf("Size() for range %v expected %d, got %d", test.range1, test.expected, result)
		}
	}
}

func TestRange_Length(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.Length()
	expected := range1.Size()
	
	if result != expected {
		t.Errorf("Length() should return same as Size(), expected %d, got %d", expected, result)
	}
}

func TestRange_IsEmpty(t *testing.T) {
	tests := []struct {
		range1   Range[Integer]
		expected Boolean
	}{
		{NewRange[Integer](Integer(1), Integer(5)), Boolean(false)},
		{NewRange[Integer](Integer(1), Integer(1)), Boolean(false)},
		{NewRange[Integer](Integer(5), Integer(1)), Boolean(true)},
		{NewExclusiveRange[Integer](Integer(1), Integer(2)), Boolean(false)},
		{NewExclusiveRange[Integer](Integer(1), Integer(1)), Boolean(true)},
	}

	for _, test := range tests {
		result := test.range1.IsEmpty()
		if result != test.expected {
			t.Errorf("IsEmpty() for range %v expected %t, got %t", test.range1, test.expected, result)
		}
	}
}

func TestRange_Min(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.Min()
	
	if result != 1 {
		t.Errorf("Min() expected 1, got %d", result)
	}
}

func TestRange_Max(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.Max()
	
	if result != 5 {
		t.Errorf("Max() expected 5, got %d", result)
	}
}

func TestRange_First(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.First()
	
	if result != 1 {
		t.Errorf("First() expected 1, got %d", result)
	}
}

func TestRange_Last(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.Last()
	
	if result != 5 {
		t.Errorf("Last() expected 5, got %d", result)
	}
}

func TestRange_BeginValue(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.BeginValue()
	
	if result != 1 {
		t.Errorf("BeginValue() expected 1, got %d", result)
	}
}

func TestRange_EndValue(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.EndValue()
	
	if result != 5 {
		t.Errorf("EndValue() expected 5, got %d", result)
	}
}

func TestRange_IsExclusive(t *testing.T) {
	tests := []struct {
		range1   Range[Integer]
		expected Boolean
	}{
		{NewRange[Integer](Integer(1), Integer(5)), Boolean(false)},
		{NewExclusiveRange[Integer](Integer(1), Integer(5)), Boolean(true)},
	}

	for _, test := range tests {
		result := test.range1.IsExclusive()
		if result != test.expected {
			t.Errorf("IsExclusive() for range %v expected %t, got %t", test.range1, test.expected, result)
		}
	}
}

func TestRange_IsInclusive(t *testing.T) {
	tests := []struct {
		range1   Range[Integer]
		expected Boolean
	}{
		{NewRange[Integer](Integer(1), Integer(5)), Boolean(true)},
		{NewExclusiveRange[Integer](Integer(1), Integer(5)), Boolean(false)},
	}

	for _, test := range tests {
		result := test.range1.IsInclusive()
		if result != test.expected {
			t.Errorf("IsInclusive() for range %v expected %t, got %t", test.range1, test.expected, result)
		}
	}
}

func TestRange_Reverse(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.Reverse()
	
	// Check that begin and end are swapped
	if result.Begin != 5 || result.End != 1 {
		t.Errorf("Reverse() expected begin=5, end=1, got begin=%d, end=%d", result.Begin, result.End)
	}
	
	// Check that exclusive property is preserved
	if result.Exclusive != range1.Exclusive {
		t.Errorf("Reverse() should preserve exclusive property")
	}
}

func TestRange_Overlap(t *testing.T) {
	tests := []struct {
		range1   Range[Integer]
		range2   Range[Integer]
		expected Boolean
	}{
		{NewRange[Integer](Integer(1), Integer(5)), NewRange[Integer](Integer(3), Integer(7)), Boolean(true)},
		{NewRange[Integer](Integer(1), Integer(5)), NewRange[Integer](Integer(6), Integer(10)), Boolean(false)},
		{NewRange[Integer](Integer(1), Integer(5)), NewRange[Integer](Integer(5), Integer(10)), Boolean(true)},
		{NewRange[Integer](Integer(1), Integer(5)), NewRange[Integer](Integer(0), Integer(1)), Boolean(true)},
		{NewRange[Integer](Integer(1), Integer(5)), NewRange[Integer](Integer(2), Integer(4)), Boolean(true)},
	}

	for _, test := range tests {
		result := test.range1.Overlap(test.range2)
		if result != test.expected {
			t.Errorf("Overlap() for ranges %v and %v expected %t, got %t", test.range1, test.range2, test.expected, result)
		}
	}
}

func TestRange_Contains(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	
	tests := []struct {
		other    Range[Integer]
		expected Boolean
	}{
		{NewRange[Integer](Integer(2), Integer(4)), Boolean(true)},
		{NewRange[Integer](Integer(1), Integer(5)), Boolean(true)},
		{NewRange[Integer](Integer(0), Integer(6)), Boolean(false)},
		{NewRange[Integer](Integer(6), Integer(10)), Boolean(false)},
	}

	for _, test := range tests {
		result := range1.Contains(test.other)
		if result != test.expected {
			t.Errorf("Contains() for range %v containing %v expected %t, got %t", test.range1, test.other, test.expected, result)
		}
	}
}

func TestRange_ToS(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.ToS()
	
	expected := "1..5"
	if result != expected {
		t.Errorf("ToS() expected '%s', got '%s'", expected, result)
	}
	
	exclusiveRange := NewExclusiveRange[Integer](Integer(1), Integer(5))
	result = exclusiveRange.ToS()
	
	expected = "1...5"
	if result != expected {
		t.Errorf("ToS() for exclusive range expected '%s', got '%s'", expected, result)
	}
}

func TestRange_ToStr(t *testing.T) {
	range1 := NewRange[Integer](Integer(1), Integer(5))
	result := range1.ToStr()
	
	if result != range1.ToS() {
		t.Errorf("ToStr() should return same as ToS()")
	}
}

func TestRange_Clone(t *testing.T) {
	original := NewRange[Integer](Integer(1), Integer(5))
	result := original.Clone()
	
	// Check that clone has same values
	if result.Begin != original.Begin || result.End != original.End || result.Exclusive != original.Exclusive {
		t.Errorf("Clone() failed to copy values correctly")
	}
	
	// Check that clone is independent
	result.Begin = Integer(10)
	if original.Begin == Integer(10) {
		t.Error("Clone() should create independent copy")
	}
}

func TestRange_Float(t *testing.T) {
	range1 := NewRange[Float](Float(1.0), Float(5.0))
	
	if range1.Begin != Float(1.0) || range1.End != Float(5.0) {
		t.Errorf("Range[Float] failed to create correctly: begin=%f, end=%f", range1.Begin, range1.End)
	}
	
	// Test Float range operations
	var result []Float
	range1.Each(func(f Float) {
		result = append(result, f)
	})
	
	expected := []Float{1.0, 2.0, 3.0, 4.0, 5.0}
	if len(result) != len(expected) {
		t.Errorf("Float range Each() expected length %d, got %d", len(expected), len(result))
	}
}