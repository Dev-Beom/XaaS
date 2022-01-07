package exception

import "errors"

// Common
var (
	ErrNotFoundData = errors.New("요청하신 데이터를 찾을 수 없습니다")
	ErrAlreadyExist = errors.New("요청하신 데이터가 이미 존재합니다")
)

// Node
var (
	ErrNodeNotAllowAccess = errors.New("노드에 대한 접근권한이 없습니다")
	ErrNodeStateChange    = errors.New("노드의 상태를 해당 상태로 변경할 수 없습니다")
)

// Model
var (
	ErrModelUpload = errors.New("모델을 업로드할 수 없습니다")
	ErrModelDelete = errors.New("모델을 제거할 수 없습니다")
)
