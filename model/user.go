package model

import "time"

type User struct {
	Model
	FirstNameTH string    `json:"first_name_th" gorm:"type:varchar(255);not null"`
	LastNameTH  string    `json:"last_name_th" gorm:"type:varchar(255);not null"`
	BirthDate   time.Time `json:"birth_date" gorm:"type:date;not null"`
	Address     string    `json:"address" gorm:"type:text;not null"`
}

func (User) TableName() string {
	return "user"
}
