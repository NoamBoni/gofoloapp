package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterTherapist(ctx *gin.Context) {
	var newTherapist models.User
	if err := ctx.ShouldBindBodyWith(&newTherapist, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	newTherapist.Role = models.Role.T
	setEncryptedPassword(ctx, &newTherapist)
	if err := newTherapist.Insert(true); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   newTherapist,
	})

}

func RegisterPatient(ctx *gin.Context) {
	var newPatient models.Patient
	if err := ctx.ShouldBindBodyWith(&newPatient, binding.JSON); err != nil {
		ReturnError(ctx, http.StatusBadRequest, err)
		return
	}

	newUser := models.User{
		Name: newPatient.Name,
		Role: models.Role.P,
	}
	setEncryptedPassword(ctx, &newUser)

	if err := newUser.Insert(true); err != nil {
		ReturnError(ctx, http.StatusBadRequest, err)
		return
	}

	newPatient.User_id = newUser.Id

	if err := newPatient.Insert(); err != nil {
		_ = newUser.Delete()
		ReturnError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   newPatient,
	})
}

func setEncryptedPassword(ctx *gin.Context, user *models.User) {
	password, got := ctx.Get("crypted-password")
	if !got {
		ReturnError(ctx, http.StatusBadRequest, errors.New("something's wrong, try again later"))
		return
	}
	user.Password = fmt.Sprintf("%v", password)
}
