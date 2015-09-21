package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Promocode struct {
	gorm.Model
	Name                    string `sql:"type:varchar(255)";"not null"`
	Value                   string `sql:"index:idx_value";"not null"`
	ExpiredAt               time.Time
	OnlyOnce                bool
	Discount                float32
	PromocodeDiscountTypeID int `sql:"index"`
	CompanyID               int `sql:"index"`
}
