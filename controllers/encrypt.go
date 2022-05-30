package controllers

import (
	"net/http"
	"strings"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	if strings.Contains(newUser.Password, " ") || strings.Contains(newUser.Name, " ") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid name or password",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	ctx.Set("crypted-password", string(hash))
	ctx.Next()
}
