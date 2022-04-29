package DAO

import (
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"time"
)

var MakeOrderError bool

type Cart struct {
	Product_id int `json:"product_id"`
	Quantity   int `json:"quantity"`
	Price      int `json:"price"`
}
type RequestBodyForOrder struct {
	Cart    []Cart `json:"cart"`
	Address string `json:"address"`
}

func MakeOrderTable(order entity.Orders, req RequestBodyForOrder) {
	order.Address = req.Address
	order.Order_Date = time.Now().Format("01-02-2006")
	order.Delivery_Date = time.Now().AddDate(0, 0, 10).Format("01-02-2006")
	order.Status = 0
	if err := database.Connector.Create(order).Error; err != nil {
		MakeOrderError = true
	}
}

func MakeUsersOrdersTable(order entity.Orders, user entity.Users) {
	var users_orders entity.UsersOrders
	users_orders.Order_id = order.Order_id
	users_orders.User_id = user.User_id
	if err := database.Connector.Create(users_orders).Error; err != nil {
		MakeOrderError = true
	}
}

func MakeProductsOrdersTable(order entity.Orders, req RequestBodyForOrder) int {
	total_amount := 0
	for i := 0; i < len(req.Cart); i++ {
		var products_order entity.ProductsOrders
		products_order.Order_id = order.Order_id
		products_order.Product_id = req.Cart[i].Product_id
		var product entity.Inventory
		database.Connector.Where("product_id = ?", req.Cart[i].Product_id).First(&product)
		if product.Quantity < req.Cart[i].Quantity {
			MakeOrderError = true
			product.Quantity = product.Quantity + req.Cart[i].Quantity
		}
		product.Quantity = product.Quantity - req.Cart[i].Quantity
		database.Connector.Model(&product).Where("product_id = ?", req.Cart[i].Product_id).Update("Quantity", product.Quantity)
		products_order.Quantity = req.Cart[i].Quantity
		total_amount = total_amount + req.Cart[i].Quantity*req.Cart[i].Price
		if err := database.Connector.Create(products_order).Error; err != nil {
			MakeOrderError = true
		}
	}
	return total_amount
}

func MakeInvoiceTable(order entity.Orders, total_amount int) {
	var invoice entity.Invoice
	invoice.Invoice_date = time.Now().Format("01-02-2006")
	invoice.Order_id = order.Order_id
	invoice.Total_amount = total_amount
	if err := database.Connector.Create(invoice).Error; err != nil {
		MakeOrderError = true
	}
}
