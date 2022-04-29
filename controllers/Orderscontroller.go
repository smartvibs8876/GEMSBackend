package controllers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"rest-go-demo/DAO"
	"rest-go-demo/database"
	"rest-go-demo/entity"
)

func FindIfIDExists(ordersCheck []entity.Orders, id int) bool {
	for i := 0; i < len(ordersCheck); i++ {
		if id == ordersCheck[i].Order_id {
			return true
		}
	}
	return false
}

type OrdersAndProducts struct {
	Order_id      int                `json:"order_id"`
	Order_Date    string             `json:"order_date"`
	Delivery_Date string             `json:"delivery_date"`
	Status        int                `json:"status"`
	Address       string             `json:"address"`
	ProductItems  []entity.Inventory `json:"product_items"`
}

func MakeAnOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := Authorization(w, r)
	if (user == entity.Users{}) {
		return
	}
	requestBody, _ := ioutil.ReadAll(r.Body)
	var req DAO.RequestBodyForOrder
	json.Unmarshal(requestBody, &req)
	var ordersCheck []entity.Orders
	if err := database.Connector.Find(&ordersCheck).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(false)
		return
	}
	var order entity.Orders
	order.Order_id = rand.Intn(10000) + 1
	for FindIfIDExists(ordersCheck, order.Order_id) {
		order.Order_id = (order.Order_id + 1) % 10000
	}
	DAO.MakeOrderTable(order, req)
	DAO.MakeUsersOrdersTable(order, user)
	total_amount := DAO.MakeProductsOrdersTable(order, req)
	DAO.MakeInvoiceTable(order, total_amount)
	if DAO.MakeOrderError == false {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(order.Order_id)
		return
	} else {
		database.Connector.Where("order_id = ?", order.Order_id).Delete(&entity.UsersOrders{})
		database.Connector.Where("order_id = ?", order.Order_id).Delete(&entity.ProductsOrders{})
		database.Connector.Where("order_id = ?", order.Order_id).Delete(&entity.Invoice{})
		database.Connector.Where("order_id = ?", order.Order_id).Delete(&entity.Orders{})
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(false)
		return
	}

}

// func OrdersByAUser(w http.ResponseWriter, r *http.Request) {
// 	user := Authorization(w, r)
// 	var ordersAndProducts []OrdersAndProducts
// 	var users_orders []entity.UsersOrders
// 	database.Connector.Where("user_id = ?", user.User_id).Find(&users_orders)
// 	var orders []entity.Orders
// 	for i := 0; i < len(users_orders); i++ {
// 		var order entity.Orders
// 		database.Connector.Where("order_id = ?", users_orders[i].Order_id).Find(&order)
// 		orders = append(orders, order)
// 	}
// 	var products_orders [][]entity.ProductsOrders
// 	for i := 0; i < len(users_orders); i++ {
// 		var products_order []entity.ProductsOrders
// 		database.Connector.Where("order_id = ?", users_orders[i].Order_id).Find(&products_order)
// 		products_orders = append(products_orders, products_order)
// 	}
// 	for i := 0; i < len(orders); i++ {
// 		var orderAndProduct OrdersAndProducts
// 		orderAndProduct.Order_id = orders[i].Order_id
// 		orderAndProduct.Order_Date = orders[i].Order_Date
// 		orderAndProduct.Delivery_Date = orders[i].Delivery_Date
// 		orderAndProduct.Status = orders[i].Status
// 		orderAndProduct.Address = orders[i].Address
// 		for j := 0; j < len(products_orders[i]); j++ {
// 			var obj entity.Inventory
// 			database.Connector.Where("product_id = ?", products_orders[i][j].Product_id).Find(&obj)
// 			obj.Quantity = products_orders[i][j].Quantity
// 			orderAndProduct.ProductItems = append(orderAndProduct.ProductItems, obj)
// 		}

// 		ordersAndProducts = append(ordersAndProducts, orderAndProduct)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(ordersAndProducts)
// }

func GetUsersOrders(orders []entity.Orders, users_orders []entity.UsersOrders) ([]entity.UsersOrders, []entity.Orders) {
	for i := 0; i < len(users_orders); i++ {
		var order entity.Orders
		database.Connector.Where("order_id = ?", users_orders[i].Order_id).Find(&order)
		orders = append(orders, order)
	}
	return users_orders, orders
}

func GetProductsOrders(users_orders []entity.UsersOrders) [][]entity.ProductsOrders {
	var products_orders [][]entity.ProductsOrders
	for i := 0; i < len(users_orders); i++ {
		var products_order []entity.ProductsOrders
		database.Connector.Where("order_id = ?", users_orders[i].Order_id).Find(&products_order)
		products_orders = append(products_orders, products_order)
	}
	return products_orders
}

func GetOrdersAndProducts(orders []entity.Orders, products_orders [][]entity.ProductsOrders) []OrdersAndProducts {
	var ordersAndProducts []OrdersAndProducts
	for i := 0; i < len(orders); i++ {
		var orderAndProduct OrdersAndProducts
		orderAndProduct.Order_id = orders[i].Order_id
		orderAndProduct.Order_Date = orders[i].Order_Date
		orderAndProduct.Delivery_Date = orders[i].Delivery_Date
		orderAndProduct.Status = orders[i].Status
		orderAndProduct.Address = orders[i].Address
		for j := 0; j < len(products_orders[i]); j++ {
			var obj entity.Inventory
			database.Connector.Where("product_id = ?", products_orders[i][j].Product_id).Find(&obj)
			obj.Quantity = products_orders[i][j].Quantity
			orderAndProduct.ProductItems = append(orderAndProduct.ProductItems, obj)
		}
		ordersAndProducts = append(ordersAndProducts, orderAndProduct)
	}
	return ordersAndProducts
}
func OrdersByAUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := Authorization(w, r)
	var orders []entity.Orders
	var users_orders []entity.UsersOrders
	database.Connector.Where("user_id = ?", user.User_id).Find(&users_orders)
	users_orders, orders = GetUsersOrders(orders, users_orders)
	products_orders := GetProductsOrders(users_orders)
	ordersAndProducts := GetOrdersAndProducts(orders, products_orders)
	json.NewEncoder(w).Encode(ordersAndProducts)
}
