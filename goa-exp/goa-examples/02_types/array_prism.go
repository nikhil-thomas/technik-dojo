package main

import (
	"github.com/goadesign/goa"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/app"
)

// ArrayPrismController implements the array_prism resource.
type ArrayPrismController struct {
	*goa.Controller
}

// NewArrayPrismController creates a array_prism controller.
func NewArrayPrismController(service *goa.Service) *ArrayPrismController {
	return &ArrayPrismController{Controller: service.NewController("ArrayPrismController")}
}

// Create runs the create action.
func (c *ArrayPrismController) Create(ctx *app.CreateArrayPrismContext) error {
	// ArrayPrismController_Create: start_implement

	// Put your logic here

	return nil
	// ArrayPrismController_Create: end_implement
}

// Show runs the show action.
func (c *ArrayPrismController) Show(ctx *app.ShowArrayPrismContext) error {
	// ArrayPrismController_Show: start_implement

	// Put your logic here

	res := &app.GoadesignExamplesArrayprism{}
	return ctx.OK(res)
	// ArrayPrismController_Show: end_implement
}
