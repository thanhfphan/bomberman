package engine

import (
	"fmt"
	"log/slog"
)

type EntityManager struct {
	entites    *ArrayList[*Entity]
	bodyManger *BodyManager
}

func NewEntityManager() *EntityManager {
	cap := 10
	bm := NewBodyManager(cap)

	return &EntityManager{
		entites:    NewArrayList[*Entity](cap),
		bodyManger: bm,
	}
}

func (em *EntityManager) CreateEntity() (*Entity, error) {
	body, err := em.bodyManger.Create()
	if err != nil {
		return nil, fmt.Errorf("could not create body: %v", err)
	}

	entity := &Entity{
		Body: body,
	}
	id, err := em.entites.Append(entity)
	if err != nil {
		return nil, err
	}
	entity.ID = id

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
