package store

import (
	"github.com/mitchellh/mapstructure"
	"log"
)

func whenOrderPlaced_then_CREATE_1 (inputEvent map[string]interface{}) {

	orderPlaced := NewOrderPlaced()
	mapstructure.Decode(inputEvent, &orderPlaced)

	topFood := &TopFood{}
	topFood.Fooid = orderPlaced.FoodId
	topFood.Count = orderPlaced.직접입력

	// view 레파지 토리에 save
	repository := TopFoodRepository()
	err := repository.save(topFood)
	if err != nil {
		// TODO error control
		log.Printf("Create error: %v \n", err)
	}
}

func whenOrderPlaced_then_UPDATE_1 (inputEvent map[string]interface{}) {

	orderPlaced := NewOrderPlaced()
	mapstructure.Decode(inputEvent,&orderPlaced)

	var topFoods []TopFood
	repository := TopFoodRepository()
	// FIXME geom lib define snake_case as column name (eg: user_id), so if your query column is 'userId' then change 'user_id'
	err := repository.db.Where("fooid = ?", orderPlaced.FoodId).Find(&topFoods).Error
	if err != nil {
		// TODO error control
		log.Printf("Select error: %v \n", err)
	}
	for _, viewEntity := range topFoods {
		viewEntity.Count = orderPlaced.직접입력
		err1 := repository.db.Updates(viewEntity).Error
		if err1 != nil {
			// TODO error control
			log.Printf("Update error: %v \n", err1)
		}
	}
}

