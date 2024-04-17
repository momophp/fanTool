package data

import "time"

type IDNumberInfo struct {
	Birthday       time.Time `json:"birthday"`
	BirthdayString string    `json:"birthday_string"`
	Sex            int       `json:"sex"`
	SexName        string    `json:"sex_name"`
}
