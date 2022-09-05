package component

type Render struct {
	Value int
}

func NewRender(value int) Render {
	return Render{Value: value}
}
