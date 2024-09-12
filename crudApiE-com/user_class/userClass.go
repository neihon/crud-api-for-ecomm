package user_class

import (
	"github.com/neihon/crud-api-with-authentication/order_class"
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
	UserOrders          []order_class.Order
	UserCreatedAt       time.Time
	UserUpdatedAt       time.Time
	UserDeletedAt       gorm.DeletedAt `gorm:"index"`
}

type ProductInfo struct {
	ProductId          uint `gorm:"primaryKey"`
	ProductName        string
	ProductDescription string
	ProductPrice       uint
	ProductInStock     uint
	ProductOrders      []Order
}
