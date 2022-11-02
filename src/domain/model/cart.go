package model

type Cart struct {
	Product   Product `gorm:"foreignkey:ProductID" json:"product_id"`
	ProductID uint    `gorm:"foreignkey:ProductID" json:"p"`
	User      User    `gorm:"foreignkey:UserID"`
	UserID    uint    `json:"used_id"`
	Quantity  uint    `gorm:"default:0" json:"quantity"`
}
