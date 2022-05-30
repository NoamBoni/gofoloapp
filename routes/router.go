package routes

import (
	"github.com/NoamBoni/gofoloapp/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	auth := router.Group("/auth")
	auth.Use(controllers.Encrypt)
	{
		auth.POST("/register_therapist", controllers.RegisterTherapist)
		auth.POST("/register_patient", controllers.RegisterPatient)
	}
	router.POST("/login", controllers.Login)
	return router
}
