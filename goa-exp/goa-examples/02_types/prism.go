package main

import (
	"github.com/goadesign/goa"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/app"
)

// PrismController implements the prism resource.
type PrismController struct {
	*goa.Controller
}

// NewPrismController creates a prism controller.
func NewPrismController(service *goa.Service) *PrismController {
	return &PrismController{Controller: service.NewController("PrismController")}
}

// Create runs the create action.
func (c *PrismController) Create(ctx *app.CreatePrismContext) error {
	// PrismController_Create: start_implement

	// Put your logic here

	return nil
	// PrismController_Create: end_implement
}

// Show runs the show action.
func (c *PrismController) Show(ctx *app.ShowPrismContext) error {
	// PrismController_Show: start_implement

	// Put your logic here

	res := &app.GoadesignExamplesPrism{}
	return ctx.OK(res)
	// PrismController_Show: end_implement
}
