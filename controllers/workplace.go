package controllers

import (
	. "../db"
	. "../models"
	"../utils/jwt"
	"../utils/uuid"
	"fmt"
	"github.com/martini-contrib/render"
	"time"
)

//Autentificate
type InputWorkplaceAuth struct {
	FrontolCode int    `json:"frontol_code"`
	Uuid        string `json:"uuid"`
}

func WorkplaceAuth(r render.Render, json InputWorkplaceAuth) {
	fmt.Printf("%+v\n", json)
	if len(json.Uuid) == 0 {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid JSON"})
		return
	}

	workplace := Workplace{}

	DB.Where("uuid = ?", json.Uuid).First(&workplace)
	fmt.Printf("%+v\n", workplace)
	if json.FrontolCode == workplace.FrontolCode && json.Uuid == workplace.Uuid {
		r.JSON(200, map[string]interface{}{"status": "ok", "token": jwt.CreateToken(workplace.ID, workplace.CompanyID), "lifetime": time.Now().Add(time.Hour * 24).Unix()})
	} else {
		r.JSON(200, map[string]interface{}{"status": "error", "message": "Invalid UUID"})
	}
}

//ping function
type Test struct {
	Token string `"json:token"`
}

/*func WorkplacePing(r render.Render, json Test) {
	fmt.Printf("%+v\n", json)
	uuid, err := jwt.ParseToken(json.Token)
	if err == nil {
		r.JSON(200, map[string]interface{}{"message": "your uuid is " + uuid})
	} else {
		r.JSON(200, map[string]interface{}{"error": err})
	}
}*/

//add new workplace
type Add struct {
	Name        string `"json:name"`
	FrontolCode int    `json:"frontol_code"`
}

func WorkplaceAdd(r render.Render, json Add) {
	fmt.Printf("%+v\n", json)
	workplace := Workplace{}
	workplace.Name = json.Name
	workplace.FrontolCode = json.FrontolCode
	workplace.Uuid = uuid.CreateUuid()

	DB.Create(&workplace)

	r.JSON(200, map[string]interface{}{"message": "your uuid is"})
}
