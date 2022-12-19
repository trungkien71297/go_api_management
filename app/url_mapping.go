package app

import (
	auth_controller "go_api_management/controllers/auth"
	country_controller "go_api_management/controllers/country"
	"go_api_management/controllers/ping"
	user_controller "go_api_management/controllers/user"
)

func mapUrls() {

	public.GET("/ping", ping.Ping)
	//auth
	public.POST("/login", auth_controller.Login)
	public.POST("/register", nil)

	//user
	private.GET("/user/:user_id", user_controller.FindUser)
	private.POST("/user", user_controller.CreateUser)
	private.PUT("/user/:user_id", user_controller.EditUser)
	private.DELETE("/user/:user_id", user_controller.DeleteUser)
	private.GET("/users", user_controller.GetAll)

	//country
	public.GET("/countries/", country_controller.GetCountry)
	public.GET("/country/:country_id", country_controller.GetCountry)
	private.POST("/country", country_controller.AddCountry)
	private.PUT("/country/:country_id", country_controller.EditCountry)
	private.DELETE("/country/:country_id", country_controller.DeleteCountry)
}
