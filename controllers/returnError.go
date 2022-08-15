package controllers

import (

	"github.com/gin-gonic/gin"
)

func ReturnError(ctx *gin.Context, code int, err error){
	ctx.AbortWithStatusJSON(code, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
}