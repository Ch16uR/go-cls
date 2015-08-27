package models

import (
	"github.com/jinzhu/gorm"
)

type Workplace struct {
	gorm.Model
	Name        string `sql:"type:varchar(255)";"not null";`
	Uuid        string `sql:"type=uuid;index:idx_uuid";"not null;unique"`
	FrontolCode int    `sql:"index:idx_frontol_code";"not null"`
}
