package controllers

import (
	"net/http"

	"github.com/NoamBoni/gofoloapp/helpers"
	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

var Db = helpers.ConnectDB()

func RegisterTherapist(ctx *gin.Context) {
	var newTherapist models.User
	if err := ctx.Bind(&newTherapist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	newTherapist.Role = "Therapist"
	result, err := Db.Model(&newTherapist).Insert()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   result,
	})

}
