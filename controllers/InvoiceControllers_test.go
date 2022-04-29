package controllers

// func buildRequest(method string, url string, doctype uint32, docid uint32) *http.Request {
// 	req, _ := http.NewRequest(method, url, nil)
// 	req.ParseForm()
// 	var vars = map[string]string{
// 		"doctype": strconv.FormatUint(uint64(doctype), 10),
// 		"docid":   strconv.FormatUint(uint64(docid), 10),
// 	}
// 	context.DefaultContext.Set(req, mux.ContextKey(0), vars) // mux.ContextKey exported
// 	return req
// }
// func TestGetInvoiceByOrderId(t *testing.T) {
// 	postBody, _ := json.Marshal(map[string]string{
// 		"Email":    "yusuf@gmail.com",
// 		"Password": "yusuf",
// 	})
// 	requestBody := bytes.NewBuffer(postBody)
// 	req, _ := http.NewRequest("POST", "/user/login", requestBody)
// 	handler := http.HandlerFunc(Login)
// 	res := httptest.NewRecorder()
// 	handler.ServeHTTP(res, req)
// 	bodyBytes, _ := io.ReadAll(res.Body)
// 	assert.Equal(t, 200, res.Code, "OK logged in successfully")
// 	tokenString := string(bodyBytes)

// 	req2, _ := http.NewRequest("GET", "/orders/1848", nil)
// 	req2.Header.Add("Authorization", "Bearer "+tokenString[0:len(tokenString)-1])
// 	handler2 := http.HandlerFunc(GetInvoiceByOrderId)
// 	res2 := httptest.NewRecorder()
// 	handler2.ServeHTTP(res2, req2)
// 	fmt.Println(res2.Code)
// 	assert.Equal(t, 200, res.Code, "Order Fetched")
// }
