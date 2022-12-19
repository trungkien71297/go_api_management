package auth_controller

import (
	"go_api_management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var params = make(map[string]string)
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad params"})
		return
	}
	username, passwrod := params["username"], params["password"]
	if username == "" || passwrod == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad params"})
		return
	}
	res, err := services.AuthService.Login(username, passwrod)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	if res != nil {
		c.JSON(http.StatusOK, res)
	}
}
