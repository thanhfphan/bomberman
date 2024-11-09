package game

import (
	"thanhfphan.com/bomberman/src/engine"
	"thanhfphan.com/bomberman/src/engine/dt"
)

type EntityManager struct {
	entites *dt.ArrayList[engine.Entity]
}

func NewEntityManager() *EntityManager {
	cap := 10
	return &EntityManager{
		entites: dt.NewArrayList[engine.Entity](cap),
	}
}

func (em *EntityManager) Create(entity engine.Entity) int {
	return em.entites.Append(entity)
}

func (em *EntityManager) GetEntity(id int) (engine.Entity, error) {
	entity, err := em.entites.Get(id)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (em *EntityManager) Size() int {
	return em.entites.Size()
}

func (em *EntityManager) Remove(e engine.Entity) error {
	if err := em.entites.Remove(e.GetID()); err != nil {
		return err
	}
	e.Destroy()
	return nil
}
