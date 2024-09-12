package product_class

// Product: ID, Name, Description, Price, StockQuantity.

type ProductInfo struct {
	ProductId          uint `gorm:"primaryKey"`
	ProductName        string
	ProductDescription string
	ProductPrice       uint
	ProductInStock     uint
}
