package entity

type UsersOrders struct {
	Order_id int //`gorm:"primary_key;auto_Increment:false"`
	User_id  int //`gorm:"primary_key;auto_Increment:false"`
}
