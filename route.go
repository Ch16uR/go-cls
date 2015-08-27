package main

import (
	"./controllers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func apiv1() {
	url := "/api/v1"
	m.Group(url+"/workplace", func(router martini.Router) {

		router.Post("/auth", binding.Json(controllers.Auth{}), controllers.WorkplaceAuth)
		router.Post("/test", binding.Json(controllers.Test{}), controllers.WorkplacePing)
		router.Post("/add", binding.Json(controllers.Add{}), controllers.WorkplaceAdd)
		//update
		//delete
		//import

	}, render.Renderer())

}
