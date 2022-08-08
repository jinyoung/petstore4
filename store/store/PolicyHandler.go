package store

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverPaid_주문목록에추가(data map[string]interface{}){
	
	event := NewPaid()
	mapstructure.Decode(data,&event)
	storeOrder := &StoreOrder{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	storeOrderrepository.save(storeOrder)
	menu := &Menu{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	menurepository.save(menu)
	
}

func wheneverOrderCanceled_주문취소알림(data map[string]interface{}){
	
	event := NewOrderCanceled()
	mapstructure.Decode(data,&event)
	storeOrder := &StoreOrder{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	storeOrderrepository.save(storeOrder)
	menu := &Menu{}
	// TODO Set value from event ( ex: delivery.OrderId = event.Id )
	// TODO Change according to the event (save, delete, put..)
	menurepository.save(menu)
	
}

