package user_class

import (
	"gorm.io/gorm"
	"time"
)

// User: ID, Name, Email, ShippingAddress.

type User struct {
	UserId              uint `gorm:"primaryKey"`
	UserName            string
	UserEmail           string
	UserShippingAddress string
	UserCreatedAt       time.Time
	UserUpdatedAt       time.Time
	UserDeletedAt       gorm.DeletedAt `gorm:"index"`
}
