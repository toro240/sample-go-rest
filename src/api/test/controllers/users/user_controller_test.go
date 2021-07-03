package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	cUsers "mj-app/controllers/users"
	exceptions "mj-app/exceptions"
	mUsers "mj-app/model/users"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func getRequestHandler(id string) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)

	param := gin.Param{Key: "user_id", Value: id}
	context.Params = gin.Params{param}

	context.Request, _ = http.NewRequest(
		http.MethodGet,
		"/user/:user_id",
		nil,
	)

	return context, response
}

// 正常系
func TestGetNoError(t *testing.T) {
	// Arrange
	u := mUsers.User{Model: gorm.Model{ID: 1}, Name: "coca cola"}
	context, _ := requestHandler(u)
	cUsers.CreateUser(context)

	context2, response := getRequestHandler("1")

	// Act ---
	cUsers.GetUser(context2)

	// Assert ---
	var user mUsers.User
	error := json.Unmarshal(response.Body.Bytes(), &user)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.Nil(t, error)
	assert.EqualValues(t, uint64(1), user.ID)
}

// 不正なIDのテスト
func TestGetWithInvalidID(t *testing.T) {
	// Arrange
	context, response := getRequestHandler("a")

	// Act ---
	cUsers.GetUser(context)

	// Assert ---
	var apiError exceptions.ApiError
	json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, apiError.Message, "user id should be a number")
	assert.EqualValues(t, apiError.Status, 400)
	assert.EqualValues(t, apiError.Error, "bad_request")
}

// 指定したIDのプロダクトが存在しないテスト
func TestGetWithNo(t *testing.T) {
	// Arrange ---
	context, response := getRequestHandler("10000")

	// Act ---
	cUsers.GetUser(context)

	// Assert ---
	var apiError exceptions.ApiError
	json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.EqualValues(t, http.StatusNotFound, response.Code)
	assert.NotNil(t, apiError)
	assert.EqualValues(t, apiError.Message, "not found user id: 10000")
	assert.EqualValues(t, apiError.Status, 404)
	assert.EqualValues(t, apiError.Error, "not_found")
}

func requestHandler(value interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	buff, _ := json.Marshal(value)
	context, _ := gin.CreateTestContext(response)
	context.Request, _ = http.NewRequest(
		http.MethodPost,
		"/user",
		bytes.NewBuffer(buff),
	)
	return context, response
}

func TestCreateUserNoError(t *testing.T) {
	// Arrange ---
	requestValue := mUsers.User{Model: gorm.Model{ID: 123}, Name: "coca cola"}
	context, response := requestHandler(requestValue)

	// Act ---
	cUsers.CreateUser(context)

	// Assert ---
	var user mUsers.User
	error := json.Unmarshal(response.Body.Bytes(), &user)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	assert.Nil(t, error)
	fmt.Println(user)
	assert.EqualValues(t, uint64(123), user.ID)
	assert.EqualValues(t, "coca cola", user.Name)
}

func TestCreateUserWith404Error(t *testing.T) {
	type demiUser struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	// Arrange ---
	requestValue := demiUser{ID: "123", Name: "coca cola"}
	context, response := requestHandler(requestValue)

	// Act ---
	cUsers.CreateUser(context)

	// Assert ---
	var apiError exceptions.ApiError
	error := json.Unmarshal(response.Body.Bytes(), &apiError)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.Nil(t, error)
	assert.EqualValues(t, "invalid json body", apiError.Message)
	assert.EqualValues(t, 400, apiError.Status)
	assert.EqualValues(t, "bad_request", apiError.Error)
}
