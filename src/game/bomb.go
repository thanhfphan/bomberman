package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine/audio"
	"thanhfphan.com/bomberman/src/engine/math"
)

type Bomb struct {
	ID          int
	Countdown   time.Duration
	PlacedAt    time.Time
	Position    math.Vec2
	AnimationID int

	Deleted bool
}

func (b *Bomb) GetID() int {
	return b.ID
}

func (b *Bomb) Update(deltaTime float64) {
	if !b.IsActive() {
		return
	}

	if time.Since(b.PlacedAt) >= b.Countdown {
		audio.Play(soundBombExplode)
		gs.entityManager.Remove(b)
	}
}

func (b *Bomb) IsActive() bool {
	return b != nil && !b.Deleted
}

func (b *Bomb) Render(screen *ebiten.Image) {
	if !b.IsActive() {
		return
	}
	if b.AnimationID < 0 {
		return
	}
	animation := gs.animationManager.GetAnimation(b.AnimationID)
	if animation == nil {
		return
	}
	aframe := animation.Definition.Frames[animation.CurrentFrameIndex]
	animation.Definition.SpriteSheet.DrawFrame(screen, float64(aframe.Row), float64(aframe.Column), b.Position, animation.IsFlipped)
}

func (b *Bomb) Destroy() {
	b.Deleted = true
}