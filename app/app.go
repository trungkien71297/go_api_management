package app

import (
	"github.com/gin-gonic/gin"
	"github.com/trungkien71297/go_api_management/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Nani")
	router.Run(":8080")
}
