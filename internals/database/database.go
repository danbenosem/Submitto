package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

	func Connect() (*gorm.DB, error) {

      err:=godotenv.Load()

	  dsn:=fmt.Sprintf("host=localhost user=")
	  if err!=nil{
		return nil, err		
	  }


	}