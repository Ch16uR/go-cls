package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Document struct {
	gorm.Model
	DateOpened      time.Time
	DateClosed      time.Time
	WorkplaceID     int `sql:"index"`
	CompanyID       int `sql:"index"`
	Education       int
	OpenSellerCode  int `sql:"index"`
	OpenSellerName  string
	CloseSellerCode int `sql:"index"`
	CloseSellerName string
	ReceiptNo       string
	ReceiptTypeCode int
	SessionNo       int
	SummForD        float64
	SummWD          float64
	TotalDiscSumm   float64
	TotalDiscValue  float64
	Cards           []Card
	Discounts       []Discount
}
