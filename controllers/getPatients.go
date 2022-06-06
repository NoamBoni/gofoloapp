package controllers

import (
	"net/http"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

func GetPatients(ctx *gin.Context) {
	id, got := ctx.Get("user-id")
	if !got {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  "something's wrong, try again later",
		})
		return
	}
	user := models.User{
		Id: id.(uint),
	}
	_ = user.Select(true)
	if err := user.GetPatients(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user.Patients,
		"length": len(user.Patients),
	})
}
