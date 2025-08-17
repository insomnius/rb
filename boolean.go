// Package rb provides Ruby-inspired utility methods for Go types.
package rb

// Boolean represents a boolean value with Ruby-inspired methods.
type Boolean bool

// ToS converts the Boolean to a String representation.
func (b Boolean) ToS() String {
	if b {
		return "true"
	}
	return "false"
}

// ToStr is an alias for ToS.
func (b Boolean) ToStr() String {
	return b.ToS()
}

// And performs logical AND operation with another Boolean.
// Example: Boolean(true).And(Boolean(false)) -> false
func (b Boolean) And(other Boolean) Boolean {
	return Boolean(bool(b) && bool(other))
}

// Or performs logical OR operation with another Boolean.
// Example: Boolean(true).Or(Boolean(false)) -> true
func (b Boolean) Or(other Boolean) Boolean {
	return Boolean(bool(b) || bool(other))
}

// Not performs logical NOT operation.
// Example: Boolean(true).Not() -> false
func (b Boolean) Not() Boolean {
	return Boolean(!bool(b))
}

// Xor performs logical XOR operation with another Boolean.
// Example: Boolean(true).Xor(Boolean(true)) -> false
func (b Boolean) Xor(other Boolean) Boolean {
	return Boolean(bool(b) != bool(other))
}

// Nand performs logical NAND operation.
func (b Boolean) Nand(other Boolean) Boolean {
	return Boolean(!bool(b) || !bool(other))
}

// Nor performs logical NOR operation.
func (b Boolean) Nor(other Boolean) Boolean {
	return Boolean(!bool(b) && !bool(other))
}

// Xnor performs logical XNOR operation with another Boolean.
// Example: Boolean(true).Xnor(Boolean(true)) -> true
func (b Boolean) Xnor(other Boolean) Boolean {
	return Boolean(bool(b) == bool(other))
}

// Implies performs logical implication (b implies other).
// Example: Boolean(false).Implies(Boolean(true)) -> true
func (b Boolean) Implies(other Boolean) Boolean {
	return Boolean(!bool(b) || bool(other))
}

// ToI converts the Boolean to an Integer (true -> 1, false -> 0).
// Example: Boolean(true).ToI() -> 1
func (b Boolean) ToI() Integer {
	if b {
		return 1
	}
	return 0
}

// ToF converts the Boolean to a Float (true -> 1.0, false -> 0.0).
// Example: Boolean(true).ToF() -> 1.0
func (b Boolean) ToF() Float {
	if b {
		return 1.0
	}
	return 0.0
}

// IsTrue checks if the Boolean is true.
// Example: Boolean(true).IsTrue() -> true
func (b Boolean) IsTrue() Boolean {
	return b
}

// IsFalse checks if the Boolean is false.
// Example: Boolean(false).IsFalse() -> true
func (b Boolean) IsFalse() Boolean {
	return Boolean(!bool(b))
}

// IfTrue executes the given function if the Boolean is true.
// Example: Boolean(true).IfTrue(func() { fmt.Println("It's true!") })
func (b Boolean) IfTrue(fn func()) {
	if b {
		fn()
	}
}

// IfFalse executes the given function if the Boolean is false.
// Example: Boolean(false).IfFalse(func() { fmt.Println("It's false!") })
func (b Boolean) IfFalse(fn func()) {
	if !b {
		fn()
	}
}

// If executes the appropriate function based on the Boolean value.
// Example: Boolean(true).If(func() { fmt.Println("True") }, func() { fmt.Println("False") })
func (b Boolean) If(ifTrue, ifFalse func()) {
	if b {
		ifTrue()
	} else {
		ifFalse()
	}
}

// Ternary returns the first value if true, second value if false.
// Example: Boolean(true).Ternary("yes", "no") -> "yes"
func (b Boolean) Ternary(ifTrue, ifFalse any) any {
	if b {
		return ifTrue
	}
	return ifFalse
}
