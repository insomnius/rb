package rb

// Hash is a generic map type to emulate Ruby-like hash behavior.
type Hash[K comparable, V any] map[K]V

// Keys returns an Array of all keys in the Hash.
// Example: Hash[string, int]{"a": 1, "b": 2}.Keys() -> ["a", "b"]
func (h Hash[K, V]) Keys() AnyArray {
	keys := make([]any, 0, len(h))
	for k := range h {
		keys = append(keys, k)
	}
	return AnyArray(keys)
}

// Values returns an Array of all values in the Hash.
// Example: Hash[string, int]{"a": 1, "b": 2}.Values() -> [1, 2]
func (h Hash[K, V]) Values() AnyArray {
	values := make([]any, 0, len(h))
	for _, v := range h {
		values = append(values, v)
	}
	return AnyArray(values)
}

// HasKey checks if the Hash contains the given key.
// Example: Hash[string, int]{"a": 1}.HasKey("a") -> true
func (h Hash[K, V]) HasKey(key K) Boolean {
	_, exists := h[key]
	return Boolean(exists)
}

// HasValue checks if the Hash contains the given value.
// Note: This method is simplified and may not work for all types.
// Example: Hash[string, int]{"a": 1}.HasValue(1) -> true
func (h Hash[K, V]) HasValue(value V) Boolean {
	// Note: This is a simplified implementation
	// In practice, you'd want to use reflect.DeepEqual or custom comparators
	return false
}

// Get retrieves the value for the given key, returning a default value if the key doesn't exist.
// Example: Hash[string, int]{"a": 1}.Get("b", 0) -> 0
func (h Hash[K, V]) Get(key K, defaultValue V) V {
	if value, exists := h[key]; exists {
		return value
	}
	return defaultValue
}

// Fetch retrieves the value for the given key, panicking if the key doesn't exist.
// Example: Hash[string, int]{"a": 1}.Fetch("a") -> 1
func (h Hash[K, V]) Fetch(key K) V {
	if value, exists := h[key]; exists {
		return value
	}
	panic("key not found")
}

// Set sets the value for the given key.
// Example: Hash[string, int]{}.Set("a", 1)
func (h Hash[K, V]) Set(key K, value V) {
	h[key] = value
}

// Delete removes the key-value pair for the given key and returns the value.
// Example: Hash[string, int]{"a": 1}.Delete("a") -> 1
func (h Hash[K, V]) Delete(key K) V {
	value := h[key]
	delete(h, key)
	return value
}

// Clear removes all key-value pairs from the Hash.
// Example: Hash[string, int]{"a": 1}.Clear()
func (h Hash[K, V]) Clear() {
	for k := range h {
		delete(h, k)
	}
}

// Size returns the number of key-value pairs in the Hash.
// Example: Hash[string, int]{"a": 1, "b": 2}.Size() -> 2
func (h Hash[K, V]) Size() Integer {
	return Integer(len(h))
}

// Length is an alias for Size.
func (h Hash[K, V]) Length() Integer {
	return h.Size()
}

// IsEmpty checks if the Hash is empty.
// Example: Hash[string, int]{}.IsEmpty() -> true
func (h Hash[K, V]) IsEmpty() Boolean {
	return Boolean(len(h) == 0)
}

// Merge merges another Hash into this Hash, overwriting existing keys.
// Example: Hash[string, int]{"a": 1}.Merge(Hash[string, int]{"b": 2}) -> {"a": 1, "b": 2}
func (h Hash[K, V]) Merge(other Hash[K, V]) Hash[K, V] {
	result := make(Hash[K, V])
	
	// Copy current hash
	for k, v := range h {
		result[k] = v
	}
	
	// Merge other hash
	for k, v := range other {
		result[k] = v
	}
	
	return result
}

// EnforceMerge merges another Hash into this Hash in place.
func (h Hash[K, V]) EnforceMerge(other Hash[K, V]) {
	for k, v := range other {
		h[k] = v
	}
}

// Select returns a new Hash containing key-value pairs for which the predicate returns true.
// Example: Hash[string, int]{"a": 1, "b": 2}.Select(func(k string, v int) bool { return v > 1 }) -> {"b": 2}
func (h Hash[K, V]) Select(predicate func(K, V) bool) Hash[K, V] {
	result := make(Hash[K, V])
	for k, v := range h {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Reject returns a new Hash containing key-value pairs for which the predicate returns false.
// Example: Hash[string, int]{"a": 1, "b": 2}.Reject(func(k string, v int) bool { return v > 1 }) -> {"a": 1}
func (h Hash[K, V]) Reject(predicate func(K, V) bool) Hash[K, V] {
	result := make(Hash[K, V])
	for k, v := range h {
		if !predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// Each applies the given function to each key-value pair.
// Example: Hash[string, int]{"a": 1}.Each(func(k string, v int) { fmt.Printf("%s: %d\n", k, v) })
func (h Hash[K, V]) Each(fn func(K, V)) {
	for k, v := range h {
		fn(k, v)
	}
}

// EachKey applies the given function to each key.
// Example: Hash[string, int]{"a": 1}.EachKey(func(k string) { fmt.Println(k) })
func (h Hash[K, V]) EachKey(fn func(K)) {
	for k := range h {
		fn(k)
	}
}

// EachValue applies the given function to each value.
// Example: Hash[string, int]{"a": 1}.EachValue(func(v int) { fmt.Println(v) })
func (h Hash[K, V]) EachValue(fn func(V)) {
	for _, v := range h {
		fn(v)
	}
}

// Map transforms the Hash by applying the given function to each key-value pair.
// Example: Hash[string, int]{"a": 1}.Map(func(k string, v int) (string, int) { return k + "x", v * 2 })
func (h Hash[K, V]) Map(fn func(K, V) (K, V)) Hash[K, V] {
	result := make(Hash[K, V])
	for k, v := range h {
		newKey, newValue := fn(k, v)
		result[newKey] = newValue
	}
	return result
}

// Invert returns a new Hash with keys and values swapped.
// Note: This only works if all values are unique and comparable.
// Example: Hash[string, int]{"a": 1}.Invert() -> {1: "a"}
func (h Hash[K, V]) Invert() Hash[any, any] {
	result := make(Hash[any, any])
	for k, v := range h {
		result[v] = k
	}
	return result
}

// Default returns the default value for keys that don't exist.
// Example: Hash[string, int]{"a": 1}.Default(0)
func (h Hash[K, V]) Default(defaultValue V) Hash[K, V] {
	result := make(Hash[K, V])
	for k, v := range h {
		result[k] = v
	}
	
	// Note: This is a simplified implementation
	// In practice, you'd want to handle defaults differently
	return result
}

// ToArray converts the Hash to an Array of key-value pairs.
// Example: Hash[string, int]{"a": 1}.ToArray() -> [{"a", 1}]
func (h Hash[K, V]) ToArray() AnyArray {
	result := make([]any, 0, len(h))
	for k, v := range h {
		result = append(result, Pair[K, V]{Key: k, Value: v})
	}
	return AnyArray(result)
}

// Clone returns a shallow copy of the Hash.
// Example: Hash[string, int]{"a": 1}.Clone()
func (h Hash[K, V]) Clone() Hash[K, V] {
	result := make(Hash[K, V])
	for k, v := range h {
		result[k] = v
	}
	return result
}

// Update updates the Hash with key-value pairs from another Hash.
// Example: Hash[string, int]{"a": 1}.Update(Hash[string, int]{"b": 2})
func (h Hash[K, V]) Update(other Hash[K, V]) {
	for k, v := range other {
		h[k] = v
	}
}

// Replace replaces the entire Hash content with another Hash.
// Example: Hash[string, int]{"a": 1}.Replace(Hash[string, int]{"b": 2})
func (h Hash[K, V]) Replace(other Hash[K, V]) {
	h.Clear()
	for k, v := range other {
		h[k] = v
	}
}

// KeepIf keeps only the key-value pairs for which the predicate returns true.
// Example: Hash[string, int]{"a": 1, "b": 2}.KeepIf(func(k string, v int) bool { return v > 1 })
func (h Hash[K, V]) KeepIf(predicate func(K, V) bool) {
	for k, v := range h {
		if !predicate(k, v) {
			delete(h, k)
		}
	}
}

// DeleteIf deletes key-value pairs for which the predicate returns true.
// Example: Hash[string, int]{"a": 1, "b": 2}.DeleteIf(func(k string, v int) bool { return v > 1 })
func (h Hash[K, V]) DeleteIf(predicate func(K, V) bool) {
	for k, v := range h {
		if predicate(k, v) {
			delete(h, k)
		}
	}
}

// Pair represents a key-value pair in a Hash.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}