package user_class

import (
	"github.com/neihon/crud-api-with-authentication/order_class"
	"gorm.io/gorm"
	"time"
)

// User: ID, Name, Email, ShippingAddress.

type User struct {
	UserId              uint `gorm:"primaryKey"`
	UserName            string
	UserEmail           string
	UserShippingAddress string
	UserOrders          []order_class.Order
	UserCreatedAt       time.Time
	UserUpdatedAt       time.Time
	UserDeletedAt       gorm.DeletedAt `gorm:"index"`
}
