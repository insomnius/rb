package rb

type Boolean bool

func (b Boolean) ToS() String {
	if b {
		return "true"
	}

	return "false"
}
