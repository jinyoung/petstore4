package store

type TopFood struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	Fooid int `json:"fooid" type:"int"`
	Count int `json:"count" type:"int"`
}