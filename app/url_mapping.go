package app

import (
	"github.com/trungkien71297/go_api_management/controllers/ping"
	"github.com/trungkien71297/go_api_management/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/user/:user_id", user.FindUser)
	router.POST("/user", user.CreateUser)
	router.PUT("/user/:user_id", user.EditUser)
	router.DELETE("/user/:user_id", user.DeleteUser)
	router.GET("/users", user.GetAll)
}
