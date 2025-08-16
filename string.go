package rb

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// String is a custom string type to emulate Ruby-like behavior.
type String string

// Chars splits the String into an Array of single-character Strings.
// Example: String("hello").Chars() -> ["h", "e", "l", "l", "o"]
func (s String) Chars() Array[String] {
	runes := []rune(string(s))
	chars := make([]String, len(runes))
	for k, c := range runes {
		chars[k] = String(c)
	}
	return chars
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

// Length returns the length of the String as an Integer.
// Example: String("hello").Length() -> 5
func (s String) Length() Integer {
	return Integer(utf8.RuneCountInString(string(s)))
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

func (s String) Split(sep String) Array[String] {
	splittedString := strings.Split(string(s), string(sep))
	arr := make(Array[String], len(splittedString))
	for k, v := range splittedString {
		arr[k] = String(v)
	}
	return arr
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

// Capitalize returns a new String with the first character capitalized and the rest lowercase.
// Example: String("hello world").Capitalize() -> "Hello world"
func (s String) Capitalize() String {
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
	return String(string(runes))
}

// EnforceCapitalize capitalizes the String in place and returns it.
func (s *String) EnforceCapitalize() String {
	*s = s.Capitalize()
	return *s
}

// Strip returns a new String with leading and trailing whitespace removed.
// Example: String("  hello world  ").Strip() -> "hello world"
func (s String) Strip() String {
	return String(strings.TrimSpace(string(s)))
}

// EnforceStrip removes leading and trailing whitespace in place and returns it.
func (s *String) EnforceStrip() String {
	*s = s.Strip()
	return *s
}

// Reverse returns a new String with characters in reverse order.
// Example: String("hello").Reverse() -> "olleh"
func (s String) Reverse() String {
	runes := []rune(string(s))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return String(string(runes))
}

// EnforceReverse reverses the String in place and returns it.
func (s *String) EnforceReverse() String {
	*s = s.Reverse()
	return *s
}

// IsEmpty checks if the String is empty (length 0).
// Example: String("").IsEmpty() -> true
func (s String) IsEmpty() Boolean {
	return Boolean(len(s) == 0)
}

// IsBlank checks if the String is blank (empty or contains only whitespace).
// Example: String("  ").IsBlank() -> true
func (s String) IsBlank() Boolean {
	return Boolean(strings.TrimSpace(string(s)) == "")
}

// StartWith checks if the String starts with the given prefix.
// Example: String("hello world").StartWith("hello") -> true
func (s String) StartWith(prefix String) Boolean {
	return Boolean(strings.HasPrefix(string(s), string(prefix)))
}

// EndWith checks if the String ends with the given suffix.
// Example: String("hello world").EndWith("world") -> true
func (s String) EndWith(suffix String) Boolean {
	return Boolean(strings.HasSuffix(string(s), string(suffix)))
}

// Include checks if the String contains the given substring.
// Example: String("hello world").Include("world") -> true
func (s String) Include(substr String) Boolean {
	return Boolean(strings.Contains(string(s), string(substr)))
}

// Gsub performs global substitution, replacing all occurrences of a pattern with a replacement.
// Example: String("hello world").Gsub("o", "0") -> "hell0 w0rld"
func (s String) Gsub(pattern, replacement String) String {
	// Special case: if both string and pattern are empty, return empty string
	if len(s) == 0 && len(pattern) == 0 {
		return String("")
	}
	return String(strings.ReplaceAll(string(s), string(pattern), string(replacement)))
}

// EnforceGsub performs global substitution in place and returns it.
func (s *String) EnforceGsub(pattern, replacement String) String {
	*s = s.Gsub(pattern, replacement)
	return *s
}

// Sub performs substitution, replacing the first occurrence of a pattern with a replacement.
// Example: String("hello world").Sub("o", "0") -> "hell0 world"
func (s String) Sub(pattern, replacement String) String {
	// Special case: if both string and pattern are empty, return empty string
	if len(s) == 0 && len(pattern) == 0 {
		return String("")
	}
	return String(strings.Replace(string(s), string(pattern), string(replacement), 1))
}

// EnforceSub performs substitution in place and returns it.
func (s *String) EnforceSub(pattern, replacement String) String {
	*s = s.Sub(pattern, replacement)
	return *s
}

// Lstrip returns a new String with leading whitespace removed.
// Example: String("  hello world").Lstrip() -> "hello world"
func (s String) Lstrip() String {
	return String(strings.TrimLeftFunc(string(s), unicode.IsSpace))
}

// EnforceLstrip removes leading whitespace in place and returns it.
func (s *String) EnforceLstrip() String {
	*s = s.Lstrip()
	return *s
}

// Rstrip returns a new String with trailing whitespace removed.
// Example: String("hello world  ").Rstrip() -> "hello world"
func (s String) Rstrip() String {
	return String(strings.TrimRightFunc(string(s), unicode.IsSpace))
}

// EnforceRstrip removes trailing whitespace in place and returns it.
func (s *String) EnforceRstrip() String {
	*s = s.Rstrip()
	return *s
}

// Swapcase returns a new String with uppercase letters converted to lowercase and vice versa.
// Example: String("Hello World").Swapcase() -> "hELLO wORLD"
func (s String) Swapcase() String {
	runes := []rune(string(s))
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return String(string(runes))
}

// EnforceSwapcase swaps case in place and returns it.
func (s *String) EnforceSwapcase() String {
	*s = s.Swapcase()
	return *s
}

// Title returns a new String with the first character of each word capitalized.
// Example: String("hello world").Title() -> "Hello World"
func (s String) Title() String {
	words := strings.Fields(string(s))
	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		runes := []rune(word)
		for j := 0; j < len(word); j++ {
			if j == 0 {
				runes[j] = unicode.ToUpper(runes[j])
			} else {
				runes[j] = unicode.ToLower(runes[j])
			}
		}
		words[i] = string(runes)
	}
	return String(strings.Join(words, " "))
}

// EnforceTitle capitalizes the first character of each word in place and returns it.
func (s *String) EnforceTitle() String {
	*s = s.Title()
	return *s
}

// Lines splits the String into an Array of lines.
// Example: String("hello\nworld").Lines() -> ["hello", "world"]
func (s String) Lines() Array[String] {
	lines := strings.Split(string(s), "\n")
	arr := make(Array[String], len(lines))
	for k, v := range lines {
		arr[k] = String(v)
	}
	return arr
}

// Words splits the String into an Array of words (splitting on whitespace).
// Example: String("hello world").Words() -> ["hello", "world"]
func (s String) Words() Array[String] {
	words := strings.Fields(string(s))
	arr := make(Array[String], len(words))
	for k, v := range words {
		arr[k] = String(v)
	}
	return arr
}

// ToI attempts to convert the String to an Integer.
// Returns 0 if the conversion fails.
// Example: String("123").ToI() -> 123
func (s String) ToI() Integer {
	// Simple integer conversion - in a real implementation you might want more robust parsing
	var result int
	_, err := fmt.Sscanf(string(s), "%d", &result)
	if err != nil {
		return 0
	}
	return Integer(result)
}

// ToF attempts to convert the String to a Float.
// Returns 0.0 if the conversion fails.
// Example: String("123.45").ToF() -> 123.45
func (s String) ToF() Float {
	var result float64
	_, err := fmt.Sscanf(string(s), "%f", &result)
	if err != nil {
		return 0.0
	}
	return Float(result)
}
