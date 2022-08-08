package store
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type TopFoodDB struct{
	db *gorm.DB
}

var topFoodrepository *TopFoodDB

func TopFoodDBInit() {
	var err error
	topFoodrepository = &TopFoodDB{}
	topFoodrepository.db, err = gorm.Open(sqlite.Open("TopFood_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	topFoodrepository.db.AutoMigrate(&TopFood{})

}

func TopFoodRepository() *TopFoodDB {
	return topFoodrepository
}

func (self *TopFoodDB)save(entity *TopFood) error {
	tx := self.db.Create(entity)
	if tx.Error != nil {
		tx1 := self.db.Save(entity)
		if tx1.Error != nil {
			log.Print(tx1.Error)
			return tx1.Error
		}
	}
	return nil
}

func (self *TopFoodDB)GetList() []TopFood{
	
	entities := []TopFood{}
	self.db.Find(&entities)

	return entities
}

func (self *TopFoodDB)GetID(id int) *TopFood{
	entity := &TopFood{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *TopFoodDB) Delete(entity *TopFood) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *TopFoodDB) Update(id int, params map[string]string) error{
	entity := &TopFood{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	}else {
		update := &TopFood{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}