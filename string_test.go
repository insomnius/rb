package rb

import (
	"testing"
)

func TestString_Chars(t *testing.T) {
	tests := []struct {
		input    String
		expected Array[String]
	}{
		{String("hello"), Array[String]{"h", "e", "l", "l", "o"}},
		{String(""), Array[String]{}},
		{String("a"), Array[String]{"a"}},
		{String("世界"), Array[String]{"世", "界"}},
	}

	for _, test := range tests {
		result := test.input.Chars()
		if len(result) != len(test.expected) {
			t.Errorf("Chars() for '%s' expected length %d, got %d", test.input, len(test.expected), len(result))
		}
		for i, char := range result {
			if char != test.expected[i] {
				t.Errorf("Chars() for '%s' at index %d expected '%s', got '%s'", test.input, i, test.expected[i], char)
			}
		}
	}
}

func TestString_Downcase(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("Hello"), String("hello")},
		{String("WORLD"), String("world")},
		{String("MiXeD"), String("mixed")},
		{String(""), String("")},
		{String("123"), String("123")},
	}

	for _, test := range tests {
		result := test.input.Downcase()
		if result != test.expected {
			t.Errorf("Downcase() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceDowncase(t *testing.T) {
	str := String("Hello")
	result := str.EnforceDowncase()

	if result != String("hello") {
		t.Errorf("EnforceDowncase() expected 'hello', got '%s'", result)
	}

	if str != String("hello") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Length(t *testing.T) {
	tests := []struct {
		input    String
		expected Integer
	}{
		{String("hello"), Integer(5)},
		{String(""), Integer(0)},
		{String("世界"), Integer(2)},
		{String("a"), Integer(1)},
	}

	for _, test := range tests {
		result := test.input.Length()
		if result != test.expected {
			t.Errorf("Length() for '%s' expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestString_ToS(t *testing.T) {
	input := String("hello")
	result := input.ToS()

	if result != input {
		t.Errorf("ToS() expected '%s', got '%s'", input, result)
	}
}

func TestString_ToStr(t *testing.T) {
	input := String("hello")
	result := input.ToStr()

	if result != input {
		t.Errorf("ToStr() expected '%s', got '%s'", input, result)
	}
}

func TestString_Split(t *testing.T) {
	tests := []struct {
		input    String
		sep      String
		expected Array[String]
	}{
		{String("hello,world"), String(","), Array[String]{"hello", "world"}},
		{String("a-b-c"), String("-"), Array[String]{"a", "b", "c"}},
		{String("hello"), String(""), Array[String]{"h", "e", "l", "l", "o"}},
		{String(""), String(","), Array[String]{""}},
	}

	for _, test := range tests {
		result := test.input.Split(test.sep)
		if len(result) != len(test.expected) {
			t.Errorf("Split() for '%s' with separator '%s' expected length %d, got %d", test.input, test.sep, len(test.expected), len(result))
		}
		for i, part := range result {
			if part != test.expected[i] {
				t.Errorf("Split() for '%s' with separator '%s' at index %d expected '%s', got '%s'", test.input, test.sep, i, test.expected[i], part)
			}
		}
	}
}

func TestString_Upcase(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("hello"), String("HELLO")},
		{String("world"), String("WORLD")},
		{String("MiXeD"), String("MIXED")},
		{String(""), String("")},
		{String("123"), String("123")},
	}

	for _, test := range tests {
		result := test.input.Upcase()
		if result != test.expected {
			t.Errorf("Upcase() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_UpcaseBang(t *testing.T) {
	str := String("hello")
	str.UpcaseBang()

	if str != String("HELLO") {
		t.Errorf("UpcaseBang() expected 'HELLO', got '%s'", str)
	}
}

func TestString_Capitalize(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("hello"), String("Hello")},
		{String("WORLD"), String("World")},
		{String("hello world"), String("Hello world")},
		{String(""), String("")},
		{String("a"), String("A")},
	}

	for _, test := range tests {
		result := test.input.Capitalize()
		if result != test.expected {
			t.Errorf("Capitalize() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceCapitalize(t *testing.T) {
	str := String("hello")
	result := str.EnforceCapitalize()

	if result != String("Hello") {
		t.Errorf("EnforceCapitalize() expected 'Hello', got '%s'", result)
	}

	if str != String("Hello") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Strip(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("  hello  "), String("hello")},
		{String("hello"), String("hello")},
		{String("  "), String("")},
		{String(""), String("")},
		{String("\t\nhello\t\n"), String("hello")},
	}

	for _, test := range tests {
		result := test.input.Strip()
		if result != test.expected {
			t.Errorf("Strip() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceStrip(t *testing.T) {
	str := String("  hello  ")
	result := str.EnforceStrip()

	if result != String("hello") {
		t.Errorf("EnforceStrip() expected 'hello', got '%s'", result)
	}

	if str != String("hello") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Reverse(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("hello"), String("olleh")},
		{String(""), String("")},
		{String("a"), String("a")},
		{String("ab"), String("ba")},
		{String("世界"), String("界世")},
	}

	for _, test := range tests {
		result := test.input.Reverse()
		if result != test.expected {
			t.Errorf("Reverse() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceReverse(t *testing.T) {
	str := String("hello")
	result := str.EnforceReverse()

	if result != String("olleh") {
		t.Errorf("EnforceReverse() expected 'olleh', got '%s'", result)
	}

	if str != String("olleh") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_IsEmpty(t *testing.T) {
	tests := []struct {
		input    String
		expected Boolean
	}{
		{String(""), Boolean(true)},
		{String("hello"), Boolean(false)},
		{String(" "), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsEmpty()
		if result != test.expected {
			t.Errorf("IsEmpty() for '%s' expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestString_IsBlank(t *testing.T) {
	tests := []struct {
		input    String
		expected Boolean
	}{
		{String(""), Boolean(true)},
		{String("  "), Boolean(true)},
		{String("\t\n"), Boolean(true)},
		{String("hello"), Boolean(false)},
		{String(" hello "), Boolean(false)},
	}

	for _, test := range tests {
		result := test.input.IsBlank()
		if result != test.expected {
			t.Errorf("IsBlank() for '%s' expected %t, got %t", test.input, test.expected, result)
		}
	}
}

func TestString_StartWith(t *testing.T) {
	tests := []struct {
		input    String
		prefix   String
		expected Boolean
	}{
		{String("hello world"), String("hello"), Boolean(true)},
		{String("hello world"), String("world"), Boolean(false)},
		{String("hello"), String("hello"), Boolean(true)},
		{String(""), String(""), Boolean(true)},
		{String("hello"), String(""), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.StartWith(test.prefix)
		if result != test.expected {
			t.Errorf("StartWith() for '%s' with prefix '%s' expected %t, got %t", test.input, test.prefix, test.expected, result)
		}
	}
}

func TestString_EndWith(t *testing.T) {
	tests := []struct {
		input    String
		suffix   String
		expected Boolean
	}{
		{String("hello world"), String("world"), Boolean(true)},
		{String("hello world"), String("hello"), Boolean(false)},
		{String("hello"), String("hello"), Boolean(true)},
		{String(""), String(""), Boolean(true)},
		{String("hello"), String(""), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.EndWith(test.suffix)
		if result != test.expected {
			t.Errorf("EndWith() for '%s' with suffix '%s' expected %t, got %t", test.input, test.suffix, test.expected, result)
		}
	}
}

func TestString_Include(t *testing.T) {
	tests := []struct {
		input    String
		substr   String
		expected Boolean
	}{
		{String("hello world"), String("world"), Boolean(true)},
		{String("hello world"), String("hello"), Boolean(true)},
		{String("hello world"), String("xyz"), Boolean(false)},
		{String(""), String(""), Boolean(true)},
		{String("hello"), String(""), Boolean(true)},
	}

	for _, test := range tests {
		result := test.input.Include(test.substr)
		if result != test.expected {
			t.Errorf("Include() for '%s' with substring '%s' expected %t, got %t", test.input, test.substr, test.expected, result)
		}
	}
}

func TestString_Gsub(t *testing.T) {
	tests := []struct {
		input       String
		pattern     String
		replacement String
		expected    String
	}{
		{String("hello world"), String("o"), String("0"), String("hell0 w0rld")},
		{String("hello"), String("l"), String("L"), String("heLLo")},
		{String("hello"), String("x"), String("y"), String("hello")},
		{String(""), String(""), String("x"), String("")},
	}

	for _, test := range tests {
		result := test.input.Gsub(test.pattern, test.replacement)
		if result != test.expected {
			t.Errorf("Gsub() for '%s' with pattern '%s' and replacement '%s' expected '%s', got '%s'", test.input, test.pattern, test.replacement, test.expected, result)
		}
	}
}

func TestString_EnforceGsub(t *testing.T) {
	str := String("hello world")
	result := str.EnforceGsub("o", "0")

	if result != String("hell0 w0rld") {
		t.Errorf("EnforceGsub() expected 'hell0 w0rld', got '%s'", result)
	}

	if str != String("hell0 w0rld") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Sub(t *testing.T) {
	tests := []struct {
		input       String
		pattern     String
		replacement String
		expected    String
	}{
		{String("hello world"), String("o"), String("0"), String("hell0 world")},
		{String("hello"), String("l"), String("L"), String("heLlo")},
		{String("hello"), String("x"), String("y"), String("hello")},
		{String(""), String(""), String("x"), String("")},
	}

	for _, test := range tests {
		result := test.input.Sub(test.pattern, test.replacement)
		if result != test.expected {
			t.Errorf("Sub() for '%s' with pattern '%s' and replacement '%s' expected '%s', got '%s'", test.input, test.pattern, test.replacement, test.expected, result)
		}
	}
}

func TestString_EnforceSub(t *testing.T) {
	str := String("hello world")
	result := str.EnforceSub("o", "0")

	if result != String("hell0 world") {
		t.Errorf("EnforceSub() expected 'hell0 world', got '%s'", result)
	}

	if str != String("hell0 world") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Lstrip(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("  hello"), String("hello")},
		{String("hello"), String("hello")},
		{String("  "), String("")},
		{String(""), String("")},
		{String("\t\nhello"), String("hello")},
	}

	for _, test := range tests {
		result := test.input.Lstrip()
		if result != test.expected {
			t.Errorf("Lstrip() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceLstrip(t *testing.T) {
	str := String("  hello")
	result := str.EnforceLstrip()

	if result != String("hello") {
		t.Errorf("EnforceLstrip() expected 'hello', got '%s'", result)
	}

	if str != String("hello") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Rstrip(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("hello  "), String("hello")},
		{String("hello"), String("hello")},
		{String("  "), String("")},
		{String(""), String("")},
		{String("hello\t\n"), String("hello")},
	}

	for _, test := range tests {
		result := test.input.Rstrip()
		if result != test.expected {
			t.Errorf("Rstrip() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceRstrip(t *testing.T) {
	str := String("hello  ")
	result := str.EnforceRstrip()

	if result != String("hello") {
		t.Errorf("EnforceRstrip() expected 'hello', got '%s'", result)
	}

	if str != String("hello") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Swapcase(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("Hello World"), String("hELLO wORLD")},
		{String("hello"), String("HELLO")},
		{String("WORLD"), String("world")},
		{String(""), String("")},
		{String("123"), String("123")},
	}

	for _, test := range tests {
		result := test.input.Swapcase()
		if result != test.expected {
			t.Errorf("Swapcase() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceSwapcase(t *testing.T) {
	str := String("Hello World")
	result := str.EnforceSwapcase()

	if result != String("hELLO wORLD") {
		t.Errorf("EnforceSwapcase() expected 'hELLO wORLD', got '%s'", result)
	}

	if str != String("hELLO wORLD") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Title(t *testing.T) {
	tests := []struct {
		input    String
		expected String
	}{
		{String("hello world"), String("Hello World")},
		{String("HELLO WORLD"), String("Hello World")},
		{String("hello"), String("Hello")},
		{String(""), String("")},
	}

	for _, test := range tests {
		result := test.input.Title()
		if result != test.expected {
			t.Errorf("Title() for '%s' expected '%s', got '%s'", test.input, test.expected, result)
		}
	}
}

func TestString_EnforceTitle(t *testing.T) {
	str := String("hello world")
	result := str.EnforceTitle()

	if result != String("Hello World") {
		t.Errorf("EnforceTitle() expected 'Hello World', got '%s'", result)
	}

	if str != String("Hello World") {
		t.Errorf("String should be modified in place, got '%s'", str)
	}
}

func TestString_Lines(t *testing.T) {
	tests := []struct {
		input    String
		expected Array[String]
	}{
		{String("hello\nworld"), Array[String]{"hello", "world"}},
		{String("hello"), Array[String]{"hello"}},
		{String(""), Array[String]{""}},
		{String("hello\n\nworld"), Array[String]{"hello", "", "world"}},
	}

	for _, test := range tests {
		result := test.input.Lines()
		if len(result) != len(test.expected) {
			t.Errorf("Lines() for '%s' expected length %d, got %d", test.input, len(test.expected), len(result))
		}
		for i, line := range result {
			if line != test.expected[i] {
				t.Errorf("Lines() for '%s' at index %d expected '%s', got '%s'", test.input, i, test.expected[i], line)
			}
		}
	}
}

func TestString_Words(t *testing.T) {
	tests := []struct {
		input    String
		expected Array[String]
	}{
		{String("hello world"), Array[String]{"hello", "world"}},
		{String("hello"), Array[String]{"hello"}},
		{String(""), Array[String]{}},
		{String("  hello  world  "), Array[String]{"hello", "world"}},
	}

	for _, test := range tests {
		result := test.input.Words()
		if len(result) != len(test.expected) {
			t.Errorf("Words() for '%s' expected length %d, got %d", test.input, len(test.expected), len(result))
		}
		for i, word := range result {
			if word != test.expected[i] {
				t.Errorf("Words() for '%s' at index %d expected '%s', got '%s'", test.input, i, test.expected[i], word)
			}
		}
	}
}

func TestString_ToI(t *testing.T) {
	tests := []struct {
		input    String
		expected Integer
	}{
		{String("123"), Integer(123)},
		{String("0"), Integer(0)},
		{String("-456"), Integer(-456)},
		{String("abc"), Integer(0)}, // Should return 0 for invalid input
		{String(""), Integer(0)},
	}

	for _, test := range tests {
		result := test.input.ToI()
		if result != test.expected {
			t.Errorf("ToI() for '%s' expected %d, got %d", test.input, test.expected, result)
		}
	}
}

func TestString_ToF(t *testing.T) {
	tests := []struct {
		input    String
		expected Float
	}{
		{String("123.45"), Float(123.45)},
		{String("0.0"), Float(0.0)},
		{String("-456.78"), Float(-456.78)},
		{String("abc"), Float(0.0)}, // Should return 0.0 for invalid input
		{String(""), Float(0.0)},
	}

	for _, test := range tests {
		result := test.input.ToF()
		if result != test.expected {
			t.Errorf("ToF() for '%s' expected %f, got %f", test.input, test.expected, result)
		}
	}
}
