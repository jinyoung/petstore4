package store
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type StoreOrderDB struct{
	db *gorm.DB
}

var storeOrderrepository *StoreOrderDB

func StoreOrderDBInit() {
	var err error
	storeOrderrepository = &StoreOrderDB{}
	storeOrderrepository.db, err = gorm.Open(sqlite.Open("StoreOrder_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	storeOrderrepository.db.AutoMigrate(&StoreOrder{})

}

func StoreOrderRepository() *StoreOrderDB {
	return storeOrderrepository
}

func (self *StoreOrderDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *StoreOrderDB)GetList() []StoreOrder{
	
	entities := []StoreOrder{}
	self.db.Find(&entities)

	return entities
}

func (self *StoreOrderDB)GetID(id int) *StoreOrder{
	entity := &StoreOrder{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *StoreOrderDB) Delete(entity *StoreOrder) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *StoreOrderDB) Update(id int, params map[string]string) error{
	entity := &StoreOrder{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	}else {
		update := &StoreOrder{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}