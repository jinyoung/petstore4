package store

import (
	"time"
)

type Cooked struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	FoodId string `json:"foodId" type:"string"` 
	Preference string `json:"preference" type:"string"` 
	OrderId int `json:"orderId" type:"int"` 
	Status string `json:"status" type:"string"` 
	
}

func NewCooked() *Cooked{
	event := &Cooked{EventType:"Cooked", TimeStamp:time.Now().String()}

	return event
}
