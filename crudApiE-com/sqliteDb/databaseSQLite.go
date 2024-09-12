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

func DbMigrateModels(database *gorm.DB) error {
	err := database.AutoMigrate(
		&user_class.User{},
		&product_class.Product{},
		&order_class.Order{},
	)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CreateUser(database *gorm.DB, user user_class.User) user_class.User {
	userCreated := database.Create(&user)
	if userCreated.Error != nil {
		log.Fatal(userCreated.Error)
	}
	return user
}

func GetUserById(database *gorm.DB, userId uint) (user_class.User, error) {
	var user user_class.User
	result := database.First(&user, userId)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}
