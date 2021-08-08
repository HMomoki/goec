package models

import(
	"github.com/jinzhu/gorm"
)

type Profile struct{
	gorm.Model
	Name string `gorm:"not null;type:VARCHAR(255)"`
}