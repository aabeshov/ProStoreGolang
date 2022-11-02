package repository

import (
	"GolangwithFrame/src/domain/model"
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin12345"
	dbname   = "postgres"
)

type database struct {
	connection *gorm.DB
}

func NewRepository() ProductRepository {

	db := NewDB()
	return &database{
		connection: db,
	}

}

func NewDB() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic("Can't connect to database")
	}
	db.AutoMigrate(&model.Token{}, &model.User{}, &model.Product{}, &model.Category{}, &model.Cart{})

	return db
}
