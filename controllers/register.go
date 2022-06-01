package controllers

import (
	"fmt"
	"net/http"

	"github.com/NoamBoni/gofoloapp/helpers"
	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-pg/pg"
)

var Db *pg.DB

func init(){
	Db = helpers.ConnectDB()
}

func RegisterTherapist(ctx *gin.Context) {
	var newTherapist models.User
	if err := ctx.ShouldBindBodyWith(&newTherapist, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	newTherapist.Role = "Therapist"
	updateUser(ctx, &newTherapist)
	if _, err := Db.Model(&newTherapist).Insert(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	newTherapist.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   newTherapist,
	})

}

func RegisterPatient(ctx *gin.Context) {
	var newPatient models.Patient
	if err := ctx.ShouldBindBodyWith(&newPatient, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	newUser := models.User{
		Name:     newPatient.Name,
		Password: newPatient.Password,
		Role:     "Patient",
	}
	updateUser(ctx, &newUser)

	if _, err := Db.Model(&newUser).Insert(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	newPatient.User_id = newUser.Id

	if _, err := Db.Model(&newPatient).Insert(); err != nil {
		_, _ = Db.Model(&newUser).Where("id = ?", newUser.Id).Delete()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	newPatient.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   newPatient,
	})
}

func updateUser(ctx *gin.Context, user *models.User) {
	password, got := ctx.Get("crypted-password")
	if !got {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  "something's wrong, try again later",
		})
		return
	}
	user.Password = fmt.Sprintf("%v", password)
}
