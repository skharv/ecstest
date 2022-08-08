package component

type Velocity struct {
	X, Y float64
}

func NewVelocityI(x, y int) Velocity {
	return Velocity{float64(x), float64(y)}
}

func NewVelocityF(x, y float64) Velocity {
	return Velocity{x, y}
}
