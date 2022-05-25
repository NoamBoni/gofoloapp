package models

import (
	"time"
)

type Meeting struct {
	Id             int       `json:"id" pg:",pk"`
	Meeting_number uint16    `json:"meeting_number"`
	Date           time.Time `json:"date"`
	Description    string    `json:"description" pg:"type:varchar(400)"`
	Video_id       string    `json:"video_id" pg:"type:varchar(50)"`
	Status         bool      `json:"status"`
	Patient        *Patient  `json:"patient" pg:"rel:has-one"`
}

// func (m *Meeting) String() string {
// 	return fmt.Sprintf("<id:%d patient:%s therapist: %s status:%v>", m.Id, m.Patient.Name, m.Patient.Therapist.Name, m.Status)
// }
