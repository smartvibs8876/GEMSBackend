package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func setHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	initDB()
	router := mux.NewRouter()
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", setHeaders(router)))
}

func initaliseHandlers(router *mux.Router) {
	//API for users
	router.HandleFunc("/user/registration", controllers.Registration).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")
	router.HandleFunc("/user", controllers.GetUserDetailsWithToken).Methods("GET")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/orders", controllers.MakeAnOrder).Methods("POST")
	router.HandleFunc("/orders", controllers.OrdersByAUser).Methods("GET")
	router.HandleFunc("/invoice/{id}", controllers.GetInvoiceByOrderId).Methods("GET")
	//API for inventory
	router.HandleFunc("/products", controllers.AddProducts).Methods("POST")
	router.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
	router.HandleFunc("/products/{id}", controllers.DeleteProductByID).Methods("DELETE")
	router.HandleFunc("/products/{id}", controllers.UpdateProductsByID).Methods("PUT")
}

func initDB() {

	// config :=
	// 	database.Config{
	// 		ServerName: "169.51.195.241:32000",
	// 		User:       "admin",
	// 		Password:   "secret",
	// 		DB:         "mysqldb",
	// 	}
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "vibhavdhaimode123",
			DB:         "golang_api",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Users{}, &entity.Inventory{}, &entity.Orders{}, &entity.Invoice{}, &entity.UsersOrders{}, &entity.ProductsOrders{})
}

//Queries for iventory
// insert into inventories(name,description,price,quantity,image_url)
// values ("Laptop","Intel i7 11th gen 512 GB SSD",50000,60,"../../../.././assets/Laptop.jpg");
// insert into inventories(name,description,price,quantity,image_url)
// values ("Mouse","Wireless Dell",250,75,"../../../.././assets/Mouse.gopng");
// insert into inventories(name,description,price,quantity,image_url)
// values ("Iphone","64 GB Rom 16 GB RAM",75000,100,"../../../.././assets/Phone.png");
// insert into inventories(name,description,price,quantity,image_url)
// values ("Keyboard","Wirelss Lenovo",850,120,"../../../.././assets/Keyboard.png");
