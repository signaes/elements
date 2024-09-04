package elements

type text struct {
	value string
}

func Text(value string) text {
	return text{value: value}
}

func (t text) String() string {
	return t.value
}
