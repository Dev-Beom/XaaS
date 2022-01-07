package models

import "time"

const (
	Creating = "Creating"
	Running  = "Running"
	Stop     = "Stop"
	Deleted  = "Deleted"
	Pause    = "Pause"
)

type Node struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	State       string    `json:"state"`
}

type NodeCreateRequestDto struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

func (node *Node) SetStateCreating() {
	node.State = Creating
}

func (node *Node) SetStateRunning() {
	node.State = Running
}

func (node *Node) SetStateStop() {
	node.State = Stop
}

func (node *Node) SetStateDeleted() {
	node.State = Deleted
}

func (node *Node) SetStatePause() {
	node.State = Pause
}
