package product_class

import "github.com/neihon/crud-api-with-authentication/order_class"

// Product: ID, Name, Description, Price, StockQuantity.

type Product struct {
	ProductId          uint `gorm:"primaryKey"`
	ProductName        string
	ProductDescription string
	ProductPrice       uint
	ProductInStock     uint
	ProductOrders      []order_class.Order
}
