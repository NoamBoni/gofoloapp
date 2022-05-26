package routes

import (
	"github.com/NoamBoni/gofoloapp/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register_therapist", controllers.RegisterTherapist)
	router.POST("/register_patient", controllers.RegisterPatient)
	return router
}
