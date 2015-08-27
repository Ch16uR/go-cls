package models

import (
	_ "../utils/uuid"
)

type Workplace struct {
	ID          int    `gorm:"primary_key"`
	Name        string `sql:"size:300"`
	Uuid        string
	FrontolCode int
}
