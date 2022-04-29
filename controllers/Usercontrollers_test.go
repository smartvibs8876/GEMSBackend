package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistration(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"F_name":   "Yusuf",
		"L_name":   "pathan",
		"Email":    "yusuf@gmail.com",
		"Password": "yusuf",
		"Mo_no":    "3123131132",
		"Address":  "Punjab",
	})
	requestBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "/user/registration", requestBody)
	handler := http.HandlerFunc(Registration)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	assert.Equal(t, 409, res.Code, "OK record created")

	postBody2, _ := json.Marshal(map[string]string{
		"F_name":   "",
		"L_name":   "pathan",
		"Email":    "yusuf@gmail.com",
		"Password": "yusuf",
		"Mo_no":    "3123131132",
		"Address":  "Punjab",
	})
	requestBody2 := bytes.NewBuffer(postBody2)
	req2, _ := http.NewRequest("POST", "/user/registration", requestBody2)
	handler2 := http.HandlerFunc(Registration)
	res2 := httptest.NewRecorder()
	handler2.ServeHTTP(res2, req2)
	assert.Equal(t, 204, res2.Code, "OK record created")
}

func TestLogin(t *testing.T) {
	postBody, _ := json.Marshal(map[string]string{
		"Email":    "yusuf@gmail.com",
		"Password": "yusuf",
	})
	requestBody := bytes.NewBuffer(postBody)
	req, _ := http.NewRequest("POST", "/user/login", requestBody)
	handler := http.HandlerFunc(Login)
	res := httptest.NewRecorder()
	handler.ServeHTTP(res, req)
	assert.Equal(t, 200, res.Code, "OK record created")

	postBody2, _ := json.Marshal(map[string]string{
		"Email":    "yusuf1@gmail.com",
		"Password": "yusuf",
	})
	requestBody2 := bytes.NewBuffer(postBody2)
	req2, _ := http.NewRequest("POST", "/user/login", requestBody2)
	handler2 := http.HandlerFunc(Login)
	res2 := httptest.NewRecorder()
	handler2.ServeHTTP(res2, req2)
	assert.Equal(t, 401, res2.Code, "OK record created")

	postBody3, _ := json.Marshal(map[string]string{
		"Email":    "yusuf@gmail.com",
		"Password": "yusuf1",
	})
	requestBody3 := bytes.NewBuffer(postBody3)
	req3, _ := http.NewRequest("POST", "/user/login", requestBody3)
	handler3 := http.HandlerFunc(Login)
	res3 := httptest.NewRecorder()
	handler3.ServeHTTP(res3, req3)
	assert.Equal(t, 401, res3.Code, "OK record created")
}
