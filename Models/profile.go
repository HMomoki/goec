package models

import(
	"github.com/jinzhu/gorm"
)

type Profile struct{
	gorm.Model
	First_name string `gorm:"not null;type:VARCHAR(255)"`
	Last_name string `gorm:"not null;type:VARCHAR(255)"`
	Phonenumber string `gorm:"not null;type:VARCHAR(11)"`
	sex string
}