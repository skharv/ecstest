package component

type Control struct {
	MoveSpeed float64
}

func NewControl(moveSpeed float64) Control {
	return Control{
		MoveSpeed: moveSpeed,
	}
}
