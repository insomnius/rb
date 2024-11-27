package ruby

type String string

func (s String) Chars() Array[String] {
	chars := make([]String, len(s))
	for _, c := range s {
		chars = append(chars, String(c))
	}
	return chars
}

func (s String) Length() Int {
	return Int(len(s))
}
