package animation

import (
	"fmt"

	"thanhfphan.com/bomberman/src/engine/dt"
	"thanhfphan.com/bomberman/src/engine/spritesheet"
)

const (
	MaxFrame = 16
)

type Frame struct {
	Duration float32
	Row      uint8
	Column   uint8
}

type Definition struct {
	SpriteSheet *spritesheet.SpriteSheet
	Frames      [MaxFrame]Frame
	FrameCount  uint8
}

type Animation struct {
	Definition        *Definition
	CurrentFrameTime  float32
	CurrentFrameIndex uint8
	DoesLoop          bool
	IsFlipped         bool
}

func (a *Animation) Reset() {
	a.CurrentFrameTime = a.Definition.Frames[0].Duration
	a.CurrentFrameIndex = 0
}

type Manager struct {
	definitions *dt.ArrayList[*Definition]
	animations  *dt.ArrayList[*Animation]
}

func NewManager() *Manager {
	return &Manager{
		definitions: dt.NewArrayList[*Definition](0),
		animations:  dt.NewArrayList[*Animation](0),
	}
}

func (m *Manager) CreateDefinition(spriteSheet *spritesheet.SpriteSheet, duration float32, rows []uint8, column []uint8, frameCount uint8) int {
	if frameCount <= 0 {
		panic(fmt.Errorf("frame count must be greater than 0"))
	}
	if frameCount > MaxFrame {
		// Create definition only called when initializing the game, so this ok to panic
		panic(fmt.Errorf("frame count exceeds maximum frame count"))
	}

	def := &Definition{
		SpriteSheet: spriteSheet,
		FrameCount:  frameCount,
	}
	for i := uint8(0); i < frameCount; i++ {
		def.Frames[i] = Frame{
			Duration: duration,
			Row:      rows[i],
			Column:   column[i],
		}
	}

	return m.definitions.Append(def)
}

func (m *Manager) CreateAnimation(definitionID int, doesLoop bool) int {
	def, err := m.definitions.Get(definitionID)
	if err != nil {
		// Create animation only called when initializing the game, so this ok to panic
		panic(fmt.Errorf("could not find definition with id %d", definitionID))
	}
	animation := &Animation{
		Definition:       def,
		DoesLoop:         doesLoop,
		CurrentFrameTime: def.Frames[0].Duration,
	}

	id := m.animations.Append(animation)

	return id
}

func (m *Manager) DestroyAnimation(index int) error {
	return m.animations.Remove(index)
}

func (m *Manager) GetAnimation(index int) *Animation {
	animation, _ := m.animations.Get(index)
	return animation
}

func (m *Manager) Update(deltaTime float64) {
	for i := 0; i < m.animations.Size(); i++ {
		animation, err := m.animations.Get(i)
		if err == dt.ErrIndexOutOfRange {
			continue
		}
		if err != nil {
			panic(fmt.Errorf("could not get animation: %v", err)) // Should never happen
		}

		animation.CurrentFrameTime -= float32(deltaTime)
		if animation.CurrentFrameTime <= 0 {
			animation.CurrentFrameIndex += 1

			// Loop for staying in the last frame
			if animation.CurrentFrameIndex >= animation.Definition.FrameCount {
				if animation.DoesLoop {
					animation.CurrentFrameIndex = 0
				} else {
					animation.CurrentFrameIndex = animation.Definition.FrameCount - 1
				}
			}

			animation.CurrentFrameTime = animation.Definition.Frames[animation.CurrentFrameIndex].Duration
		}
	}
}
