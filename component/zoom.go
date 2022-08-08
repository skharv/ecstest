package component

type Zoom struct {
	Value float64
}

func NewZoomI(value int) Zoom {
	return Zoom{float64(value)}
}

func NewZoomF(value float64) Zoom {
	return Zoom{value}
}
