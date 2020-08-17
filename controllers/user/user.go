package user

import (
	_ "encoding/json"
	_ "fmt"
	_ "io/ioutil"
	"net/http"
	"strconv"

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
	res, err := services.UserService.CreateUser(s)

	if res != nil {
		c.JSON(http.StatusCreated, res.Marshall(c.GetHeader("X-Public") == "true"))
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
	res, err := services.UserService.GetUser(userId)
	if res != nil {
		c.JSON(http.StatusOK, res.Marshall(c.GetHeader("X-Public") == "true"))
	}

	if err != nil {
		c.String(http.StatusNotFound, "Not found")
	}

}

func GetAll(c *gin.Context) {
	res, err := services.UserService.GetAllUsers()
	if err != nil {
		c.String(http.StatusNotFound, "Not found")
	}

	if res != nil {

		result := make([]interface{}, len(res))

		for index, user := range res {
			result[index] = user.Marshall(c.GetHeader("X-Public") == "true")
		}
		c.JSON(http.StatusOK, result)
	}
}

func EditUser(c *gin.Context) {

	var s users.User
	id := c.Param("user_id")
	userId, parseErr := strconv.ParseInt(id, 10, 64)
	if err := c.ShouldBindJSON(&s); err != nil {
		c.String(http.StatusNonAuthoritativeInfo, "Chicharito")
	}
	if parseErr != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "wrong id"})
	}
	s.Id = userId

	res, err := services.UserService.EditUser(&s)

	if err != nil {
	}

	if res != nil {
		c.JSON(http.StatusOK, res)
	} else {
		c.String(http.StatusNotModified, "Cant modified")
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("user_id")

	userId, parseErr := strconv.ParseInt(id, 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusConflict, gin.H{"message": "wrong id"})
	}
	res, err := services.UserService.DeleteUser(userId)
	if err != nil {
		c.String(http.StatusExpectationFailed, "Error")
	}
	if res {
		c.JSON(http.StatusGone, gin.H{"message": "deleted"})
	} else {
		c.JSON(http.StatusGone, gin.H{"message": "Cant deleted"})
	}
}
