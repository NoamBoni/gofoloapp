package controllers

import (
	"fmt"
	"net/http"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

func GetPatients(ctx *gin.Context) {
	id, got := ctx.Get("user-id")
	if !got {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  "something's wrong, try again later",
		})
		return
	}
	var user models.User
	_ = Db.Model(&user).Where("id = ?", id.(uint)).Select()
	query := fmt.Sprintf("select * from users join patients on users.id = patients.user_id where patients.therapist_id = %v", id)
	_, _ = Db.Query(&user.Patients, query)
	for _, val := range user.Patients {
		val.Password = ""
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   user.Patients,
		"length": len(user.Patients),
	})
}
