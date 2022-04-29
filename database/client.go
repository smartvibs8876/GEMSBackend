package database

import (
	"fmt"
	"log"
	"rest-go-demo/entity"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func Migrate(users *entity.Users, inventory *entity.Inventory, order *entity.Orders, invoice *entity.Invoice, usersOrders *entity.UsersOrders, productOrders *entity.ProductsOrders) {
	if !Connector.HasTable("users") {
		Connector.Exec("create table users(user_id integer primary key auto_increment,f_name varchar(100) not null,l_name varchar(100) not null,password varchar(100) not null,email varchar(100) unique not null,mo_no varchar(100) not null,address varchar(100) not null);")
	}
	if !Connector.HasTable("inventories") {
		Connector.Exec("create table inventories(product_id integer primary key auto_increment ,name varchar(100) not null,description varchar(100) not null,price integer not null,quantity integer not null,image_url varchar(500) not null);")
	}
	if !Connector.HasTable("orders") {
		Connector.Exec("create table orders(order_id integer primary key ,order_date varchar(100) not null,delivery_date varchar(100) not null,status integer not null,address varchar(100) not null);")
	}
	if !Connector.HasTable("products_orders") {
		Connector.Exec("create table products_orders(product_id integer not null,order_id integer not null,quantity integer not null,primary key(product_id,order_id),foreign key (product_id) references inventories(product_id),foreign key (order_id) references orders(order_id));")
	}
	if !Connector.HasTable("users_orders") {
		Connector.Exec("create table users_orders(user_id integer not null,order_id integer not null,primary key(user_id,order_id),foreign key (user_id) references users(user_id),foreign key (order_id) references orders(order_id));")
	}
	if !Connector.HasTable("invoices") {
		Connector.Exec("create table invoices(invoice_id integer primary key auto_increment,total_amount integer not null,invoice_date varchar(100) not null,order_id integer not null,foreign key (order_id) references orders (order_id));")
	}
	Connector.AutoMigrate(&users, &inventory, &order, &invoice, &usersOrders, &productOrders)
	fmt.Println("Auto Migrated tables")
}

// func Migrate(users *entity.Users, inventory *entity.Inventory, order *entity.Orders, invoice *entity.Invoice, users_orders *entity.UsersOrders, products_orders *entity.ProductsOrders) {
// 	Connector.AutoMigrate(&users, &inventory, &order, &invoice, &users_orders, &products_orders)
// 	Connector.Model(&entity.Invoice{}).AddForeignKey("Order_id", "Orders(Order_id)", "RESTRICT", "RESTRICT")
// 	Connector.Model(&entity.ProductsOrders{}).AddForeignKey("Product_id", "Inventories(product_id)", "RESTRICT", "RESTRICT")
// 	Connector.Model(&entity.ProductsOrders{}).AddForeignKey("Order_id", "Orders(Order_id)", "RESTRICT", "RESTRICT")
// 	Connector.Model(&entity.UsersOrders{}).AddForeignKey("User_id", "Users(user_id)", "RESTRICT", "RESTRICT")
// 	Connector.Model(&entity.UsersOrders{}).AddForeignKey("Order_id", "Orders(Order_id)", "RESTRICT", "RESTRICT")
// }
