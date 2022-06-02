package controllers

import (
	"fmt"
	"net/http"
	"time"

	// "github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type meetingToParseDate struct {
	Id             int
	Meeting_number uint16
	Date           string
	Description    string
	Video_id       string
	Status         bool
	Patient_id     uint
}

func AddMeeting(ctx *gin.Context) {
	var tempMeeting meetingToParseDate
	// var newMeeting models.Meeting
	if err := ctx.ShouldBindBodyWith(&tempMeeting, binding.JSON); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed1",
			"error":  err.Error(),
		})
		return
	}
	tempDate, err := time.Parse("07-20-2018", tempMeeting.Date)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "failed2",
			"error":  err.Error(),
		})
		return
	}
	fmt.Println(tempDate)
	// if _, err := Db.Model(&newMeeting).Insert(); err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status": "failed2",
	// 		"error":  err.Error(),
	// 	})
	// 	return
	// }
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   tempMeeting,
	})
}
