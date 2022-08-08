package component

type Position struct {
	X, Y float64
}

func NewPositionI(x, y int) Position {
	return Position{float64(x), float64(y)}
}
