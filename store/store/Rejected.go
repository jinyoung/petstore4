package store

import (
	"time"
)

type Rejected struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	
}

func NewRejected() *Rejected{
	event := &Rejected{EventType:"Rejected", TimeStamp:time.Now().String()}

	return event
}
