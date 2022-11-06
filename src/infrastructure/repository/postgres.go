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

type Database struct {
	Connection *gorm.DB
}

func NewRepository() Database {

	db := NewDB()
	return Database{
		Connection: db,
	}

}

func (db *Database) CloseDB() {
	err := db.Connection.Close()
	if err != nil {
		panic("Failed to Connect")
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
