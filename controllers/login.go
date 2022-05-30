package controllers

import (
	"fmt"
	"net/http"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var loginUser models.User
	var usersList []models.User
	if err := ctx.ShouldBindBodyWith(&loginUser, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}
	if err := Db.Model(&usersList).Where("name = ?", loginUser.Name).Select(); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  "invalid email or password",
		})
		return
	}
	fmt.Printf("%+v\n", loginUser)
	fmt.Printf("%+v\n", usersList)
	for _, user := range usersList {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err == nil {
			if user.Role == "Therapist" {
				loginUser.Password = ""
				loginUser.Role = user.Role
				ctx.JSON(http.StatusOK, gin.H{
					"status": "successful authentication",
					"data":   loginUser,
				})
				fmt.Println("fmt")
				return
			} else {
				patient := models.Patient{
					User_id: user.Id,
					Role:    user.Role,
					Name:    user.Name,
				}
				e := Db.Model(&patient).Select()
				if e != nil {
					fmt.Println(e)
				}
				patient.Password = ""
				ctx.JSON(http.StatusOK, gin.H{
					"status": "successful authentication",
					"data":   patient,
				})
				return
			}
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"status": "failed",
		"error":  "invalid email or password",
	})
}
