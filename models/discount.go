package models

import (
	"github.com/jinzhu/gorm"
)

type Discount struct {
	gorm.Model
	DocumentID int `sql:index`
	Code       int
	Kind       int
	Summ       float64
	Text       string
	Type       int
	Value      float64
}
