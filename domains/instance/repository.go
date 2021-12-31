package instance

import (
	"github.com/dev-beom/faas/exception"
	"github.com/dev-beom/faas/models"
)

var store = make(map[string]models.Instance)

type Repository interface {
	Find(id string) (models.Instance, error)
	Create(instance models.Instance) error
	Delete(id string) error
	Update(id string, instance models.Instance) (models.Instance, error)
}

type repository struct {
}

func (this *repository) Find(id string) (models.Instance, error) {
	instance, ok := store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	return instance, nil
}

func (this *repository) Create(instance models.Instance) error {
	_, ok := store[instance.Id]
	if ok {
		return exception.ErrAlreadyExist
	}
	store[instance.Id] = instance
	return nil
}

func (this *repository) Delete(id string) error {
	_, ok := store[id]
	if !ok {
		return exception.ErrNotFoundData
	}
	delete(store, id)
	return nil
}

func (this *repository) Update(id string, instance models.Instance) (models.Instance, error) {
	_, ok := store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	store[id] = instance
	return store[id], nil
}
