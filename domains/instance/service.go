package instance

/**
todo 비즈니스 로직 구현
인스턴스 생성, 삭제, 설명 변경, 상태 변경
*/

import (
	"github.com/dev-beom/xaas/exception"
	"github.com/dev-beom/xaas/models"
	"time"
)

type Service interface {
	Get(id string) (models.Instance, error)
	GetAll() map[string]models.Instance
	Create(instance models.Instance) error
	Delete(id string) error
	UpdateDescription(id string, description string) (models.Instance, error)
	UpdateState(id string, state string) (models.Instance, error)
}

type service struct {
	instanceRepository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Get(id string) (models.Instance, error) {
	find, err := s.instanceRepository.Find(id)
	return find, err
}

func (s *service) GetAll() map[string]models.Instance {
	return s.instanceRepository.FindAll()
}

func (s *service) Create(instance models.Instance) error {
	err := s.instanceRepository.Create(instance)
	if err != nil {
		return exception.ErrAlreadyExist
	}
	return nil
}
func (s *service) Delete(id string) error {
	err := s.instanceRepository.Delete(id)
	if err != nil {
		return exception.ErrNotFoundData
	}
	return nil
}
func (s *service) UpdateDescription(id string, description string) (models.Instance, error) {
	foundInstance, err := s.instanceRepository.Find(id)
	if err != nil {
		return models.Instance{}, err
	}
	foundInstance = models.Instance{
		Id:          foundInstance.Id,
		Description: description,
		CreateAt:    foundInstance.CreateAt,
		UpdateAt:    time.Now(),
		State:       foundInstance.State,
	}
	updatedInstance, err := s.instanceRepository.Update(id, foundInstance)
	if err != nil {
		return models.Instance{}, err
	}
	return updatedInstance, nil
}

func (s *service) UpdateState(id string, state string) (models.Instance, error) {
	foundInstance, err := s.instanceRepository.Find(id)
	if err != nil {
		return models.Instance{}, nil
	}
	foundInstance = models.Instance{
		Id:          foundInstance.Id,
		Description: foundInstance.Description,
		CreateAt:    foundInstance.CreateAt,
		UpdateAt:    time.Now(),
		State:       state,
	}
	updatedInstance, err := s.instanceRepository.Update(id, foundInstance)
	if err != nil {
		return models.Instance{}, nil
	}
	return updatedInstance, nil
}
