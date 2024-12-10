package rb

import (
	"strings"
)

type String string

func (s String) Chars() Array[String] {
	chars := make([]String, len(s))
	for k, c := range s {
		chars[k] = String(c)
	}
	return chars
}

func (s String) Length() Integer {
	return Integer(len(s))
}

func (s String) ToS() String {
	return s
}

func (s String) ToStr() String {
	return s.ToS()
}

func (s String) Downcase() String {
	return String(strings.ToLower(string(s)))
}

func (s *String) DowncaseBang() {
	*s = s.Downcase()
}

func (s String) Upcase() String {
	return String(strings.ToUpper(string(s)))
}

func (s *String) UpcaseBang() {
	*s = s.Upcase()
}
