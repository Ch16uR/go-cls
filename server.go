package main

import (
	"github.com/go-martini/martini"
)

var m = martini.Classic()

func server() {

	apiv1()
	static := martini.Static("ui", martini.StaticOptions{Fallback: "/index.html", Exclude: "/api/v"})
	m.Use(static)

	m.Run()
}
