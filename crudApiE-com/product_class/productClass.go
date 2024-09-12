package product_class

// Product: ID, Name, Description, Price, StockQuantity.

type Product struct {
	ProductId          uint `gorm:"primaryKey"`
	ProductName        string
	ProductDescription string
	ProductPrice       uint
	ProductInStock     uint
}
