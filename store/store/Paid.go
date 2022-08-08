package store

import (
	"time"
)

type Paid struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	OrderId int `json:"orderId" type:"int"` 
	
}

func NewPaid() *Paid{
	event := &Paid{EventType:"Paid", TimeStamp:time.Now().String()}

	return event
}
