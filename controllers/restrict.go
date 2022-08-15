package controllers

import (
	"net/http"
	"errors"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

func RestrictToTherapists(ctx *gin.Context){
	role, got := ctx.Get("role")
	if !got{
		ReturnError(ctx, http.StatusInternalServerError, errors.New("something's wrong, try again later"))
		return
	}
	if role != models.Role.T{
		ReturnError(ctx, http.StatusUnauthorized, errors.New("not allowed"))
		return
	}
	ctx.Next()
}