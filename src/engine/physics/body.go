package physics

import (
	"thanhfphan.com/bomberman/src/engine/math"
)

type Body struct {
	ID           int
	AABB         AABB
	Velocity     math.Vec2
	Acceleration math.Vec2

	IsActive bool
}
