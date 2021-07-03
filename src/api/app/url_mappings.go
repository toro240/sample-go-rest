package app

import (
	cUsers "mj-app/controllers/users"
)

func mapUrls() {
	router.GET("/user/:user_id", cUsers.GetUser)
	router.POST("/user", cUsers.CreateUser)
	router.PATCH("/user/:user_id", cUsers.UpdateUser)
	router.DELETE("/user/:user_id", cUsers.DeleteUser)
}
