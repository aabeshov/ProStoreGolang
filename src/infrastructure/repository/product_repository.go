package repository

import (
	model "GolangwithFrame/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "admin12345"
//	dbname   = "postgres"
//)

type ProductRepository interface {
	CreateProduct(product model.Product)
	UpdateProduct(product model.Product) error
	DeleteProduct(product model.Product) error
	FindAllProducts() []model.Product
	CloseDB()
	GetProduct(id int) (model.Product, error)
	//Get(product model.Product) model.Product
}

//func NewProductRepository() ProductRepository {
//
//	db := NewDB()
//	return &database{
//		connection: db,
//	}
//
//}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to Connect")
	}

}

func (db *database) GetProduct(id int) (model.Product, error) {
	product := model.Product{}
	err := db.connection.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	} else {
		return product, nil
	}
	//db.connection.First()
	//db.connection.First(&prod, "id=?", product.Id)
	////db.connection.Where("name = ?", "jinzhu").First(&product)
	//return db.connection.First(&prod, "id=?", product.Id), nil
}

func (db *database) CreateProduct(product model.Product) {
	db.connection.Create(&product)
}
func (db *database) UpdateProduct(product model.Product) error {
	currentProduct := model.Product{}

	err := db.connection.Where("id = ?", product.Id).First(&currentProduct).Error
	//fmt.Println(err)
	//db.connection.Save(&product)
	//db.connection.Where("id = ?", product.Id).First(&product).Error
	if err != nil {
		return db.connection.Where("id = ?", product.Id).First(&currentProduct).Error
	}
	db.connection.Save(&product)
	return nil
	//return nil

}
func (db *database) DeleteProduct(product model.Product) error {
	err := db.connection.Where("id = ?", product.Id).First(&product).Error
	db.connection.Delete(&product)
	return err
}
func (db *database) FindAllProducts() []model.Product {
	var products []model.Product
	db.connection.Set("gorm:auto_preload", true).Order("id").Find(&products)
	return products
}

//
//func (db *database) Get(id uint64) {
//	var product model.Product
//	return
//}
