package models

import (
	"github.com/jinzhu/gorm"
)

type Promocode struct {
	gorm.Model
	Name           string `sql:"type:varchar(255)";"not null"`
	Value          string `sql:"index:idx_value";"not null"`
	Discount       Discount
	DiscountID     int `sql:"index"`
	DiscountType   DiscountType
	DiscountTypeID int `sql:"index"`
}
