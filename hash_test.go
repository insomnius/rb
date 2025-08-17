package rb

import (
	"testing"
)

func TestHash_Keys(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}
	result := hash.Keys()

	if len(result) != 3 {
		t.Errorf("Keys() expected length 3, got %d", len(result))
	}

	// Check that all keys are present
	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	for _, key := range result {
		if !expectedKeys[string(key.(String))] {
			t.Errorf("Keys() contains unexpected key: %s", key)
		}
	}
}

func TestHash_Values(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}
	result := hash.Values()

	if len(result) != 3 {
		t.Errorf("Values() expected length 3, got %d", len(result))
	}

	// Check that all values are present
	expectedValues := map[int]bool{1: true, 2: true, 3: true}
	for _, value := range result {
		if !expectedValues[int(value.(Integer))] {
			t.Errorf("Values() contains unexpected value: %d", value)
		}
	}
}

func TestHash_HasKey(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		key      String
		expected Boolean
	}{
		{String("a"), Boolean(true)},
		{String("b"), Boolean(true)},
		{String("c"), Boolean(true)},
		{String("d"), Boolean(false)},
		{String(""), Boolean(false)},
	}

	for _, test := range tests {
		result := hash.HasKey(test.key)
		if result != test.expected {
			t.Errorf("HasKey() for key '%s' expected %t, got %t", test.key, test.expected, result)
		}
	}
}

func TestHash_Get(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	tests := []struct {
		key          String
		defaultValue Integer
		expected     Integer
	}{
		{String("a"), Integer(0), Integer(1)},
		{String("b"), Integer(0), Integer(2)},
		{String("c"), Integer(0), Integer(3)},
		{String("d"), Integer(0), Integer(0)},
		{String(""), Integer(42), Integer(42)},
	}

	for _, test := range tests {
		result := hash.Get(test.key, test.defaultValue)
		if result != test.expected {
			t.Errorf("Get() for key '%s' with default %d expected %d, got %d", test.key, test.defaultValue, test.expected, result)
		}
	}
}

func TestHash_Fetch(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	// Test existing key
	result := hash.Fetch(String("a"))
	if result != 1 {
		t.Errorf("Fetch() for key 'a' expected 1, got %d", result)
	}

	// Test non-existing key (should panic)
	defer func() {
		if r := recover(); r == nil {
			t.Error("Fetch() for non-existing key should panic")
		}
	}()
	hash.Fetch(String("d"))
}

func TestHash_Set(t *testing.T) {
	hash := Hash[String, Integer]{}

	// Set new key
	hash.Set(String("a"), Integer(1))
	if hash["a"] != 1 {
		t.Errorf("Set() failed to set key 'a' to value 1")
	}

	// Overwrite existing key
	hash.Set(String("a"), Integer(42))
	if hash["a"] != 42 {
		t.Errorf("Set() failed to overwrite key 'a' with value 42")
	}
}

func TestHash_Delete(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	// Delete existing key
	result := hash.Delete(String("b"))
	if result != 2 {
		t.Errorf("Delete() for key 'b' expected to return 2, got %d", result)
	}
	if hash.HasKey(String("b")) {
		t.Error("Delete() failed to remove key 'b'")
	}

	// Delete non-existing key
	result = hash.Delete(String("d"))
	if result != 0 {
		t.Errorf("Delete() for non-existing key 'd' expected to return 0, got %d", result)
	}
}

func TestHash_Clear(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	hash.Clear()
	if len(hash) != 0 {
		t.Errorf("Clear() failed to clear hash, got length %d", len(hash))
	}
}

func TestHash_Size(t *testing.T) {
	tests := []struct {
		hash     Hash[String, Integer]
		expected Integer
	}{
		{Hash[String, Integer]{}, Integer(0)},
		{Hash[String, Integer]{"a": 1}, Integer(1)},
		{Hash[String, Integer]{"a": 1, "b": 2, "c": 3}, Integer(3)},
	}

	for _, test := range tests {
		result := test.hash.Size()
		if result != test.expected {
			t.Errorf("Size() for %v expected %d, got %d", test.hash, test.expected, result)
		}
	}
}

func TestHash_Length(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}
	result := hash.Length()
	expected := hash.Size()

	if result != expected {
		t.Errorf("Length() should return same as Size(), expected %d, got %d", expected, result)
	}
}

func TestHash_IsEmpty(t *testing.T) {
	tests := []struct {
		hash     Hash[String, Integer]
		expected Boolean
	}{
		{Hash[String, Integer]{}, Boolean(true)},
		{Hash[String, Integer]{"a": 1}, Boolean(false)},
		{Hash[String, Integer]{"a": 1, "b": 2}, Boolean(false)},
	}

	for _, test := range tests {
		result := test.hash.IsEmpty()
		if result != test.expected {
			t.Errorf("IsEmpty() for %v expected %t, got %t", test.hash, test.expected, result)
		}
	}
}

func TestHash_Merge(t *testing.T) {
	hash1 := Hash[String, Integer]{"a": 1, "b": 2}
	hash2 := Hash[String, Integer]{"b": 3, "c": 4}

	result := hash1.Merge(hash2)

	// Check that all keys are present
	expected := map[string]int{"a": 1, "b": 3, "c": 4}
	if len(result) != len(expected) {
		t.Errorf("Merge() expected length %d, got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[String(key)] != Integer(value) {
			t.Errorf("Merge() for key '%s' expected %d, got %d", key, value, result[String(key)])
		}
	}

	// Check that original hash is not modified
	if hash1["b"] != 2 {
		t.Error("Merge() should not modify original hash")
	}
}

func TestHash_EnforceMerge(t *testing.T) {
	hash1 := Hash[String, Integer]{"a": 1, "b": 2}
	hash2 := Hash[String, Integer]{"b": 3, "c": 4}

	hash1.EnforceMerge(hash2)

	// Check that hash1 is modified
	expected := map[string]int{"a": 1, "b": 3, "c": 4}
	if len(hash1) != len(expected) {
		t.Errorf("EnforceMerge() expected length %d, got %d", len(expected), len(hash1))
	}

	for key, value := range expected {
		if hash1[String(key)] != Integer(value) {
			t.Errorf("EnforceMerge() for key '%s' expected %d, got %d", key, value, hash1[String(key)])
		}
	}
}

func TestHash_Select(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3, "d": 4}

	result := hash.Select(func(_ String, v Integer) bool {
		return v > 2
	})

	// Check that only values > 2 are selected
	expected := map[string]int{"c": 3, "d": 4}
	if len(result) != len(expected) {
		t.Errorf("Select() expected length %d, got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[String(key)] != Integer(value) {
			t.Errorf("Select() for key '%s' expected %d, got %d", key, value, result[String(key)])
		}
	}
}

func TestHash_Reject(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3, "d": 4}

	result := hash.Reject(func(_ String, v Integer) bool {
		return v <= 2
	})

	// Check that only values > 2 are kept (Reject returns pairs where predicate is false)
	expected := map[string]int{"c": 3, "d": 4}
	if len(result) != len(expected) {
		t.Errorf("Reject() expected length %d, got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[String(key)] != Integer(value) {
			t.Errorf("Reject() for key '%s' expected %d, got %d", key, value, result[String(key)])
		}
	}
}

func TestHash_Each(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}
	var result []String
	var values []Integer

	hash.Each(func(k String, v Integer) {
		result = append(result, k)
		values = append(values, v)
	})

	if len(result) != 3 {
		t.Errorf("Each() should process all key-value pairs, expected 3, got %d", len(result))
	}

	// Check that all keys and values are processed
	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	expectedValues := map[int]bool{1: true, 2: true, 3: true}

	for _, key := range result {
		if !expectedKeys[string(key)] {
			t.Errorf("Each() processed unexpected key: %s", key)
		}
	}

	for _, value := range values {
		if !expectedValues[int(value)] {
			t.Errorf("Each() processed unexpected value: %d", value)
		}
	}
}

func TestHash_EachKey(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}
	var result []String

	hash.EachKey(func(k String) {
		result = append(result, k)
	})

	if len(result) != 3 {
		t.Errorf("EachKey() should process all keys, expected 3, got %d", len(result))
	}

	expectedKeys := map[string]bool{"a": true, "b": true, "c": true}
	for _, key := range result {
		if !expectedKeys[string(key)] {
			t.Errorf("EachKey() processed unexpected key: %s", key)
		}
	}
}

func TestHash_EachValue(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}
	var result []Integer

	hash.EachValue(func(v Integer) {
		result = append(result, v)
	})

	if len(result) != 3 {
		t.Errorf("EachValue() should process all values, expected 3, got %d", len(result))
	}

	expectedValues := map[int]bool{1: true, 2: true, 3: true}
	for _, value := range result {
		if !expectedValues[int(value)] {
			t.Errorf("EachValue() processed unexpected value: %d", value)
		}
	}
}

func TestHash_Map(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	result := hash.Map(func(k String, v Integer) (String, Integer) {
		return k + "x", v * 2
	})

	// Check that keys and values are transformed
	expected := map[string]int{"ax": 2, "bx": 4, "cx": 6}
	if len(result) != len(expected) {
		t.Errorf("Map() expected length %d, got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[String(key)] != Integer(value) {
			t.Errorf("Map() for key '%s' expected %d, got %d", key, value, result[String(key)])
		}
	}
}

func TestHash_Invert(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	result := hash.Invert()

	// Check that keys and values are swapped
	expected := map[int]string{1: "a", 2: "b", 3: "c"}
	if len(result) != len(expected) {
		t.Errorf("Invert() expected length %d, got %d", len(expected), len(result))
	}

	for key, value := range expected {
		if result[Integer(key)] != String(value) {
			t.Errorf("Invert() for key %d expected '%s', got '%v'", key, value, result[Integer(key)])
		}
	}
}

func TestHash_ToArray(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	result := hash.ToArray()

	if len(result) != 3 {
		t.Errorf("ToArray() expected length 3, got %d", len(result))
	}

	// Check that all pairs are present
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	for _, pair := range result {
		p := pair.(Pair[String, Integer])
		key := string(p.Key)
		value := int(p.Value)

		if expectedValue, exists := expected[key]; !exists || expectedValue != value {
			t.Errorf("ToArray() contains unexpected pair: %s -> %d", key, value)
		}
	}
}

func TestHash_Clone(t *testing.T) {
	original := Hash[String, Integer]{"a": 1, "b": 2, "c": 3}

	result := original.Clone()

	// Check that clone has same content
	if len(result) != len(original) {
		t.Errorf("Clone() expected length %d, got %d", len(original), len(result))
	}

	for key, value := range original {
		if result[key] != value {
			t.Errorf("Clone() for key '%s' expected %d, got %d", key, value, result[key])
		}
	}

	// Check that clone is independent
	result["d"] = 4
	if original["d"] == 4 {
		t.Error("Clone() should create independent copy")
	}
}

func TestHash_Update(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2}
	other := Hash[String, Integer]{"b": 3, "c": 4}

	hash.Update(other)

	// Check that hash is updated
	expected := map[string]int{"a": 1, "b": 3, "c": 4}
	if len(hash) != len(expected) {
		t.Errorf("Update() expected length %d, got %d", len(expected), len(hash))
	}

	for key, value := range expected {
		if hash[String(key)] != Integer(value) {
			t.Errorf("Update() for key '%s' expected %d, got %d", key, value, hash[String(key)])
		}
	}
}

func TestHash_Replace(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2}
	other := Hash[String, Integer]{"c": 3, "d": 4}

	hash.Replace(other)

	// Check that hash is completely replaced
	expected := map[string]int{"c": 3, "d": 4}
	if len(hash) != len(expected) {
		t.Errorf("Replace() expected length %d, got %d", len(expected), len(hash))
	}

	for key, value := range expected {
		if hash[String(key)] != Integer(value) {
			t.Errorf("Replace() for key '%s' expected %d, got %d", key, value, hash[String(key)])
		}
	}

	// Check that old keys are gone
	if hash.HasKey(String("a")) || hash.HasKey(String("b")) {
		t.Error("Replace() should remove old keys")
	}
}

func TestHash_KeepIf(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3, "d": 4}

	hash.KeepIf(func(_ String, v Integer) bool {
		return v > 2
	})

	// Check that only values > 2 are kept
	expected := map[string]int{"c": 3, "d": 4}
	if len(hash) != len(expected) {
		t.Errorf("KeepIf() expected length %d, got %d", len(expected), len(hash))
	}

	for key, value := range expected {
		if hash[String(key)] != Integer(value) {
			t.Errorf("KeepIf() for key '%s' expected %d, got %d", key, value, hash[String(key)])
		}
	}
}

func TestHash_DeleteIf(t *testing.T) {
	hash := Hash[String, Integer]{"a": 1, "b": 2, "c": 3, "d": 4}

	hash.DeleteIf(func(_ String, v Integer) bool {
		return v <= 2
	})

	// Check that only values > 2 are kept (DeleteIf deletes entries where predicate is true)
	expected := map[string]int{"c": 3, "d": 4}
	if len(hash) != len(expected) {
		t.Errorf("DeleteIf() expected length %d, got %d", len(expected), len(hash))
	}

	for key, value := range expected {
		if hash[String(key)] != Integer(value) {
			t.Errorf("DeleteIf() for key '%s' expected %d, got %d", key, value, hash[String(key)])
		}
	}
}
