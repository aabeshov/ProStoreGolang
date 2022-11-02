package model

import "time"

type Product struct {
	Id          uint64    `gorm:"primary_key;autoincrementincrement;NOT NULL" json:"id"`
	Name        string    `gorm:"type:varchar(50)" json:"name"`
	Description string    `gorm:"type:varchar(100)" json:"description"`
	Category    Category  `gorm:"foreign_key:CategoryID" json:"-"`
	CategoryID  uint      `gorm:"not_null" json:"category_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
