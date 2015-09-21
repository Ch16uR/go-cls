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
		router.Post("/auth", binding.Json(controllers.InputWorkplaceAuth{}), controllers.WorkplaceAuth)
		//router.Post("/test", binding.Json(controllers.Test{}), controllers.WorkplacePing)
		//router.Post("/add", binding.Json(controllers.Add{}), controllers.WorkplaceAdd)
		//update
		//delete
		//import

	}, render.Renderer())

	m.Group(url+"/promocode", func(router martini.Router) {

		router.Post("/calc", binding.Json(controllers.InputPromocodeCalc{}), controllers.PromocodeCalc)
		//update
		//delete
		//import

	}, render.Renderer())

	m.Group(url+"/document", func(router martini.Router) {

		router.Post("/add", binding.Json(controllers.InputDocumentAdd{}), controllers.DocumentAdd)
		//update
		//delete
		//import

	}, render.Renderer())

}
