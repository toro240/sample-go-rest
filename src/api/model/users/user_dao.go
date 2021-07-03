package user

import (
	"mj-app/datasources/mysql"
	"mj-app/exceptions"
)

func (u *User) Get() *exceptions.ApiError {
	if result := mysql.Client.Where("id = ?", u.Model.ID).Find(&u); result.Error != nil {
		return mysql.ParseError(result)
	}
	return nil
}

func (u *User) Save() *exceptions.ApiError {
	if result := mysql.Client.Create(&u); result.Error != nil {
		return mysql.ParseError(result)
	}
	return nil
}

func (u *User) Update() *exceptions.ApiError {
	if result := mysql.Client.Save(&u); result.Error != nil {
		return mysql.ParseError(result)
	}
	return nil
}

func (u *User) PartialUpdate() *exceptions.ApiError {
	if result := mysql.Client.
		Table("users").
		Where("id IN (?)", u.ID).
		Updates(&u); result.Error != nil {
		return mysql.ParseError(result)
	}
	return nil
}

func (u *User) Delete() *exceptions.ApiError {
	if result := mysql.Client.Delete(&u); result.Error != nil {
		return mysql.ParseError(result)
	}
	return nil
}
