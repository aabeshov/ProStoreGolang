package repository

import (
	model "GolangwithFrame/src/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin12345"
	dbname   = "postgres"
)

type ProductRepository interface {
	Save(product model.Product)
	Update(product model.Product)
	Delete(product model.Product)
	FindAll() []model.Product
	CloseDB()
	Get(id int) (model.Product, error)
	//Get(product model.Product) model.Product
}

type database struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("Can't connect to database")
	}
	db.AutoMigrate(&model.Token{}, &model.User{}, &model.Product{}, &model.Category{})
	return &database{
		connection: db,
	}

}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to Connect")
	}

}

func (db *database) Get(id int) (model.Product, error) {
	product := model.Product{}
	db.connection.Where("id = ?", id).First(&product)
	//if err != nil {
	//	panic("error")
	//}
	return product, nil
	//db.connection.First()
	//db.connection.First(&prod, "id=?", product.Id)
	////db.connection.Where("name = ?", "jinzhu").First(&product)
	//return db.connection.First(&prod, "id=?", product.Id), nil
}

func (db *database) Save(product model.Product) {
	db.connection.Create(&product)
}
func (db *database) Update(product model.Product) {
	db.connection.Save(&product)
}
func (db *database) Delete(product model.Product) {
	db.connection.Delete(&product)
}
func (db *database) FindAll() []model.Product {
	var products []model.Product
	db.connection.Set("gorm:auto_preload", true).Find(&products)
	return products
}

//
//func (db *database) Get(id uint64) {
//	var product model.Product
//	return
//}
