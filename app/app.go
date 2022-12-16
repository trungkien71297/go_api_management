package app

import (
	"go_api_management/logger"

	"github.com/gin-gonic/gin"
)

var (
	router  = gin.Default()
	public  = router.Group("/api")
	private = router.Group("/api")
)

func StartApplication() {
	mapUrls()
	logger.Info("Nani")
	router.Run(":8080")
}
