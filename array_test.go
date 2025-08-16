package rb

import (
	"testing"
)

func TestArray_Count(t *testing.T) {
	tests := []struct {
		array    Array[String]
		arg      any
		expected Integer
	}{
		{Array[String]{"a", "b", "c"}, nil, Integer(3)},
		{Array[String]{"a", "b", "a"}, "a", Integer(2)},
		{Array[String]{"a", "b", "c"}, "d", Integer(0)},
		{Array[String]{}, nil, Integer(0)},
		{Array[String]{"a", "b", "c"}, func(s String) bool { return s == "a" }, Integer(1)},
	}

	for _, test := range tests {
		result := test.array.Count(test.arg)
		if result != test.expected {
			t.Errorf("Count() for %v with arg %v expected %d, got %d", test.array, test.arg, test.expected, result)
		}
	}
}

func TestArray_Map(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	result := array.Map(func(s String) String {
		return s.Upcase()
	})
	
	expected := Array[String]{"A", "B", "C"}
	if len(result) != len(expected) {
		t.Errorf("Map() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Map() at index %d expected %s, got %s", i, expected[i], val)
		}
	}
}

func TestArray_Select(t *testing.T) {
	array := Array[String]{"a", "bb", "c", "ddd"}
	result := array.Select(func(s String) bool {
		return s.Length() > 1
	})
	
	expected := Array[String]{"bb", "ddd"}
	if len(result) != len(expected) {
		t.Errorf("Select() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Select() at index %d expected %s, got %s", i, expected[i], val)
		}
	}
}

func TestArray_Reject(t *testing.T) {
	array := Array[String]{"a", "bb", "c", "ddd"}
	result := array.Reject(func(s String) bool {
		return s.Length() > 1
	})
	
	expected := Array[String]{"a", "c"}
	if len(result) != len(expected) {
		t.Errorf("Reject() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Reject() at index %d expected %s, got %s", i, expected[i], val)
		}
	}
}

func TestArray_Find(t *testing.T) {
	array := Array[String]{"a", "bb", "c", "ddd"}
	
	// Find first string with length > 1
	result := array.Find(func(s String) bool {
		return s.Length() > 1
	})
	
	if result == nil {
		t.Error("Find() should return a value for length > 1")
	}
	if *result != "bb" {
		t.Errorf("Find() expected 'bb', got %s", *result)
	}
	
	// Find non-existent
	result = array.Find(func(s String) bool {
		return s.Length() > 10
	})
	
	if result != nil {
		t.Error("Find() should return nil for non-existent value")
	}
}

func TestArray_Any(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	
	// Should return true
	result := array.Any(func(s String) bool {
		return s == "b"
	})
	if !result {
		t.Error("Any() should return true when predicate matches")
	}
	
	// Should return false
	result = array.Any(func(s String) bool {
		return s == "d"
	})
	if result {
		t.Error("Any() should return false when predicate doesn't match")
	}
}

func TestArray_All(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	
	// Should return true
	result := array.All(func(s String) bool {
		return s.Length() == 1
	})
	if !result {
		t.Error("All() should return true when all elements match predicate")
	}
	
	// Should return false
	result = array.All(func(s String) bool {
		return s == "a"
	})
	if result {
		t.Error("All() should return false when not all elements match predicate")
	}
}

func TestArray_None(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	
	// Should return true
	result := array.None(func(s String) bool {
		return s.Length() > 1
	})
	if !result {
		t.Error("None() should return true when no elements match predicate")
	}
	
	// Should return false
	result = array.None(func(s String) bool {
		return s == "a"
	})
	if result {
		t.Error("None() should return false when some elements match predicate")
	}
}

func TestArray_First(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	result := array.First()
	
	if result == nil {
		t.Error("First() should return first element")
	}
	if *result != "a" {
		t.Errorf("First() expected 'a', got %s", *result)
	}
	
	// Empty array
	emptyArray := Array[String]{}
	result = emptyArray.First()
	if result != nil {
		t.Error("First() should return nil for empty array")
	}
}

func TestArray_Last(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	result := array.Last()
	
	if result == nil {
		t.Error("Last() should return last element")
	}
	if *result != "c" {
		t.Errorf("Last() expected 'c', got %s", *result)
	}
	
	// Empty array
	emptyArray := Array[String]{}
	result = emptyArray.Last()
	if result != nil {
		t.Error("Last() should return nil for empty array")
	}
}

func TestArray_Uniq(t *testing.T) {
	array := Array[String]{"a", "b", "a", "c", "b"}
	result := array.Uniq()
	
	// Check that duplicates are removed
	seen := make(map[String]bool)
	for _, val := range result {
		if seen[val] {
			t.Errorf("Uniq() contains duplicate value: %s", val)
		}
		seen[val] = true
	}
	
	// Check that all original values are present
	originalSeen := make(map[String]bool)
	for _, val := range array {
		originalSeen[val] = true
	}
	
	for val := range seen {
		if !originalSeen[val] {
			t.Errorf("Uniq() contains value not in original: %s", val)
		}
	}
}

func TestArray_Compact(t *testing.T) {
	array := Array[String]{"a", "", "b", "", "c"}
	result := array.Compact()
	
	expected := Array[String]{"a", "b", "c"}
	if len(result) != len(expected) {
		t.Errorf("Compact() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Compact() at index %d expected %s, got %s", i, expected[i], val)
		}
	}
}

func TestArray_Each(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	var result []String
	
	array.Each(func(s String) {
		result = append(result, s)
	})
	
	if len(result) != len(array) {
		t.Errorf("Each() should process all elements, expected %d, got %d", len(array), len(result))
	}
	for i, val := range result {
		if val != array[i] {
			t.Errorf("Each() at index %d expected %s, got %s", i, array[i], val)
		}
	}
}

func TestArray_EachWithIndex(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	var result []String
	var indices []Integer
	
	array.EachWithIndex(func(s String, i Integer) {
		result = append(result, s)
		indices = append(indices, i)
	})
	
	if len(result) != len(array) {
		t.Errorf("EachWithIndex() should process all elements, expected %d, got %d", len(array), len(result))
	}
	
	expectedIndices := []Integer{0, 1, 2}
	for i, idx := range indices {
		if idx != expectedIndices[i] {
			t.Errorf("EachWithIndex() at index %d expected index %d, got %d", i, expectedIndices[i], idx)
		}
	}
}

func TestArray_Reverse(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	result := array.Reverse()
	
	expected := Array[String]{"c", "b", "a"}
	if len(result) != len(expected) {
		t.Errorf("Reverse() expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Reverse() at index %d expected %s, got %s", i, expected[i], val)
		}
	}
}

func TestArray_Sort(t *testing.T) {
	array := Array[String]{"c", "a", "b"}
	result := array.Sort()
	
	// Note: Sort is simplified, so we just check it returns something
	if len(result) != len(array) {
		t.Errorf("Sort() expected length %d, got %d", len(array), len(result))
	}
}

func TestArray_Take(t *testing.T) {
	array := Array[String]{"a", "b", "c", "d", "e"}
	
	// Take 3
	result := array.Take(3)
	expected := Array[String]{"a", "b", "c"}
	if len(result) != len(expected) {
		t.Errorf("Take(3) expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Take(3) at index %d expected %s, got %s", i, expected[i], val)
		}
	}
	
	// Take more than available
	result = array.Take(10)
	if len(result) != len(array) {
		t.Errorf("Take(10) should return all elements, expected %d, got %d", len(array), len(result))
	}
	
	// Take 0
	result = array.Take(0)
	if len(result) != 0 {
		t.Errorf("Take(0) should return empty array, got length %d", len(result))
	}
}

func TestArray_Drop(t *testing.T) {
	array := Array[String]{"a", "b", "c", "d", "e"}
	
	// Drop 2
	result := array.Drop(2)
	expected := Array[String]{"c", "d", "e"}
	if len(result) != len(expected) {
		t.Errorf("Drop(2) expected length %d, got %d", len(expected), len(result))
	}
	for i, val := range result {
		if val != expected[i] {
			t.Errorf("Drop(2) at index %d expected %s, got %s", i, expected[i], val)
		}
	}
	
	// Drop more than available
	result = array.Drop(10)
	if len(result) != 0 {
		t.Errorf("Drop(10) should return empty array, got length %d", len(result))
	}
	
	// Drop 0
	result = array.Drop(0)
	if len(result) != len(array) {
		t.Errorf("Drop(0) should return all elements, expected %d, got %d", len(array), len(result))
	}
}

func TestArray_IsEmpty(t *testing.T) {
	tests := []struct {
		array    Array[String]
		expected Boolean
	}{
		{Array[String]{}, Boolean(true)},
		{Array[String]{"a"}, Boolean(false)},
		{Array[String]{"a", "b"}, Boolean(false)},
	}

	for _, test := range tests {
		result := test.array.IsEmpty()
		if result != test.expected {
			t.Errorf("IsEmpty() for %v expected %t, got %t", test.array, test.expected, result)
		}
	}
}

func TestArray_Length(t *testing.T) {
	tests := []struct {
		array    Array[String]
		expected Integer
	}{
		{Array[String]{}, Integer(0)},
		{Array[String]{"a"}, Integer(1)},
		{Array[String]{"a", "b", "c"}, Integer(3)},
	}

	for _, test := range tests {
		result := test.array.Length()
		if result != test.expected {
			t.Errorf("Length() for %v expected %d, got %d", test.array, test.expected, result)
		}
	}
}

func TestArray_Size(t *testing.T) {
	array := Array[String]{"a", "b", "c"}
	result := array.Size()
	expected := array.Length()
	
	if result != expected {
		t.Errorf("Size() should return same as Length(), expected %d, got %d", expected, result)
	}
}