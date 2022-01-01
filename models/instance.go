package models

import "time"

const (
	Creating = "Creating"
	Running  = "Running"
	Stop     = "Stop"
	Deleted  = "Deleted"
	Pause    = "Pause"
)

type Instance struct {
	Id          string    `bson:"id"`
	Description string    `bson:"description"`
	CreateAt    time.Time `bson:"create_at"`
	UpdateAt    time.Time `bson:"update_at"`
	State       string    `bson:"state"`
}

type InstanceCreateRequestDto struct {
	Id          string `bson:"id"`
	Description string `bson:"description"`
}

func (instance *Instance) SetStateCreating() {
	instance.State = Creating
}

func (instance *Instance) SetStateRunning() {
	instance.State = Running
}

func (instance *Instance) SetStateStop() {
	instance.State = Stop
}

func (instance *Instance) SetStateDeleted() {
	instance.State = Deleted
}

func (instance *Instance) SetStatePause() {
	instance.State = Pause
}
