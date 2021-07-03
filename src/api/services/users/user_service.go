package users

import (
	"mj-app/exceptions"
	mUsers "mj-app/model/users"

	"github.com/jinzhu/gorm"
)

func GetUser(userID uint) (*mUsers.User, *exceptions.ApiError) {
	u := &mUsers.User{Model: gorm.Model{ID: userID}}
	if error := u.Get(); error != nil {
		return nil, error
	}
	return u, nil
}

func CreateUser(user mUsers.User) (*mUsers.User, *exceptions.ApiError) {
	if error := user.Validate(); error != nil {
		return nil, error
	}

	if error := user.Save(); error != nil {
		return nil, error
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user mUsers.User) (*mUsers.User, *exceptions.ApiError) {
	current, error := GetUser(user.ID)
	if error != nil {
		return nil, error
	}

	if isPartial {
		if user.Name != "" {
			current.Name = user.Name
		}
		if err := current.PartialUpdate(); err != nil {
			return nil, err
		}
	} else {
		current.Name = user.Name
	}

	if error := current.Update(); error != nil {
		return nil, error
	}

	return current, nil
}

func DeleteUser(userID uint) *exceptions.ApiError {
	user := &mUsers.User{Model: gorm.Model{ID: userID}}
	return user.Delete()
}
