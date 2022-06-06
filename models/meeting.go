package models

import (
	"fmt"
	"time"
)

type Meeting struct {
	Id             int       `json:"id" pg:",pk"`
	Meeting_number uint16    `json:"meeting_number" binding:"required"`
	Date           time.Time `json:"date" binding:"required"`
	Description    string    `json:"description" pg:"type:varchar(400)" binding:"required"`
	Video_id       string    `json:"video_id" pg:"type:varchar(50)" binding:"required"`
	Status         bool      `json:"status" binding:"required"`
	Patient_id     uint      `json:"patient_id" pg:"rel:has-one, fk:patient_id, on_delete:CASCADE" binding:"required"`
}

func (m *Meeting) String() string {
	return fmt.Sprintf("%+v\n", *m)
}

// func (m *Meeting) Insert(){

// }