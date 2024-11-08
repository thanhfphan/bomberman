package engine

import (
	"fmt"
	"log/slog"

	"thanhfphan.com/bomberman/src/engine/dt"
	"thanhfphan.com/bomberman/src/engine/physics"
)

type EntityManager struct {
	entites    *dt.ArrayList[*Entity]
	bodyManger *physics.BodyManager
}

func NewEntityManager() *EntityManager {
	cap := 10
	bm := physics.NewBodyManager(cap)

	return &EntityManager{
		entites:    dt.NewArrayList[*Entity](cap),
		bodyManger: bm,
	}
}

func (em *EntityManager) CreateEntity() (*Entity, error) {
	body, err := em.bodyManger.Create()
	if err != nil {
		return nil, fmt.Errorf("could not create body: %v", err)
	}

	entity := &Entity{
		Body:        body,
		AnimationID: -1,
	}
	entity.ID = em.entites.Append(entity)

	return entity, nil
}

func (em *EntityManager) GetEntity(id int) (*Entity, error) {
	entity, err := em.entites.Get(id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (em *EntityManager) RemoveEntity(e *Entity) error {
	if err := em.bodyManger.Remove(e.Body); err != nil {
		slog.Error("remove body entity error:" + err.Error())
		return err
	}
	return em.entites.Remove(e.ID)
}

func (em *EntityManager) Size() int {
	return em.entites.Size()
}

func (em *EntityManager) Update(deltaTime float64) {
	em.bodyManger.Update(deltaTime)
}
