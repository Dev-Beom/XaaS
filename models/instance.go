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

func (instance *Instance) setStateCreating() {
	instance.State = Creating
}

func (instance *Instance) setStateRunning() {
	instance.State = Running
}

func (instance *Instance) setStateStop() {
	instance.State = Stop
}

func (instance *Instance) setStateDeleted() {
	instance.State = Deleted
}

func (instance *Instance) setStatePause() {
	instance.State = Pause
}
