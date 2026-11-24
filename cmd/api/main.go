package main

import (
	"log"
	"submitto/internals/database"
	"submitto/internals/models"
)

func main(){


	db,err:= database.Connect()

	if err!=nil{
		log.Fatal("Database Connection failed:",err)
	}

	log.Println("Database Connection successful")
	

	err= db.AutoMigrate(
		&models.Form{},
		&models.FormField{},
		&models.Submission{},
		&models.SubmissionAnswer{},

	
	)

	 if err != nil {
        log.Fatal("Database migration failed:", err)
    }

    log.Println("Database migration successful")

	
      
}