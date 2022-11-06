package model

type Cart struct {
	Product   Product `gorm:"foreignkey:ProductID" json:"-"`
	ProductID uint    `gorm:"NOT NULL" json:"product_id"`
	User      User    `gorm:"foreignkey:UserID"`
	UserID    uint    `gorm:"NOT NULL" json:"used_id"`
	Quantity  uint    `gorm:"type:integer;default:0" json:"quantity"`
}
