package controllers

import (
	"os"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestMain(m *testing.M) {
	initDB()
	os.Exit(m.Run())
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
