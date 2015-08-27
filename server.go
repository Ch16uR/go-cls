package main

import (
	"github.com/go-martini/martini"
)

var m = martini.Classic()

func server() {

	apiv1()

	m.Run()
}
