package controllers

import (
	. "../db"
	. "../models"
	"../utils/jwt"
	"fmt"

	"github.com/martini-contrib/render"
)

//Autentificate
type Auth struct {
	FrontolCode int    `json:"frontol_code"`
	Uuid        string `json:"uuid"`
}

func WorkplaceAuth(r render.Render, json Auth) {
	fmt.Printf("%+v\n", json)

	workplace := Workplace{}

	DB.Where("uuid = ?", json.Uuid).First(&workplace)
	fmt.Printf("%+v\n", workplace)
	if json.FrontolCode == workplace.FrontolCode && json.Uuid == workplace.Uuid {
		r.JSON(200, map[string]interface{}{"token": jwt.CreateToken(workplace.Uuid)})
	} else {
		r.JSON(200, map[string]interface{}{"error": "invalid ID or pincode"})
	}
}

//ping function
type Test struct {
	Token string `"json:token"`
}

func WorkplacePing(r render.Render, json Test) {
	fmt.Printf("%+v\n", json)
	uuid, err := jwt.ParseToken(json.Token)
	if err == nil {
		r.JSON(200, map[string]interface{}{"message": "your uuid is " + uuid})
	} else {
		r.JSON(200, map[string]interface{}{"error": err})
	}
}

//add new workplace
type Add struct {
	Name        string
	FrontolCode int
}

func WorkplaceAdd(r render.Render, json Add) {
	fmt.Printf("%+v\n", json)
	workplace := Workplace{}
	workplace.Name = json.Name
	workplace.FrontolCode = json.FrontolCode

	DB.Create(&workplace)

}
