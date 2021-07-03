package user

import (
	exceptions "mj-app/exceptions"
	"strings"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

// New
// Validate - check parameters user inputs
func (u *User) Validate() *exceptions.ApiError {
	u.Name = strings.TrimSpace(strings.ToLower(u.Name))
	if u.Name == "" {
		return exceptions.NewBadRequestError("invalid user name")
	}
	return nil
}
