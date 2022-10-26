package model

import "time"

type Product struct {
	Id          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"type:varchar(50)" json:"name"`
	Description string    `gorm:"type:varchar(100)" json:"description"`
	Category    Category  `gorm:"foreignkey:CategoryID"json:"category_id"`
	CategoryID  uint      `json:"-"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
