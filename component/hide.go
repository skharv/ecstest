package component

type Hide struct {
	Value bool
}

func NewHide(value bool) Hide {
	return Hide{value}
}
