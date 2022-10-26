package model

type Category struct {
	Id   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(50)" json:"name"`
}
