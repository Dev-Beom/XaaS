package node

import (
	"encoding/json"
	"fmt"
	"github.com/dev-beom/xaas/apiserver/constants/IPCMessage"
	"github.com/dev-beom/xaas/apiserver/exception"
	"github.com/dev-beom/xaas/apiserver/models"
	ipc "github.com/james-barrow/golang-ipc"
)

type Repository interface {
	Find(id string) (models.Node, error)
	FindAll() map[string]models.Node
	Create(node models.Node) error
	Delete(id string) error
	Update(id string, node models.Node) (models.Node, error)
}

type repository struct {
	store     map[string]models.Node
	ipcServer *ipc.Server
}

func NewRepository() Repository {
	ipcServer, _ := ipc.StartServer("XaaS", nil)
	return &repository{
		store:     make(map[string]models.Node),
		ipcServer: ipcServer,
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
	data, ok := r.store[node.Id]
	if ok {
		fmt.Println("이미 존재", data)
		return exception.ErrAlreadyExist
	}
	bytes, _ := json.Marshal(node)
	err := r.ipcServer.Write(IPCMessage.CREATE, bytes)
	if err != nil {
		return exception.ErrNodeCreate
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
