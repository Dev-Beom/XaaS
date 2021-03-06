package node

import (
	"encoding/json"
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

func NewRepository(ipcServer *ipc.Server) Repository {
	return &repository{
		store:     make(map[string]models.Node),
		ipcServer: ipcServer,
	}
}

func NewMockRepository(mockDB map[string]models.Node) Repository {
	return &repository{store: mockDB, ipcServer: nil}
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
	if _, ok := r.store[node.Id]; ok {
		return exception.ErrAlreadyExist
	}
	err := r.sendNodeByIPC(node, IPCMessage.CREATE)
	if err != nil {
		return exception.ErrNodeCreate
	}
	r.store[node.Id] = node
	return nil
}

func (r *repository) Delete(id string) error {
	if _, ok := r.store[id]; !ok {
		return exception.ErrNotFoundData
	}
	err := r.sendNodeByIPC(r.store[id], IPCMessage.DELETE)
	if err != nil {
		return exception.ErrNodeDelete
	}
	delete(r.store, id)
	return nil
}

func (r *repository) Update(id string, node models.Node) (models.Node, error) {
	if _, ok := r.store[id]; !ok {
		return models.Node{}, exception.ErrNotFoundData
	}
	err := r.sendNodeByIPC(node, IPCMessage.UPDATE)
	if err != nil {
		return models.Node{}, exception.ErrNodeUpdate
	}
	r.store[id] = node
	return r.store[id], nil
}

func (r *repository) sendNodeByIPC(node models.Node, msgType int) error {
	if r.ipcServer == nil {
		return nil
	}
	bytes, err := json.Marshal(node)
	if err != nil {
		return err
	}
	err = r.ipcServer.Write(msgType, bytes)
	if err != nil {
		return err
	}
	return nil
}
