package node

/**
todo 비즈니스 로직 구현
인스턴스 생성, 삭제, 설명 변경, 상태 변경
*/

import (
	"github.com/dev-beom/xaas/apiserver/exception"
	"github.com/dev-beom/xaas/apiserver/models"
	"time"
)

type Service interface {
	Get(id string) (models.Node, error)
	GetAll() map[string]models.Node
	Create(node models.Node) error
	Delete(id string) error
	UpdateDescription(id string, description string) (models.Node, error)
	UpdateState(id string, state string) (models.Node, error)
}

type service struct {
	nodeRepository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Get(id string) (models.Node, error) {
	find, err := s.nodeRepository.Find(id)
	return find, err
}

func (s *service) GetAll() map[string]models.Node {
	return s.nodeRepository.FindAll()
}

func (s *service) Create(node models.Node) error {
	err := s.nodeRepository.Create(node)
	if err != nil {
		return exception.ErrAlreadyExist
	}
	return nil
}
func (s *service) Delete(id string) error {
	err := s.nodeRepository.Delete(id)
	if err != nil {
		return exception.ErrNotFoundData
	}
	return nil
}
func (s *service) UpdateDescription(id string, description string) (models.Node, error) {
	foundNode, err := s.nodeRepository.Find(id)
	if err != nil {
		return models.Node{}, err
	}
	foundNode = models.Node{
		Id:          foundNode.Id,
		Description: description,
		CreateAt:    foundNode.CreateAt,
		UpdateAt:    time.Now(),
		State:       foundNode.State,
	}
	updatedNode, err := s.nodeRepository.Update(id, foundNode)
	if err != nil {
		return models.Node{}, err
	}
	return updatedNode, nil
}

func (s *service) UpdateState(id string, state string) (models.Node, error) {
	foundNode, err := s.nodeRepository.Find(id)
	if err != nil {
		return models.Node{}, nil
	}
	foundNode = models.Node{
		Id:          foundNode.Id,
		Description: foundNode.Description,
		CreateAt:    foundNode.CreateAt,
		UpdateAt:    time.Now(),
		State:       state,
	}
	updatedNode, err := s.nodeRepository.Update(id, foundNode)
	if err != nil {
		return models.Node{}, nil
	}
	return updatedNode, nil
}
