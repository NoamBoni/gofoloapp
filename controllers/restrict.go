package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RestrictToTherapists(ctx *gin.Context){
	role, got := ctx.Get("role")
	if !got{
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
			"status": "failed",
			"error":  "something's wrong, try again later",
		})
		return
	}
	if role != "Therapist"{
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
			"status": "failed",
			"error":  "not allowed",
		})
		return
	}
	ctx.Next()
}