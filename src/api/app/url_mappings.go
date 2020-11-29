package app

import (
	"github.com/ag3ntsc4rn/golang-microservices/src/api/controllers/repositories"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/controllers/polo"
)

func mapUrls(){
	router.GET("/marco", polo.Polo)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}