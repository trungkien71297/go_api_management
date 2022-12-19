package country_controller

import (
	"go_api_management/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCountry(c *gin.Context) {

}

func GetCountries(c *gin.Context) {
	res, err := services.CountryServices.GetAllCountries()
	if res != nil {
		c.JSON(http.StatusOK, res)
	}

	if err != nil {
		c.String(http.StatusInternalServerError, "Server error")
	}

}

func EditCountry(c *gin.Context) {

}

func DeleteCountry(c *gin.Context) {

}

func AddCountry(c *gin.Context) {

}
