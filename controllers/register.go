package controllers

import (
	"fmt"
	"net/http"

	"github.com/NoamBoni/gofoloapp/helpers"
	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

var Db = helpers.ConnectDB()

func RegisterTherapist(ctx *gin.Context) {
	var newTherapist models.User
	if err := ctx.Bind(&newTherapist); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	newTherapist.Role = "Therapist"
	_, err := Db.Model(&newTherapist).Insert()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
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
	if err := ctx.Bind(&newPatient); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newUser := models.User{
		Name:     newPatient.Name,
		Password: newPatient.Password,
		Role:     "Patient",
	}

	_, err := Db.Model(&newUser).Insert()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newPatient.User_id = newUser.Id

	if _, err := Db.Model(&newPatient).Insert(); err != nil {
		if _, e := Db.Model(&newUser).Where("id = ?", newUser.Id).Delete(); e != nil {
			fmt.Println(e.Error())
		}
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newPatient.Password = ""
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"data":   newPatient,
	})
}
