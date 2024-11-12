package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine/math"
)

type Bat struct {
	EntityID    int
	AnimationID int
	Position    math.Vec2

	Deleted bool
}

func NewBat(position math.Vec2) *Bat {
	bat := &Bat{
		Position: math.Vec2{
			X: float64(int(position.X/TileSize)*TileSize + TileSize/2),
			Y: float64(int(position.Y/TileSize)*TileSize + TileSize/2),
		},
		AnimationID: global.animation.CreateAnimation(batDefWalkRightID, true),
	}
	bat.EntityID = global.entity.Create(bat)

	return bat
}

func (b *Bat) GetID() int {
	return b.EntityID
}

func (b *Bat) Update(deltaTime float64) {

}

func (b *Bat) IsActive() bool {
	return b != nil && !b.Deleted
}

func (b *Bat) Render(screen *ebiten.Image) {
	if !b.IsActive() {
		return
	}

	animation := global.animation.GetAnimation(b.AnimationID)
	aframe := animation.Definition.Frames[animation.CurrentFrameIndex]
	animation.Definition.SpriteSheet.DrawFrame(screen, float64(aframe.Row), float64(aframe.Column), b.Position, animation.IsFlipped)
}

func (b *Bat) Destroy() {
	b.Deleted = true
}
