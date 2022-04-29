package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
)

func GetInvoiceByOrderId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := Authorization(w, r)
	if (user == entity.Users{}) {
		return
	}
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key)
	var invoiceBill entity.Invoice
	database.Connector.Where("order_id = ?", key).First(&invoiceBill)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoiceBill)

}
