package entity

import (
	_ "gorm.io/gorm"
)

type Invoice struct {
	Invoice_id   int    //`gorm:"primary_key"`
	Total_amount int    //`gorm:"not null"`
	Invoice_date string //`gorm:"not null"`
	Order_id     int    //`gorm:"not null"`
}
