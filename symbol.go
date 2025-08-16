package rb

import (
	"strconv"
	"strings"
	"unicode"
)

// Symbol represents a Ruby-like symbol, implemented as a string with special behavior.
type Symbol string

// NewSymbol creates a new Symbol from a string.
// Example: NewSymbol("hello") -> :hello
func NewSymbol(s string) Symbol {
	return Symbol(s)
}

// ToS converts the Symbol to a String.
// Example: Symbol("hello").ToS() -> "hello"
func (s Symbol) ToS() String {
	return String(s)
}

// ToStr is an alias for ToS.
func (s Symbol) ToStr() String {
	return s.ToS()
}

// ToI attempts to convert the Symbol to an Integer.
// Returns 0 if the conversion fails.
// Example: Symbol("123").ToI() -> 123
func (s Symbol) ToI() Integer {
	if i, err := strconv.Atoi(string(s)); err == nil {
		return Integer(i)
	}
	return 0
}

// ToF attempts to convert the Symbol to a Float.
// Returns 0.0 if the conversion fails.
// Example: Symbol("123.45").ToF() -> 123.45
func (s Symbol) ToF() Float {
	if f, err := strconv.ParseFloat(string(s), 64); err == nil {
		return Float(f)
	}
	return 0.0
}

// ToSym returns the Symbol itself (identity method).
// Example: Symbol("hello").ToSym() -> :hello
func (s Symbol) ToSym() Symbol {
	return s
}

// IsEmpty checks if the Symbol is empty.
// Example: Symbol("").IsEmpty() -> true
func (s Symbol) IsEmpty() Boolean {
	return Boolean(len(s) == 0)
}

// Length returns the length of the Symbol.
// Example: Symbol("hello").Length() -> 5
func (s Symbol) Length() Integer {
	return Integer(len(s))
}

// Size is an alias for Length.
func (s Symbol) Size() Integer {
	return s.Length()
}

// IsBlank checks if the Symbol is blank (empty or contains only whitespace).
// Example: Symbol("  ").IsBlank() -> true
func (s Symbol) IsBlank() Boolean {
	return Boolean(strings.TrimSpace(string(s)) == "")
}

// IsPresent checks if the Symbol is present (not blank).
// Example: Symbol("hello").IsPresent() -> true
func (s Symbol) IsPresent() Boolean {
	return Boolean(!s.IsBlank())
}

// Upcase returns a new Symbol with all characters converted to uppercase.
// Example: Symbol("hello").Upcase() -> "HELLO"
func (s Symbol) Upcase() Symbol {
	return Symbol(strings.ToUpper(string(s)))
}

// Downcase returns a new Symbol with all characters converted to lowercase.
// Example: Symbol("HELLO").Downcase() -> "hello"
func (s Symbol) Downcase() Symbol {
	return Symbol(strings.ToLower(string(s)))
}

// Capitalize returns a new Symbol with the first character capitalized and the rest lowercase.
// Example: Symbol("hello world").Capitalize() -> "Hello world"
func (s Symbol) Capitalize() Symbol {
	if len(s) == 0 {
		return s
	}
	runes := []rune(string(s))
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return Symbol(string(runes))
}

// Swapcase returns a new Symbol with uppercase letters converted to lowercase and vice versa.
// Example: Symbol("Hello World").Swapcase() -> "hELLO wORLD"
func (s Symbol) Swapcase() Symbol {
	runes := []rune(string(s))
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return Symbol(string(runes))
}

// Title returns a new Symbol with the first character of each word capitalized.
// Example: Symbol("hello world").Title() -> "Hello World"
func (s Symbol) Title() Symbol {
	return Symbol(strings.Title(strings.ToLower(string(s))))
}

// Strip returns a new Symbol with leading and trailing whitespace removed.
// Example: Symbol("  hello world  ").Strip() -> "hello world"
func (s Symbol) Strip() Symbol {
	return Symbol(strings.TrimSpace(string(s)))
}

// Lstrip returns a new Symbol with leading whitespace removed.
// Example: Symbol("  hello world").Lstrip() -> "hello world"
func (s Symbol) Lstrip() Symbol {
	return Symbol(strings.TrimLeftFunc(string(s), unicode.IsSpace))
}

// Rstrip returns a new Symbol with trailing whitespace removed.
// Example: Symbol("hello world  ").Rstrip() -> "hello world"
func (s Symbol) Rstrip() Symbol {
	return Symbol(strings.TrimRightFunc(string(s), unicode.IsSpace))
}

// Reverse returns a new Symbol with characters in reverse order.
// Example: Symbol("hello").Reverse() -> "olleh"
func (s Symbol) Reverse() Symbol {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return Symbol(string(runes))
}

// StartWith checks if the Symbol starts with the given prefix.
// Example: Symbol("hello world").StartWith("hello") -> true
func (s Symbol) StartWith(prefix Symbol) Boolean {
	return Boolean(strings.HasPrefix(string(s), string(prefix)))
}

// EndWith checks if the Symbol ends with the given suffix.
// Example: Symbol("hello world").EndWith("world") -> true
func (s Symbol) EndWith(suffix Symbol) Boolean {
	return Boolean(strings.HasSuffix(string(s), string(suffix)))
}

// Include checks if the Symbol contains the given substring.
// Example: Symbol("hello world").Include("world") -> true
func (s Symbol) Include(substr Symbol) Boolean {
	return Boolean(strings.Contains(string(s), string(substr)))
}

// Gsub performs global substitution, replacing all occurrences of a pattern with a replacement.
// Example: Symbol("hello world").Gsub("o", "0") -> "hell0 w0rld"
func (s Symbol) Gsub(pattern, replacement Symbol) Symbol {
	return Symbol(strings.ReplaceAll(string(s), string(pattern), string(replacement)))
}

// Sub performs substitution, replacing the first occurrence of a pattern with a replacement.
// Example: Symbol("hello world").Sub("o", "0") -> "hell0 world"
func (s Symbol) Sub(pattern, replacement Symbol) Symbol {
	return Symbol(strings.Replace(string(s), string(pattern), string(replacement), 1))
}

// Split splits the Symbol by the given separator.
// Example: Symbol("hello,world").Split(",") -> ["hello", "world"]
func (s Symbol) Split(sep Symbol) Array[Symbol] {
	parts := strings.Split(string(s), string(sep))
	result := make(Array[Symbol], len(parts))
	for i, part := range parts {
		result[i] = Symbol(part)
	}
	return result
}

// Lines splits the Symbol into an Array of lines.
// Example: Symbol("hello\nworld").Lines() -> ["hello", "world"]
func (s Symbol) Lines() Array[Symbol] {
	lines := strings.Split(string(s), "\n")
	result := make(Array[Symbol], len(lines))
	for i, line := range lines {
		result[i] = Symbol(line)
	}
	return result
}

// Words splits the Symbol into an Array of words (splitting on whitespace).
// Example: Symbol("hello world").Words() -> ["hello", "world"]
func (s Symbol) Words() Array[Symbol] {
	words := strings.Fields(string(s))
	result := make(Array[Symbol], len(words))
	for i, word := range words {
		result[i] = Symbol(word)
	}
	return result
}

// Chars splits the Symbol into an Array of single-character Symbols.
// Example: Symbol("hello").Chars() -> ["h", "e", "l", "l", "o"]
func (s Symbol) Chars() Array[Symbol] {
	chars := make([]Symbol, len(s))
	for i, c := range s {
		chars[i] = Symbol(c)
	}
	return Array[Symbol](chars)
}

// IsNumeric checks if the Symbol represents a numeric value.
// Example: Symbol("123").IsNumeric() -> true
func (s Symbol) IsNumeric() Boolean {
	_, err := strconv.ParseFloat(string(s), 64)
	return Boolean(err == nil)
}

// IsInteger checks if the Symbol represents an integer value.
// Example: Symbol("123").IsInteger() -> true
func (s Symbol) IsInteger() Boolean {
	_, err := strconv.Atoi(string(s))
	return Boolean(err == nil)
}

// IsFloat checks if the Symbol represents a float value.
// Example: Symbol("123.45").IsFloat() -> true
func (s Symbol) IsFloat() Boolean {
	_, err := strconv.ParseFloat(string(s), 64)
	return Boolean(err == nil)
}

// IsAlpha checks if the Symbol contains only alphabetic characters.
// Example: Symbol("hello").IsAlpha() -> true
func (s Symbol) IsAlpha() Boolean {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return Boolean(len(s) > 0)
}

// IsAlphanumeric checks if the Symbol contains only alphanumeric characters.
// Example: Symbol("hello123").IsAlphanumeric() -> true
func (s Symbol) IsAlphanumeric() Boolean {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return Boolean(len(s) > 0)
}

// IsDigit checks if the Symbol contains only digit characters.
// Example: Symbol("123").IsDigit() -> true
func (s Symbol) IsDigit() Boolean {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return Boolean(len(s) > 0)
}

// IsSpace checks if the Symbol contains only whitespace characters.
// Example: Symbol("  ").IsSpace() -> true
func (s Symbol) IsSpace() Boolean {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return Boolean(len(s) > 0)
}

// IsUpper checks if the Symbol contains only uppercase characters.
// Example: Symbol("HELLO").IsUpper() -> true
func (s Symbol) IsUpper() Boolean {
	for _, r := range s {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return Boolean(len(s) > 0)
}

// IsLower checks if the Symbol contains only lowercase characters.
// Example: Symbol("hello").IsLower() -> true
func (s Symbol) IsLower() Boolean {
	for _, r := range s {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return Boolean(len(s) > 0)
}

// Clone returns a copy of the Symbol.
// Example: Symbol("hello").Clone() -> :hello
func (s Symbol) Clone() Symbol {
	return s
}

// String returns the string representation of the Symbol.
// This implements the fmt.Stringer interface.
func (s Symbol) String() string {
	return string(s)
}