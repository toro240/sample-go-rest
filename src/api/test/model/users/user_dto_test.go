package users

import (
	"testing"

	mUsers "mj-app/model/users"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestProductValiateNoError(t *testing.T) {
	// Arrange ---
	u := mUsers.User{Model: gorm.Model{ID: 123}, Name: "coca cola"}

	// Act ---
	error := u.Validate()

	// Assert ---
	assert.Nil(t, error)
}

func TestProductValiateBadRequestError(t *testing.T) {
	// Arrange ---
	u := mUsers.User{Model: gorm.Model{ID: 123}}

	// Act ---
	error := u.Validate()

	// Assert ---
	assert.NotNil(t, error)
	assert.EqualValues(t, "invalid user name", error.Message)
	assert.EqualValues(t, 400, error.Status)
	assert.EqualValues(t, "bad_request", error.Error)

}
