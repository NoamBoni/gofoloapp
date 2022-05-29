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

	name := strings.ReplaceAll(newUser.Name, " ", "")
	password := strings.ReplaceAll(newUser.Password, " ", "")

	if password == "" || name == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid name or password",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	password = string(hash)

	ctx.Set("valid-name", name)
	ctx.Set("crypted-password", password)
	ctx.Next()
}
