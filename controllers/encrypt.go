package controllers

import (
	"errors"
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
		ReturnError(ctx, http.StatusBadRequest, errors.New("invalid name or password"))
		return
	}

	if strings.Contains(newUser.Password, " ") || strings.Contains(newUser.Name, " ") {
		ReturnError(ctx, http.StatusBadRequest, errors.New("invalid name or password"))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ReturnError(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.Set("crypted-password", string(hash))
	ctx.Next()
}
