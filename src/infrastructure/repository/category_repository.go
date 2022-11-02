package repository

import (
	model "GolangwithFrame/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type CategoryRepository interface {
	CreateProduct(category model.Category)
	UpdateProduct(category model.Category) error
	DeleteProduct(category model.Category) error
	FindAllProducts() []model.Category
	GetProduct(id int) (model.Category, error)
	//Get(product model.Product) model.Product
}

//func NewRepository() ProductRepository {
//
//	db := NewDB()
//	return &database{
//		connection: db,
//	}
//
//}

func (db *database) GetCategory(id int) (model.Category, error) {
	category := model.Category{}
	err := db.connection.Where("id = ?", id).First(&category).Error
	if err != nil {
		return category, err
	} else {
		return category, nil
	}
	//db.connection.First()
	//db.connection.First(&prod, "id=?", product.Id)
	////db.connection.Where("name = ?", "jinzhu").First(&product)
	//return db.connection.First(&prod, "id=?", product.Id), nil
}
