package store

import (
	"time"
)

type OrderCanceled struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	FoodId string `json:"foodId" type:"string"` 
	CustomerId string `json:"customerId" type:"string"` 
	Preference string `json:"preference" type:"string"` 
	
}

func NewOrderCanceled() *OrderCanceled{
	event := &OrderCanceled{EventType:"OrderCanceled", TimeStamp:time.Now().String()}

	return event
}
