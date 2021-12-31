package instance

/**
todo 비즈니스 로직 구현
인스턴스 생성, 삭제, 설명 변경, 상태 변경
*/

import (
	"github.com/dev-beom/faas/exception"
	"github.com/dev-beom/faas/models"
	"time"
)

type Service interface {
	Create(instance models.Instance) error
	Delete(id string) error
	UpdateDescription(id string, description string) (models.Instance, error)
	UpdateState(id string, state string) (models.Instance, error)
}

type service struct {
	instanceRepository Repository
}

func (this *service) Create(instance models.Instance) error {
	err := this.instanceRepository.Create(instance)
	if err != nil {
		return exception.ErrAlreadyExist
	}
	return nil
}
func (this *service) Delete(id string) error {
	err := this.instanceRepository.Delete(id)
	if err != nil {
		return exception.ErrNotFoundData
	}
	return nil
}
func (this *service) UpdateDescription(id string, description string) (models.Instance, error) {
	foundInstance, _ := this.instanceRepository.Find(id)
	foundInstance = models.Instance{
		Id:          foundInstance.Id,
		Description: description,
		CreateAt:    foundInstance.CreateAt,
		UpdateAt:    time.Now(),
		State:       foundInstance.State,
	}
	updatedInstance, _ := this.instanceRepository.Update(id, foundInstance)
	return updatedInstance, nil
}

func (this *service) UpdateState(id string, state string) (models.Instance, error) {
	foundInstance, _ := this.instanceRepository.Find(id)
	foundInstance = models.Instance{
		Id:          foundInstance.Id,
		Description: foundInstance.Description,
		CreateAt:    foundInstance.CreateAt,
		UpdateAt:    time.Now(),
		State:       state,
	}
	updatedInstance, _ := this.instanceRepository.Update(id, foundInstance)
	return updatedInstance, nil
}
