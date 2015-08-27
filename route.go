package main

import (
	"./controllers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func apiv1() {

	m.Group("/workplaces", func(router martini.Router) {

		router.Post("/auth", binding.Json(controllers.Auth{}), controllers.WorkplaceAuth)
		router.Post("/test", binding.Json(controllers.Test{}), controllers.WorkplacePing)
		//add
		//update
		//delete
		//import

	}, render.Renderer())

}
