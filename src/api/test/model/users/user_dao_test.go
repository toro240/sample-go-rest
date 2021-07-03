package users

import (
	"testing"

	mUsers "mj-app/model/users"

	"github.com/jinzhu/gorm"

	"github.com/stretchr/testify/assert"
)

func TestGetNoError(t *testing.T) {
	// Arrange ---
	u := mUsers.User{
		Model: gorm.Model{ID: 1},
		Name:  "coca cola",
	}
	u.Save()
	newUser := mUsers.User{Model: gorm.Model{ID: 1}}

	// Act ---
	result := newUser.Get()

	// Arrange ---
	assert.Nil(t, result)
	assert.EqualValues(t, u.Name, newUser.Name)
}

func TestNotFound(t *testing.T) {
	// Arrange ---
	p := mUsers.User{Model: gorm.Model{ID: 100}}

	// Act ---
	error := p.Get()

	// Assert ---
	assert.NotNil(t, error)
	assert.EqualValues(t, error.Message, "not found user id: 100")
	assert.EqualValues(t, error.Status, 404)
	assert.EqualValues(t, error.Error, "not_found")
}

func TestSaveNoError(t *testing.T) {
	// Arrange ---
	u := mUsers.User{
		Model: gorm.Model{ID: 1},
		Name:  "coca cola",
	}

	// Act ---
	error := u.Save()

	// Assert ---
	assert.Nil(t, error)
}

// 同じIDを保存した場合エラーになる
func TestSaveBadRequestErrorWithSameID(t *testing.T) {
	// Arrange ---
	u := mUsers.User{
		Model: gorm.Model{ID: 1},
		Name:  "coca cola",
	}

	u.Save()

	u2 := mUsers.User{
		Model: gorm.Model{ID: 1},
		Name:  "orange",
	}

	// Act ---
	error := u2.Save()

	// Assert ---
	assert.NotNil(t, error)
	assert.EqualValues(t, error.Message, "already exists user id: 1")
	assert.EqualValues(t, error.Status, 400)
	assert.EqualValues(t, error.Error, "bad_request")
}
