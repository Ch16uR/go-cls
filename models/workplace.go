package models

type Workplace struct {
	ID          int    `gorm:"primary_key"`
	Name        string `sql:"size:300"`
	Uuid        string `sql:"DEFAULT:uuid_generate_v4()"`
	FrontolCode int
}
