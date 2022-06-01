package models

import "fmt"

type User struct {
	Id       uint       `json:"id" pg:",pk"`
	Name     string     `json:"name" pg:"type:varchar(50)" binding:"required"`
	Password string     `json:"password" pg:"type:varchar(150)" binding:"required"`
	Role     string     `json:"role" pg:"type:varchar(50)"`
	Patients []*Patient `json:"patients" sql:"-" pg:"rel:has-many, fk:patient_id"`
}

func (u *User) String() string {
	return fmt.Sprintf("%+v\n", *u)
}
