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

//safe decides if the password should be hidden
func (u *User) Insert(safe bool) error {
	if _, err := Db.Model(u).Insert(); err != nil {
		return err
	}
	if safe {
		u.Password = ""
	}
	return nil
}

func (u *User) Delete() error {
	_, err := Db.Model(u).Where("id = ?", u.Id).Delete()
	return err
}

//safe decides if the password should be hidden
func (u *User) Select(safe bool) error {
	if err := Db.Model(&u).Select(); err != nil {
		return err
	}
	if safe {
		u.Password = ""
	}
	return nil
}

func (u *User) GetPatients() error {
	query := fmt.Sprintf("select * from users join patients on users.id = patients.user_id where patients.therapist_id = %v", u.Id)
	fmt.Println(query)
	if _, err := Db.Query(&u.Patients, query); err != nil {
		return err
	}
	for _, val := range u.Patients {
		val.Password = ""
	}
	return nil
}

func GetUsersByName(name string, users *[]User) error {
	return Db.Model(users).Where("name = ?", name).Select()
}
