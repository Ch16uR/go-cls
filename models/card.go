package models

import (
	"github.com/jinzhu/gorm"
)

type Card struct {
	gorm.Model
	DocumentID int `sql:index`
	Code       int
	GroupCode  int
	GroupName  string
	GroupText  string
	Value      string
}
