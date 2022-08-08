package store
import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type MenuDB struct{
	db *gorm.DB
}

var menurepository *MenuDB

func MenuDBInit() {
	var err error
	menurepository = &MenuDB{}
	menurepository.db, err = gorm.Open(sqlite.Open("Menu_table.db"), &gorm.Config{})
	
	if err != nil {
		panic("DB Connection Error")
	}
	menurepository.db.AutoMigrate(&Menu{})

}

func MenuRepository() *MenuDB {
	return menurepository
}

func (self *MenuDB)save(entity interface{}) error {
	
	tx := self.db.Create(entity)

	if tx.Error != nil {
		log.Print(tx.Error)
		return tx.Error
	}
	return nil
}

func (self *MenuDB)GetList() []Menu{
	
	entities := []Menu{}
	self.db.Find(&entities)

	return entities
}

func (self *MenuDB)GetID(id int) *Menu{
	entity := &Menu{}
	self.db.Where("id = ?", id).First(entity)

	return entity
}

func (self *MenuDB) Delete(entity *Menu) error{
	err2 := self.db.Delete(&entity).Error
	return err2
}

func (self *MenuDB) Update(id int, params map[string]string) error{
	entity := &Menu{}
	err1 := self.db.Where("id = ?", id).First(entity).Error
	if err1 != nil {
		return err1
	}else {
		update := &Menu{}
		ObjectMapping(update, params)

		err2 := self.db.Model(&entity).Updates(update).Error
		return err2
	}

}