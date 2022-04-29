package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"rest-go-demo/DAO"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAnOrder(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "yusuf@gmail.com",
		"Password": "yusuf",
	})
	requestBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "/user/login", requestBody)
	handler := http.HandlerFunc(Login)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	bodyBytes, _ := io.ReadAll(res.Body)
	//assert.Equal(t, 200, res.Code, "OK logged in successfully")
	tokenString := string(bodyBytes)

	var reqBodyForOrder DAO.RequestBodyForOrder
	var cartItem DAO.Cart
	cartItem.Product_id = 1
	cartItem.Price = 50000
	cartItem.Quantity = 2
	reqBodyForOrder.Cart = append(reqBodyForOrder.Cart, cartItem)
	cartItem.Product_id = 2
	cartItem.Price = 250
	cartItem.Quantity = 3
	reqBodyForOrder.Cart = append(reqBodyForOrder.Cart, cartItem)
	reqBodyForOrder.Address = "Fatorda"
	postBody2, _ := json.Marshal(reqBodyForOrder)
	requestBody2 := bytes.NewBuffer(postBody2)
	req2, _ := http.NewRequest("POST", "/order", requestBody2)
	req2.Header.Add("Authorization", "Bearer "+tokenString[0:len(tokenString)-1])
	handler2 := http.HandlerFunc(MakeAnOrder)
	res2 := httptest.NewRecorder()
	handler2.ServeHTTP(res2, req2)
	assert.Equal(t, 200, res.Code, "Order placed successfully")
}

func TestOrdersByAUser(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "yusuf@gmail.com",
		"Password": "yusuf",
	})
	requestBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "/user/login", requestBody)
	handler := http.HandlerFunc(Login)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	bodyBytes, _ := io.ReadAll(res.Body)
	//assert.Equal(t, 200, res.Code, "OK logged in successfully")
	tokenString := string(bodyBytes)

	req2, _ := http.NewRequest("GET", "/order", nil)
	req2.Header.Add("Authorization", "Bearer "+tokenString[0:len(tokenString)-1])
	handler2 := http.HandlerFunc(OrdersByAUser)
	res2 := httptest.NewRecorder()
	handler2.ServeHTTP(res2, req2)
	assert.Equal(t, 200, res2.Code, "Orders Fetched")
}
