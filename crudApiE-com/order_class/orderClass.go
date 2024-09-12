package order_class

import (
	"github.com/neihon/crud-api-with-authentication/user_class"
	"gorm.io/gorm"
	"time"
)

// Order: ID, UserID (Foreign Key), OrderDate, Status, TotalAmount.
// OrderItem: ID, OrderID (Foreign Key), ProductID (Foreign Key), Quantity, Price.

type Order struct {
	OrderId         uint `gorm:"primaryKey"`
	UserId          uint
	User            user_class.User
	ProductId       uint
	Product         string
	OrderDate       time.Time
	OrderQuantity   int
	OrderTotalPrice float64
	OrderCreatedAt  time.Time
	OrderUpdatedAt  time.Time
	OrderDeletedAt  gorm.DeletedAt `gorm:"index"`
}
