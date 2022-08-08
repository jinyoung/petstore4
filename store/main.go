package main

import(
	"log"
	"store/store"
	"store/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	store.StoreOrderDBInit()
	store.MenuDBInit()
	store.TopFoodDBInit()
	go store.InitProducer(config.GetMode())
	// view 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go store.InitConsumer(config.GetMode())
	// policy 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go store.InitConsumer(config.GetMode())
	e := store.RouteInit()

	e.Logger.Fatal(e.Start(":8082"))
}
