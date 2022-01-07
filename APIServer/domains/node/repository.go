package node

import (
	"github.com/dev-beom/xaas/apiserver/exception"
	"github.com/dev-beom/xaas/apiserver/models"
)

type Repository interface {
	Find(id string) (models.Node, error)
	FindAll() map[string]models.Node
	Create(node models.Node) error
	Delete(id string) error
	Update(id string, node models.Node) (models.Node, error)
}

type repository struct {
	store map[string]models.Node
}

func NewRepository() Repository {
	return &repository{
		store: make(map[string]models.Node),
	}
}

func NewMockRepository(mockDB map[string]models.Node) Repository {
	return &repository{store: mockDB}
}

func (r *repository) Find(id string) (models.Node, error) {
	node, ok := r.store[id]
	if !ok {
		return models.Node{}, exception.ErrNotFoundData
	}
	return node, nil
}

func (r *repository) FindAll() map[string]models.Node {
	return r.store
}

func (r *repository) Create(node models.Node) error {
	_, ok := r.store[node.Id]
	if ok {
		return exception.ErrAlreadyExist
	}

	r.store[node.Id] = node
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

func (r *repository) Update(id string, node models.Node) (models.Node, error) {
	_, ok := r.store[id]
	if !ok {
		return models.Node{}, exception.ErrNotFoundData
	}
	r.store[id] = node
	return r.store[id], nil
}
