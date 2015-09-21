package models

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name       string `sql:"type:varchar(255)";"not null";`
	Workplaces []Workplace
}
