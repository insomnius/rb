package ruby

type String string

func (s String) Chars() Array[string] {
	chars := make([]string, len(s))
	for _, c := range s {
		chars = append(chars, string(c))
	}
	return chars
}

func (s String) Length() Int {
	return Int(len(s))
}
