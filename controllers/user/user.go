package user

import (
	_ "encoding/json"
	_ "fmt"
	_ "io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trungkien71297/go_api_management/domain/users"
	"github.com/trungkien71297/go_api_management/services"
)

func GetUsers(c *gin.Context) {
	c.String(http.StatusNotImplemented, "This function is not implemented")
}
func CreateUser(c *gin.Context) {
	//c.String(http.StatusNotImplemented, "This function is not implemented")
	var s users.User
	if err := c.ShouldBindJSON(&s); err != nil {
		c.String(http.StatusNonAuthoritativeInfo, "Chicharito")
	}
	res, err := services.CreateUser(s)

	if res != nil {
		c.JSON(http.StatusCreated, gin.H{
			"status":   "OK",
			"username": res.Username,
		})
		return
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save user")
		return
	}
}
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "This function is not implemented")
}
