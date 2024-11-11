package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/math"
)

type PlayerState int

const (
	PlayerStateIdle PlayerState = iota
	PlayerStateLongIdle
	PlayerStateWalking
)

type Player struct {
	ID               int
	Speed            float64
	Velocity         math.Vec2
	Position         math.Vec2
	AnimationID      int
	AnimationFlipped bool
	State            PlayerState
	LastKeypressed   engine.InputKey

	Deleted bool
}

func NewPlayer(speed float64, position math.Vec2) *Player {
	player := &Player{
		Speed:       speed,
		Position:    position,
		State:       PlayerStateIdle,
		AnimationID: animationIdleFrontID,
	}
	player.ID = global.entity.Create(player)
	return player
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
	if p.isKeyPressed(global.input.Left) {
		p.Velocity.X -= p.Speed
		p.LastKeypressed = engine.InputKeyLeft
	}
	if p.isKeyPressed(global.input.Right) {
		p.Velocity.X += p.Speed
		p.LastKeypressed = engine.InputKeyRight
	}
	if p.isKeyPressed(global.input.Up) {
		p.Velocity.Y -= p.Speed
		p.LastKeypressed = engine.InputKeyUp
	}
	if p.isKeyPressed(global.input.Down) {
		p.Velocity.Y += p.Speed
		p.LastKeypressed = engine.InputKeyDown
	}
	if p.Velocity.X != 0 || p.Velocity.Y != 0 {
		p.State = PlayerStateWalking
	} else {
		p.State = PlayerStateIdle
	}

	p.Position.X += p.Velocity.X * deltaTime
	p.Position.Y += p.Velocity.Y * deltaTime

	snapThreshold := BaseSnapThreshold
	if p.isMovingFast() {
		snapThreshold = HighSpeedSnapThreshold
	}

	if p.isNearTileCenter(p.Position, snapThreshold) && p.State == PlayerStateIdle {
		p.snapToTileCenter()
	}
}

func (p *Player) IsActive() bool {
	return p != nil && !p.Deleted
}

func (p *Player) Render(screen *ebiten.Image) {
	if !p.IsActive() {
		return
	}
	p.AnimationFlipped = false
	if p.State == PlayerStateWalking {
		if p.isKeyPressed(global.input.Left) {
			p.AnimationID = animationWalkRightID
			p.AnimationFlipped = true
		} else if p.isKeyPressed(global.input.Right) {
			p.AnimationID = animationWalkRightID
		} else if p.isKeyPressed(global.input.Up) {
			p.AnimationID = animationWalkBackID
		} else if p.isKeyPressed(global.input.Down) {
			p.AnimationID = animationWalkFrontID
		}
	} else {
		if p.LastKeypressed == engine.InputKeyLeft {
			p.AnimationID = animationIdleRightID
			p.AnimationFlipped = true
		} else if p.LastKeypressed == engine.InputKeyRight {
			p.AnimationID = animationIdleRightID
		} else if p.LastKeypressed == engine.InputKeyUp {
			p.AnimationID = animationIdleBackID
		} else if p.LastKeypressed == engine.InputKeyDown {
			p.AnimationID = animationIdleFrontID
		}
	}
	animation := global.animation.GetAnimation(p.AnimationID)
	aframe := animation.Definition.Frames[animation.CurrentFrameIndex]
	animation.Definition.SpriteSheet.DrawFrame(
		screen,
		float64(aframe.Row),
		float64(aframe.Column),
		p.Position,
		p.AnimationFlipped,
	)
}

func (p *Player) Destroy() {
	p.Deleted = true
}

func (p *Player) isKeyPressed(key engine.KeyState) bool {
	return key == engine.KeyStatePressed || key == engine.KeyStateHeld
}

func (p *Player) isNearTileCenter(pos math.Vec2, threshold float64) bool {
	centerX := float64(int(pos.X/TileSize)*TileSize + TileSize/2)
	centerY := float64(int(pos.Y/TileSize)*TileSize + TileSize/2)

	dx := math.Abs(pos.X - centerX)
	dy := math.Abs(pos.Y - centerY)

	return dx <= threshold && dy <= threshold
}

func (p *Player) snapToTileCenter() {
	p.Position.X = float64(int(p.Position.X/TileSize)*TileSize + TileSize/2)
	p.Position.Y = float64(int(p.Position.Y/TileSize)*TileSize + TileSize/2)
}

func (p *Player) isMovingFast() bool {
	return math.Abs(p.Velocity.X) > HighSpeedThreshold || math.Abs(p.Velocity.Y) > HighSpeedThreshold
}
