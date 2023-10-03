package handlers

// import (
// 	"julianmindria/backendCoffee/config"
// 	"julianmindria/backendCoffee/internal/models"
// 	"julianmindria/backendCoffee/internal/repositories"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var repoProductMock = repositories.RepoProductMock{}
// var reqBodyProduct = `{
// 	"id_Product": "123",
// 	"product_name": "test",
// 	"description": "testtt",
// 	"stock": 2,
// 	"price": 1000,
// 	"categories": "food",
// 	"IsFavorite": true,
// 	"image": "file.jpg"
// }`

// func TestPostProduct(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	w := httptest.NewRecorder()
// 	// r := gin.Default()
// 	a, r := gin.CreateTestContext(w)
// 	a.Set("images", "files.jpg")

// 	handler := NewProduct(&repoProductMock)
// 	expectedResult := &config.Result{Message: "200"}
// 	repoProductMock.On("CreateProduct", mock.Anything).Return(expectedResult, nil)

// 	r.POST("/create", handler.Postdata)
// 	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBodyProduct))
// 	req.Header.Set("Content-type", "application/json")
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"data": "Product created"}`, w.Body.String())
// }

// func TestDeleteProduct(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	w := httptest.NewRecorder()
// 	r := gin.Default()

// 	handler := NewProduct(&repoProductMock)
// 	expectedResult := &config.Result{Data: "Delete Success"}
// 	repoProductMock.On("DeleteProduct", mock.Anything).Return(expectedResult, nil)

// 	r.DELETE("/delete/:id", handler.Deletedata)
// 	req := httptest.NewRequest("DELETE", "/delete/123", strings.NewReader("{}"))
// 	req.Header.Set("Content-type", "application/json")
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"data": "Delete Success"}`, w.Body.String())
// }

// func TestUpdateProduct(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	w := httptest.NewRecorder()
// 	// r := gin.Default()
// 	a, r := gin.CreateTestContext(w)
// 	a.Set("image", "file.jpg")

// 	handler := NewProduct(&repoProductMock)
// 	expectedResult := &config.Result{Data: "Update Success"}
// 	repoProductMock.On("UpdateProduct", mock.Anything).Return(expectedResult, nil)

// 	r.PUT("/update/:username", handler.Updatedata)
// 	req := httptest.NewRequest("PUT", "/update/julianmindria", strings.NewReader(reqBodyProduct))
// 	req.Header.Set("Content-type", "application/json")
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"data": "Update Success"}`, w.Body.String())
// }

// func TestGetProduct(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	w := httptest.NewRecorder()
// 	r := gin.Default()

// 	handler := NewProduct(&repoProductMock)
// 	expectedResult := []models.User{{Username: "mindriajulian", Role: "user", Created_at: nil, Updated_at: nil}}
// 	repoProductMock.On("Get_Users", mock.Anything).Return(expectedResult, nil)

// 	r.GET("/getuser", handler.Getdata)
// 	req := httptest.NewRequest("GET", "/getuser", strings.NewReader(reqBodyProduct))
// 	req.Header.Set("Content-type", "application/json")
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{
// 		"data": [
// 			{
// 				"username": "mindriajulian",
// 				"roles": "user",
// 				"created_at": null,
// 				"updated_at": null
// 			}
// 		],
// 		"description": "OK",
// 		"status": 200
// 	}`, w.Body.String())
// }
