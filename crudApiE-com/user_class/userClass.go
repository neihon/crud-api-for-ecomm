package user_class

import (
	"gorm.io/gorm"
	"time"
)

// User: ID, Name, Email, ShippingAddress.

type User struct {
	userId              uint `gorm:"primaryKey"`
	userName            string
	userEmail           string
	userShippingAddress string
	userCreatedAt       time.Time
	userUpdatedAt       time.Time
	userDeletedAt       gorm.DeletedAt `gorm:"index"`
}
