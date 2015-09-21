package controllers

import (
	. "../db"
	. "../models"
	"../utils/jwt"
	"fmt"
	"github.com/martini-contrib/render"
	"time"
)

type InputPromocodeCalc struct {
	Token    string `json:"token"`
	Value    string `json:"value"`
	OpenedAt string `json:"opened_at"`
}

func PromocodeCalc(r render.Render, json InputPromocodeCalc) {
	fmt.Printf("%+v\n", json)
	id, company, err := jwt.ParseToken(json.Token)
	if err != nil {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid Token"})
		return
	}

	datetime, err := time.Parse(time.RFC3339, json.OpenedAt)
	if err != nil {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid Datetime"})
		return
	}
	fmt.Println(datetime.UTC())
	promocode := Promocode{}

	DB.Where("value = ? and company_id = ?", json.Value, company).First(&promocode)

	fmt.Printf("%+v\n", promocode)
	if promocode.ID == 0 {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Купона " + json.Value + " нет в базе данных."})
		return
	}
	if datetime.After(promocode.ExpiredAt) {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Срок действия этого купона истек."})
		return
	}

	if err != nil {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid Token"})
	} else {
		r.JSON(200, map[string]interface{}{"status": "ok", "message": "", "discount_value": promocode.Discount, "discount_type": promocode.PromocodeDiscountTypeID})
	}

}
