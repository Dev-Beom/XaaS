package instance

import (
	"github.com/dev-beom/faas/exception"
	"github.com/dev-beom/faas/models"
)

type Repository interface {
	Find(id string) (models.Instance, error)
	FindAll() map[string]models.Instance
	Create(instance models.Instance) error
	Delete(id string) error
	Update(id string, instance models.Instance) (models.Instance, error)
}

type repository struct {
	store map[string]models.Instance
}

func NewRepository() Repository {
	return &repository{}
}

func NewMockRepository(mockDB map[string]models.Instance) Repository {
	return &repository{store: mockDB}
}

func (r *repository) Find(id string) (models.Instance, error) {
	instance, ok := r.store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	return instance, nil
}

func (r *repository) FindAll() map[string]models.Instance {
	return r.store
}

func (r *repository) Create(instance models.Instance) error {
	_, ok := r.store[instance.Id]
	if ok {
		return exception.ErrAlreadyExist
	}
	r.store[instance.Id] = instance
	return nil
}

func (r *repository) Delete(id string) error {
	_, ok := r.store[id]
	if !ok {
		return exception.ErrNotFoundData
	}
	delete(r.store, id)
	return nil
}

func (r *repository) Update(id string, instance models.Instance) (models.Instance, error) {
	_, ok := r.store[id]
	if !ok {
		return models.Instance{}, exception.ErrNotFoundData
	}
	r.store[id] = instance
	return r.store[id], nil
}
