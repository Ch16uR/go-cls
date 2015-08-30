package controllers

import (
	_ "fmt"
	_ "github.com/martini-contrib/render"
	"time"
)

type Calc struct {
	Token    string    `json: token`
	OpenedAt time.Time `json: openedAt`
	Value    string    `json: value`
	Receipt  []receipt
}
type receipt struct {
	ReceiptNo       string
	ReceiptTypeCode int
	DateOpen        time.Time
	SessionNo       int
	SummForD        float32
	SummWD          float32
}
