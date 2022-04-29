package entity

import (
	_ "gorm.io/gorm"
)

//Inventory object for REST(CRUD)
type Inventory struct {
	Product_id  int    `json:"product_id"`  //gorm:"primary_key"`
	Name        string `json:"name"`        //gorm:"not null"`
	Description string `json:"description"` //gorm:"not null"`
	Price       int    `json:"price"`       //gorm:"not null"`
	Quantity    int    `json:"quantity"`    //gorm:"not null"`
	ImageURL    string `json:"imageURL"`    //gorm:"not null"`
}
