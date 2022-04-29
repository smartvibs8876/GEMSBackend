package entity

import (
	_ "gorm.io/gorm"
)

type Orders struct {
	Order_id      int    //`gorm:"primary_key;auto_Increment:false"`
	Order_Date    string //`gorm:"not null"`
	Delivery_Date string //`gorm:"not null"`
	Status        int    //`gorm:"not null"`
	Address       string //`gorm:"not null"`
}
