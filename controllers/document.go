package controllers

import (
	. "../db"
	. "../models"
	"../utils/jwt"
	"fmt"
	"github.com/martini-contrib/render"
	"time"
)

type InputDocumentAdd struct {
	Token           string  `json:"token"`
	DateOpened      string  `json:"date_opened"`
	DateClosed      string  `json:"date_closed"`
	WorkplaceID     int     `json:"workplace_id"`
	Education       int     `json:"Education"`
	OpenSellerCode  int     `json:"open_seller_code"`
	OpenSellerName  string  `json:"open_seller_name"`
	CloseSellerCode int     `json:"close_seller_code"`
	CloseSellerName string  `json:"close_seller_name"`
	ReceiptNo       string  `json:"receipt_no"`
	ReceiptTypeCode int     `json:"receipt_type_code"`
	SessionNo       int     `json:"session_no"`
	SummForD        float64 `json:"summ_for_d"`
	SummWd          float64 `json:"summ_wd"`
	TotalDiscSumm   float64 `json:"total_disc_summ"`
	TotalDiscValue  float64 `json:"total_disc_value"`
	Cards           []struct {
		Code      int    `json:"code"`
		GroupCode int    `json:"group_code"`
		GroupName string `json:"group_name"`
		GroupText string `json:"group_text"`
		Value     string `json:"value"`
	} `json:"cards"`
	Discounts []struct {
		Code  int     `json:"code"`
		Kind  int     `json:"kind"`
		Summ  float64 `json:"summ"`
		Text  string  `json:"text"`
		Type  int     `json:"type"`
		Value float64 `json:"value"`
	} `json:"discounts"`
}

func DocumentAdd(r render.Render, json InputDocumentAdd) {
	fmt.Printf("%+v\n", json)
	id, company, err := jwt.ParseToken(json.Token)
	if err != nil {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid Token"})
		return
	}

	if json.Education == 1 {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Education mod On"})
		return
	}

	dateOpened, err := time.Parse(time.RFC3339, json.DateOpened)
	if err != nil {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid Datetime"})
		return
	}
	dateClosed, err := time.Parse(time.RFC3339, json.DateClosed)
	if err != nil {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid Datetime"})
		return
	}
	card := Card{}
	cards := []Card{}
	for _, c := range json.Cards {
		card.Code = c.Code
		card.GroupCode = c.GroupCode
		card.GroupName = c.GroupName
		card.GroupText = c.GroupText
		card.Value = c.Value
		cards = append(cards, card)
	}
	discount := Discount{}
	discounts := []Discount{}
	for _, d := range json.Discounts {
		discount.Code = d.Code
		discount.Kind = d.Kind
		discount.Summ = d.Summ
		discount.Text = d.Text
		discount.Type = d.Type
		discount.Value = d.Value

		discounts = append(discounts, discount)
	}

	document := Document{
		DateOpened:      dateOpened,
		DateClosed:      dateClosed,
		WorkplaceID:     id,
		CompanyID:       company,
		Education:       json.Education,
		OpenSellerCode:  json.OpenSellerCode,
		OpenSellerName:  json.OpenSellerName,
		CloseSellerCode: json.CloseSellerCode,
		CloseSellerName: json.CloseSellerName,
		ReceiptNo:       json.ReceiptNo,
		ReceiptTypeCode: json.ReceiptTypeCode,
		SessionNo:       json.SessionNo,
		SummForD:        json.SummForD,
		SummWD:          json.SummWd,
		TotalDiscSumm:   json.TotalDiscSumm,
		TotalDiscValue:  json.TotalDiscValue,
		Cards:           cards,
		Discounts:       discounts,
	}
	DB.Create(&document)
}
