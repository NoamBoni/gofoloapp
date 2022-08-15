package controllers

import (
	"net/http"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

func AddMeeting(ctx *gin.Context) {
	var tempMeeting models.MeetingToParseDate
	if err := ctx.ShouldBindJSON(&tempMeeting); err != nil {
		ReturnError(ctx, http.StatusBadRequest, err)
		return
	}

	var newMeeting models.Meeting
	if err := newMeeting.Insert(&tempMeeting); err != nil {
		ReturnError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   newMeeting,
	})
}
