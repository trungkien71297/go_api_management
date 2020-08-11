package app

import (
	"github.com/trungkien71297/go_api_management/controllers/ping"
	"github.com/trungkien71297/go_api_management/controllers/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", user.GetUsers)
	//	router.GET("/users/search/", controllers.FindUser)
	router.POST("/users", user.CreateUser)
}
