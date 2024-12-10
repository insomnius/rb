package rb

import (
	"strings"
)

// String is a custom string type to emulate Ruby-like behavior.
type String string

// Chars splits the String into an Array of single-character Strings.
// Example: String("hello").Chars() -> ["h", "e", "l", "l", "o"]
func (s String) Chars() Array[String] {
	chars := make([]String, len(s))
	for k, c := range s {
		chars[k] = String(c)
	}
	return chars
}

// Length returns the length of the String as an Integer.
// Example: String("hello").Length() -> 5
func (s String) Length() Integer {
	return Integer(len(s))
}

// ToS returns the String itself, mimicking Ruby's to_s method.
// Example: String("hello").ToS() -> "hello"
func (s String) ToS() String {
	return s
}

// ToStr is an alias for ToS, returning the String itself.
// Example: String("hello").ToStr() -> "hello"
func (s String) ToStr() String {
	return s.ToS()
}

// Downcase returns a new String with all characters converted to lowercase.
// Example: String("Hello").Downcase() -> "hello"
func (s String) Downcase() String {
	return String(strings.ToLower(string(s)))
}

// EnforceDowncase converts the String to lowercase in place and returns it.
// Example:
// str := String("Hello")
// str.EnforceDowncase() // str is now "hello"
func (s *String) EnforceDowncase() String {
	*s = s.Downcase()
	return *s
}

// Upcase returns a new String with all characters converted to uppercase.
// Example: String("hello").Upcase() -> "HELLO"
func (s String) Upcase() String {
	return String(strings.ToUpper(string(s)))
}

// UpcaseBang converts the String to uppercase in place.
// Example:
// str := String("hello")
// str.UpcaseBang() // str is now "HELLO"
func (s *String) UpcaseBang() {
	*s = s.Upcase()
}
