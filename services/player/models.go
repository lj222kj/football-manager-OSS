package player

import positions "github.com/lj222kj/xpkg/positions"

type Attributes struct {
	Speed   float64
	Passing float64
	Stamina float64
}

type Player struct {
	ID         string
	Name       string
	Age        int
	Position   positions.Position
	Attributes Attributes
}
