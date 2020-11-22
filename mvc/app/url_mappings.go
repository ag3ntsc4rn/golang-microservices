package app

import (
	"github.com/ag3ntsc4rn/golang-microservices/mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
