package component

type Speed struct {
	Value float64
}

func NewSpeedI(value int) Speed {
	return Speed{float64(value)}
}

func NewSpeedF(value float64) Speed {
	return Speed{value}
}
