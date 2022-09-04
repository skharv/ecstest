package enum

type Group int

const (
	GroupNone = Group(iota)
	GroupUnit
	GroupTile
)
