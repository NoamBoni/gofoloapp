package models

import (
	"time"
)

type Patient struct {
	Id                  uint      `json:"id" pg:",pk"`
	Therapist_id        *User      `json:"therapist_id" pg:"rel:has-one, join_fk:user_id" binding:"required"`
	User_id             *User     `json:"user_id" binding:"required" pg:"join_fk:user_id"`
	Age                 uint      `json:"age"`
	Gender              string    `json:"gender" pg:"type:varchar(50)"`
	Image_icon          string    `json:"image_icon" pg:"type:varchar(50)"`
	Name_contact_person string    `json:"name_contact_person" pg:"type:varchar(50)"`
	Contact_person      string    `json:"contact_person" pg:"type:varchar(50)"`
	Total_meetings      uint      `json:"total_meetings"`
	Meeting_number      uint      `json:"meeting_number"`
	Next_meeting        time.Time `json:"next_meeting"`
	Meetings            *Meeting  `json:"meetings" pg:"rel:has-many,join_fk:meeting_id"`
}

// func (p *Patient) String() string {
// 	return fmt.Sprintf("<id:%d name:%s therapist: %s role:%s>", p.Id)
// }
