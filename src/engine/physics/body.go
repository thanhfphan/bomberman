package physics

import (
	"thanhfphan.com/bomberman/src/engine/math"
)

type Body struct {
	ID       int
	Position math.Vec2
	Velocity math.Vec2

	IsActive bool
}
