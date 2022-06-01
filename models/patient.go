package models

import (
	"time"
	"fmt"
)

type Patient struct {
	Id                  uint       `json:"id" pg:",pk"`
	Therapist_id        uint       `json:"therapist_id" pg:"rel:has-one, fk:user_id, on_delete:CASCADE" binding:"required"`
	User_id             uint       `json:"user_id" pg:"rel:has-one, fk:user_id, on_delete:CASCADE"`
	Age                 uint       `json:"age"`
	Gender              string     `json:"gender" pg:"type:varchar(50)"`
	Image_icon          string     `json:"image_icon" pg:"type:varchar(50)"`
	Name_contact_person string     `json:"name_contact_person" pg:"type:varchar(50)"`
	Contact_person      string     `json:"contact_person" pg:"type:varchar(50)"`
	Total_meetings      uint       `json:"total_meetings"`
	Meeting_number      uint       `json:"meeting_number"`
	Next_meeting        time.Time  `json:"next_meeting"`
	Meetings            []*Meeting `json:"meetings" sql:"-" pg:"rel:has-many, fk:meeting_id"`
	Name                string     `json:"name" sql:"-"`
	Password            string     `json:"password" sql:"-"`
	Role                string     `json:"role" sql:"-"`
}

func (p *Patient) String() string {
	return fmt.Sprintf("%+v\n", *p)
}
