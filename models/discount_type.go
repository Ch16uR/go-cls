package models

import (
	"github.com/jinzhu/gorm"
)

type DiscountType struct {
	gorm.Model
	Name string `sql:"type:varchar(255)";"not null"`
	Type int
}
