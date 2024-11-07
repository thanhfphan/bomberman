package physics

import (
	"thanhfphan.com/bomberman/src/engine/dt"
)

type BodyManager struct {
	bodies *dt.ArrayList[*Body]
}

func NewBodyManager(capacity int) *BodyManager {
	if capacity <= 0 {
		capacity = 1
	}
	return &BodyManager{
		bodies: dt.NewArrayList[*Body](capacity),
	}
}

func (b *BodyManager) Create() (*Body, error) {
	body := &Body{
		IsActive: true,
	}
	body.ID = b.bodies.Append(body)

	return body, nil
}

func (b *BodyManager) Get(id int) (*Body, error) {
	body, err := b.bodies.Get(id)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (b *BodyManager) Remove(e *Body) error {
	return b.bodies.Remove(e.ID)
}

func (b *BodyManager) RemoveID(id int) error {
	return b.bodies.Remove(id)
}

func (b *BodyManager) Size() int {
	return b.bodies.Size()
}

func (b *BodyManager) Update(deltaTime float64) {
	for i := 0; i < b.bodies.Size(); i++ {
		body, err := b.bodies.Get(i)
		if err != nil {
			continue
		}
		if !body.IsActive {
			continue
		}

		body.AABB.Position.X += body.Velocity.X * deltaTime
		body.AABB.Position.Y += body.Velocity.Y * deltaTime
	}
}
