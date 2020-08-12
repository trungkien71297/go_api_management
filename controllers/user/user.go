package user

import (
	_ "encoding/json"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/trungkien71297/go_api_management/domain/users"
	"github.com/trungkien71297/go_api_management/services"
	_ "io/ioutil"
	"net/http"
	"strconv"
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
	id := c.Param("user_id")
	userId, parseErr := strconv.ParseInt(id, 10, 64)

	if parseErr != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "wrong id"})
	}
	res, err := services.GetUsers(userId)
	if res != nil {
		c.JSON(http.StatusOK, res)
	}

	if err != nil {
		c.String(http.StatusNotFound, "Not found")
	}

}
