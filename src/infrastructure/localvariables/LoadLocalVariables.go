package localvariables

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

func LoadEnvVariables() {

	dir, err := os.Getwd()
	fmt.Println(dir)

	err = godotenv.Load(filepath.Join(dir, ".env"))
	//err := godotenv.Load()

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Error loading .env file!")
	}
}
