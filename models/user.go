package models

type User struct {
	Id       uint       `json:"id" pg:",pk"`
	Name     string     `json:"name" pg:"type:varchar(50)" binding:"required"`
	Password string     `json:"password" pg:"type:varchar(150)" binding:"required"`
	Role     string     `json:"role" pg:"type:varchar(50)"`
	Patients []*Patient `json:"patients" sql:"-"`
}
