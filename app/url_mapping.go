package app

import (
	"github.com/trungkien71297/go_api_management/controllers"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
