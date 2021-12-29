package instance

import (
	"github.com/dev-beom/faas/exception"
	"github.com/dev-beom/faas/models"
)

var store = make(map[string]models.Instance)

type Repository interface {
	Create(instance models.Instance) error
	Delete(id string) error
	UpdateDescription(instance models.Instance) (models.Instance, error)
	UpdateState(id string, state string) (models.Instance, error)
}

type repository struct {
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

func (this *repository) UpdateDescription(instance models.Instance) (models.Instance, error) {
	_, ok := store[instance.Id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	store[instance.Id] = instance
	return store[instance.Id], nil
}

func (this *repository) UpdateState(id string, state string) (models.Instance, error) {
	findInstance, ok := store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	findInstance.State = state
	return findInstance, nil
}
