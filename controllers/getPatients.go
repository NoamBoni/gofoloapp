package controllers

import (
	"errors"
	"net/http"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

func GetPatients(ctx *gin.Context) {
	id, got := ctx.Get("user-id")
	if !got {
		ReturnError(ctx, http.StatusInternalServerError, errors.New("something's wrong, try again later"))
		return
	}
	user := models.User{
		Id: id.(uint),
	}
	_ = user.Select(true)
	if err := user.GetPatients(); err != nil {
		ReturnError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user.Patients,
		"length": len(user.Patients),
	})
}
