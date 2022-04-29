package entity

type ProductsOrders struct {
	Product_id int //`gorm:"primary_key;auto_Increment:false"`
	Order_id   int //`gorm:"primary_key;auto_Increment:false"`
	Quantity   int //`gorm:"not null"`
}
