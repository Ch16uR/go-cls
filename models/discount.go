package models

import (
	"github.com/jinzhu/gorm"
)

type Discount struct {
	gorm.Model
	Value int `sql:"not null"`
}
