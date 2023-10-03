package handlers

import (
	"julianmindria/backendCoffee/config"
	"julianmindria/backendCoffee/internal/models"
	"julianmindria/backendCoffee/internal/repositories"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoUserMock = repositories.RepoUserMock{}
var reqBody = `{
	"id_user": "123",
	"username": "testing",
	"pass": "abcd1234",
	"roles": "user"
}`

func TestPostData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	// r := gin.Default()
	a, r := gin.CreateTestContext(w)
	a.Set("image", "file.jpg")

	handler := NewUser(&repoUserMock)
	expectedResult := &config.Result{Data: "User created"}
	repoUserMock.On("CreateUser", mock.Anything).Return(expectedResult, nil)

	r.POST("/create", handler.Postdata)
	req := httptest.NewRequest("POST", "/create", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data": "User created"}`, w.Body.String())
}

func TestDeleteData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := &config.Result{Data: "Delete Success"}
	repoUserMock.On("DeleteUser", mock.Anything).Return(expectedResult, nil)

	r.DELETE("/delete/:id", handler.Deletedata)
	req := httptest.NewRequest("DELETE", "/delete/123", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data": "Delete Success"}`, w.Body.String())
}

func TestUpdateData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	// r := gin.Default()
	a, r := gin.CreateTestContext(w)
	a.Set("image", "file.jpg")

	handler := NewUser(&repoUserMock)
	expectedResult := &config.Result{Data: "Update Success"}
	repoUserMock.On("UpdateUser", mock.Anything).Return(expectedResult, nil)

	r.PUT("/update/:username", handler.Updatedata)
	req := httptest.NewRequest("PUT", "/update/julianmindria", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"data": "Update Success"}`, w.Body.String())
}

func TestGetData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	r := gin.Default()

	handler := NewUser(&repoUserMock)
	expectedResult := []models.User{{Username: "mindriajulian", Role: "user", Created_at: nil, Updated_at: nil}}
	repoUserMock.On("Get_Users", mock.Anything).Return(expectedResult, nil)

	r.GET("/getuser", handler.Getdata)
	req := httptest.NewRequest("GET", "/getuser", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{
		"data": [
			{
				"username": "mindriajulian",
				"roles": "user",
				"created_at": null,
				"updated_at": null
			}
		],
		"description": "OK",
		"status": 200
	}`, w.Body.String())
}
