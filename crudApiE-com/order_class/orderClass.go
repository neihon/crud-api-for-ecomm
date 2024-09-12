package order_class

import (
	"github.com/neihon/crud-api-with-authentication/product_class"
	"github.com/neihon/crud-api-with-authentication/user_class"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	OrderId         uint `gorm:"primaryKey"`
	UserId          uint
	User            user_class.User
	ProductId       uint
	Product         product_class.Product
	OrderDate       time.Time
	OrderQuantity   int
	OrderTotalPrice float64
	OrderCreatedAt  time.Time
	OrderUpdatedAt  time.Time
	OrderDeletedAt  gorm.DeletedAt `gorm:"index"`
}
