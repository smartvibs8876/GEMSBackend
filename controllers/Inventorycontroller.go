package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var inventory []entity.Inventory
	if err := database.Connector.Find(&inventory).Error; err != nil {
		fmt.Println("Internal Server Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventory)
}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var product entity.Inventory
	database.Connector.Where("product_id = ?", key).First(&product)
	fmt.Println(product.Product_id)
	w.Header().Set("Content-Type", "application/json")
	if product.Product_id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(false)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}

}

func AddProducts(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Inventory
	json.Unmarshal(requestBody, &product)
	database.Connector.Create(product)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var product entity.Inventory
	fmt.Println(key)
	id, _ := strconv.ParseInt(key, 10, 64)
	fmt.Println(id)
	database.Connector.Where("product_id = ?", id).Delete(&product)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("deleted from database")

}

func UpdateProductsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]


	requestBody, _ := ioutil.ReadAll(r.Body)
	var product entity.Inventory
	json.Unmarshal(requestBody, &product)
	database.Connector.Model(entity.Inventory{}).Where("product_id= ?", key).Updates(entity.Inventory{Name: product.Name, Description: product.Description, Price: product.Price, Quantity: product.Quantity, ImageURL: product.ImageURL})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
