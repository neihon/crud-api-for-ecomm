package sqliteDb

import (
	"github.com/neihon/crud-api-with-authentication/order_class"
	"github.com/neihon/crud-api-with-authentication/product_class"
	"github.com/neihon/crud-api-with-authentication/user_class"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func CreateDatabase() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("crudApiE-com/sqliteDb/apiDB.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func DbMigrateModels(database *gorm.DB) {
	err := database.AutoMigrate(
		&user_class.User{},
		&product_class.Product{},
		&order_class.Order{},
		&order_class.OrderItem{},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(database *gorm.DB, user user_class.User) user_class.User {
	database.Create(&user)
	return user
}
