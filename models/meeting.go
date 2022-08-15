package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Meeting struct {
	Id             int       `json:"id" pg:",pk"`
	Meeting_number uint16    `json:"meeting_number" binding:"required"`
	Date           time.Time `json:"date" binding:"required"`
	Description    string    `json:"description" pg:"type:varchar(400)" binding:"required"`
	Video_id       string    `json:"video_id" pg:"type:varchar(50)" binding:"required"`
	Status         bool      `json:"status" pg:"default:false"`
	Patient_id     uint      `json:"patient_id" pg:"required,rel:has-one, fk:patient_id, on_delete:CASCADE,notnull" binding:"required"`
}

type MeetingToParseDate struct {
	Meeting_number uint16 `binding:"required"`
	Date           string `binding:"required"`
	Description    string `binding:"required"`
	Video_id       string `binding:"required"`
	Status         bool
	Patient_id     uint `binding:"required"`
}

func (m *Meeting) String() string {
	return fmt.Sprintf("%+v\n", *m)
}

func (m *Meeting) Insert(mtp *MeetingToParseDate) error {
	date, err := parseDate(mtp.Date)
	if err != nil {
		return err
	}
	m.Date = date
	m.Meeting_number = mtp.Meeting_number
	m.Description = mtp.Description
	m.Video_id = mtp.Video_id
	m.Status = mtp.Status
	m.Patient_id = mtp.Patient_id

	if _, err := Db.Model(m).Insert(); err != nil {
		return err
	}

	return nil
}

func GetAllMeetingsByPatientId(id uint) ([]Meeting, error) {
	var meetings []Meeting
	err := Db.Model(&meetings).Where("patient_id = ?", id).Select()
	return meetings, err
}

func parseDate(date string) (time.Time, error) {
	err := errors.New("invalid input please insert in a format of dd-mm-yyyy")
	dates := strings.Split(date, "-")
	if len(dates) != 3 {
		return time.Now(), err
	}
	year, _ := strconv.Atoi(dates[2])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[0])
	if day > 31 || month > 12 {
		return time.Now(), err
	}
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}
