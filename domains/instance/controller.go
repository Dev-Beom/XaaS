package instance

/**
todo 3 layer 구조 완성
instance 관련 end point 정리
validation 코드 추가
*/

type controller struct {
	instanceService Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

