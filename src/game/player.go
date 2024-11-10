package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/math"
)

type Player struct {
	ID          int
	Speed       float64
	Velocity    math.Vec2
	Position    math.Vec2
	AnimationID int

	Deleted bool
}

func (p *Player) GetID() int {
	return p.ID
}

func (p *Player) Update(deltaTime float64) {
	if !p.IsActive() {
		return
	}

	p.Velocity.X = 0
	p.Velocity.Y = 0
	if global.input.Left == engine.KeyStatePressed || global.input.Left == engine.KeyStateHeld {
		p.Velocity.X -= p.Speed
	}
	if global.input.Right == engine.KeyStatePressed || global.input.Right == engine.KeyStateHeld {
		p.Velocity.X += p.Speed
	}
	if global.input.Up == engine.KeyStatePressed || global.input.Up == engine.KeyStateHeld {
		p.Velocity.Y -= p.Speed
	}
	if global.input.Down == engine.KeyStatePressed || global.input.Down == engine.KeyStateHeld {
		p.Velocity.Y += p.Speed
	}

	if p.Velocity.Y < 0 {
		p.AnimationID = animationWalkBackID
	} else if p.Velocity.Y > 0 {
		p.AnimationID = animationWalkFrontID
	} else if p.Velocity.X != 0 {
		p.AnimationID = animationWalkRightID // Walk left will be flipped
	} else {
		p.AnimationID = animationIdleID
	}

	p.Position.X += p.Velocity.X * deltaTime
	p.Position.Y += p.Velocity.Y * deltaTime
}

func (p *Player) IsActive() bool {
	return p != nil && !p.Deleted
}

func (p *Player) Render(screen *ebiten.Image) {
	if !p.IsActive() {
		return
	}
	if p.AnimationID < 0 {
		return
	}
	animation := global.animation.GetAnimation(p.AnimationID)
	if animation == nil {
		return
	}
	if p.Velocity.Y == 0 {
		if p.Velocity.X < 0 {
			animation.IsFlipped = true
		} else if p.Velocity.X > 0 {
			animation.IsFlipped = false
		}
	}
	aframe := animation.Definition.Frames[animation.CurrentFrameIndex]
	animation.Definition.SpriteSheet.DrawFrame(screen, float64(aframe.Row), float64(aframe.Column), p.Position, animation.IsFlipped)

}

func (p *Player) Destroy() {
	p.Deleted = true
}
