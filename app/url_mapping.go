package app

import (
	"go_api_management/controllers/country"
	"go_api_management/controllers/ping"
	"go_api_management/controllers/user"
)

func mapUrls() {

	public.GET("/ping", ping.Ping)
	//auth
	public.POST("/login", nil)
	public.POST("/register", nil)

	//user
	private.GET("/user/:user_id", user.FindUser)
	private.POST("/user", user.CreateUser)
	private.PUT("/user/:user_id", user.EditUser)
	private.DELETE("/user/:user_id", user.DeleteUser)
	private.GET("/users", user.GetAll)

	//country
	public.GET("/countries/", country.GetCountries)
	public.GET("/country/:country_id", country.GetCountry)
	private.POST("/country", country.AddCountry)
	private.PUT("/country/:country_id", country.EditCountry)
	private.DELETE("/country/:country_id", country.DeleteCountry)
}
