# API 명세
### `GET` `/api/node/:id` 개별 노드의 정보 얻기
### `GET` `/api/nodes` 모든 노드의 정보 얻기
### `POST` `/api/node` 노드 생성하기 
`Content-Type: application/json`  
`Body: json`
```json
{
  "id": "string",
  "description": "string"
}
```

### `DELETE` `/api/node/:id` 노드 삭제하기
### `PATCH` `/api/node/:id` 노드 설명 변경하기
`Content-Type: application/json`  
`Body: json`
```json
{
  "description": "string"
}
```