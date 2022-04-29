package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	handler := http.HandlerFunc(GetAllProducts)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "Products Fetched")
}

// func TestGetProductByID(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/products/{1}", nil)
// 	fmt.Println(req.RequestURI)
// 	handler := http.HandlerFunc(GetProductByID)
// 	res := httptest.NewRecorder()
// 	handler.ServeHTTP(res, req)
// 	assert.Equal(t, 200, res.Code, "Product Fetched")
// }
