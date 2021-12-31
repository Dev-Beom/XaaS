package instance

import (
	"github.com/dev-beom/faas/exception"
	"github.com/dev-beom/faas/models"
)

var store = make(map[string]models.Instance)

type Repository interface {
	Find(id string) (models.Instance, error)
	FindAll() map[string]models.Instance
	Create(instance models.Instance) error
	Delete(id string) error
	Update(id string, instance models.Instance) (models.Instance, error)
}

type repository struct {
}

func (r *repository) Find(id string) (models.Instance, error) {
	instance, ok := store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	return instance, nil
}

func (r *repository) FindAll() map[string]models.Instance {
	return store
}

func (r *repository) Create(instance models.Instance) error {
	_, ok := store[instance.Id]
	if ok {
		return exception.ErrAlreadyExist
	}
	store[instance.Id] = instance
	return nil
}

func (r *repository) Delete(id string) error {
	_, ok := store[id]
	if !ok {
		return exception.ErrNotFoundData
	}
	delete(store, id)
	return nil
}

func (r *repository) Update(id string, instance models.Instance) (models.Instance, error) {
	_, ok := store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	store[id] = instance
	return store[id], nil
}
