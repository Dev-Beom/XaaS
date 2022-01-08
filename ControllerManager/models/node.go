package models

import "time"

type Node struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	State       string    `json:"state"`
}
