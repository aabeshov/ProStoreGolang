package model

type User struct {
	Id       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Login    string `gorm:"type:varchar(50)" json:"login"`
	Name     string `gorm:"type:varchar(50)" json:"name"`
	Password string `json:"password"`
}
