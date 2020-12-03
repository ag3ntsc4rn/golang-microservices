package app

import (
	logoptiona "github.com/ag3ntsc4rn/golang-microservices/src/api/log/option_a"
	logoptionb "github.com/ag3ntsc4rn/golang-microservices/src/api/log/option_b"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {
	logoptiona.Info("About to map urls.", "step:1", "status:pending")
	logoptionb.Info("this is from zap", logoptionb.Field("custom", "customvalue"))
	mapUrls()
	logoptiona.Info("Urls mapped successfully.")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
