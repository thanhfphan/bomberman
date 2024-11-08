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
	IsActive          bool
	IsFlipped         bool
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

func (m *Manager) CreateDefinition(spriteSheet *spritesheet.SpriteSheet, duration float32, row uint8, column []uint8, frameCount uint8) int {
	if frameCount > MaxFrame {
		panic(fmt.Errorf("frame count exceeds maximum frame count"))
	}

	def := &Definition{
		SpriteSheet: spriteSheet,
		FrameCount:  frameCount,
	}
	for i := uint8(0); i < frameCount; i++ {
		def.Frames[i] = Frame{
			Duration: duration,
			Row:      row,
			Column:   column[i],
		}
	}

	return m.definitions.Append(def)
}

func (m *Manager) CreateAnimation(definitionID int, doesLoop bool) (int, error) {
	def, err := m.definitions.Get(definitionID)
	if err != nil {
		return -1, fmt.Errorf("could not find definition with id %d", definitionID)
	}
	animation := &Animation{
		Definition:       def,
		DoesLoop:         doesLoop,
		IsActive:         true,
		CurrentFrameTime: def.Frames[0].Duration,
	}

	id := m.animations.Append(animation)

	return id, nil
}

func (m *Manager) DestroyAnimation(index int) error {
	return m.animations.Remove(index)
}

func (m *Manager) GetAnimation(index int) (*Animation, error) {
	return m.animations.Get(index)
}

func (m *Manager) Update(deltaTime float64) {
	for i := 0; i < m.animations.Size(); i++ {
		animation, err := m.animations.Get(i)
		if err != nil {
			panic(fmt.Errorf("could not get animation: %v", err)) // Should never happen
		}
		if !animation.IsActive {
			continue
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
