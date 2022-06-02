package routes

import (
	"github.com/NoamBoni/gofoloapp/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/login", controllers.Login)
	router.Use(controllers.AuthMiddleware)
	auth := router.Group("/auth")
	auth.Use(controllers.Encrypt, controllers.RestrictToTherapists)
	{
		auth.POST("/register_therapist", controllers.RegisterTherapist)
		auth.POST("/register_patient", controllers.RegisterPatient)
	}
	data := router.Group("/data")
	data.GET("/patients", controllers.RestrictToTherapists, controllers.GetPatients)
	data.POST("/meetings", controllers.RestrictToTherapists, controllers.AddMeeting)
	return router
}
