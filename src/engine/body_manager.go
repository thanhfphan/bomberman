package engine

type BodyManager struct {
	bodies *ArrayList[*Body]
}

func NewBodyManager(capacity int) *BodyManager {
	if capacity <= 0 {
		capacity = 1
	}
	return &BodyManager{
		bodies: NewArrayList[*Body](capacity),
	}
}

func (b *BodyManager) Create() (*Body, error) {
	body := &Body{}
	id, err := b.bodies.Append(body)
	if err != nil {
		return nil, err
	}
	body.ID = id

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
