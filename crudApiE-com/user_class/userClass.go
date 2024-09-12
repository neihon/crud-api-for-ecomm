package user_class

import (
	"gorm.io/gorm"
	"time"
)

// User: ID, Name, Email, ShippingAddress.
// Order: ID, UserID (Foreign Key), OrderDate, Status, TotalAmount.
// OrderItem: ID, OrderID (Foreign Key), ProductID (Foreign Key), Quantity, Price.

type User struct {
	UserId              uint `gorm:"primaryKey"`
	UserName            string
	UserEmail           string
	UserShippingAddress string
	UserCreatedAt       time.Time
	UserUpdatedAt       time.Time
	UserDeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UserOrderList struct {
	// no-op
}

type Order struct {
	// no-op
	OrderId   uint `gorm:"primaryKey"`
	User      User
	ProductId product_class.ProductId
}
