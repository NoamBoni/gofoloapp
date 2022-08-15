package controllers

import (
	"net/http"
	"strconv"

	"github.com/NoamBoni/gofoloapp/models"
	"github.com/gin-gonic/gin"
)

func GetMeetings(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ReturnError(ctx,http.StatusInternalServerError,err)
		return
	}
	meetings, err := models.GetAllMeetingsByPatientId(uint(id))
	if err != nil {
		ReturnError(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"length": len(meetings),
		"data":   meetings,
	})
}
