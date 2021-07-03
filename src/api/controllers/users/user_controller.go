package users

import (
	"net/http"
	"strconv"

	"mj-app/exceptions"
	mUsers "mj-app/model/users"
	sUsers "mj-app/services/users"

	"github.com/gin-gonic/gin"
)

func getUserID(userIDParam string) (uint, *exceptions.ApiError) {
	userId, error := strconv.ParseUint(userIDParam, 10, 64)
	if error != nil {
		return 0, exceptions.NewBadRequestError("user id should be a number")
	}
	return uint(userId), nil
}

func GetUser(c *gin.Context) {
	userID, idError := getUserID(c.Param("user_id"))
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}

	user, error := sUsers.GetUser(uint(userID))
	if error != nil {
		c.JSON(error.Status, error)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user mUsers.User
	if error := c.ShouldBindJSON(&user); error != nil {
		apiError := exceptions.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}
	newUser, error := sUsers.CreateUser(user)
	if error != nil {
		c.JSON(error.Status, error)
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c *gin.Context) {
	userID, idError := getUserID(c.Param("user_id"))
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}
	var user mUsers.User
	if error := c.ShouldBindJSON(&user); error != nil {
		apiError := exceptions.NewBadRequestError("invalid json body")
		c.JSON(apiError.Status, apiError)
		return
	}

	user.ID = uint(userID)

	isPartial := c.Request.Method == http.MethodPatch

	newUser, error := sUsers.UpdateUser(isPartial, user)
	if error != nil {
		c.JSON(error.Status, error)
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func DeleteUser(c *gin.Context) {
	userID, idError := getUserID(c.Param("user_id"))
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}
	if err := sUsers.DeleteUser(userID); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
