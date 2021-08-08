package app

import (
	"github.com/Chestermozhao/GolangFromZero/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}