package engine

import "thanhfphan.com/bomberman/src/engine/physics"

type Entity struct {
	ID          int
	Body        *physics.Body
	AnimationID int
}
