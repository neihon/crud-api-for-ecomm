package main

import (
	"github.com/neihon/crud-api-with-authentication/user_class"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	// no-op
	database, err := gorm.Open(sqlite.Open("apiDB.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	User := user_class.User
	database.AutoMigrate(&User{})

}
