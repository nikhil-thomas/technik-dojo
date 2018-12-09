package main

import (
	"github.com/goadesign/goa"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/app"
)

// HashPrismController implements the hash_prism resource.
type HashPrismController struct {
	*goa.Controller
}

// NewHashPrismController creates a hash_prism controller.
func NewHashPrismController(service *goa.Service) *HashPrismController {
	return &HashPrismController{Controller: service.NewController("HashPrismController")}
}

// Create runs the create action.
func (c *HashPrismController) Create(ctx *app.CreateHashPrismContext) error {
	// HashPrismController_Create: start_implement

	// Put your logic here

	return nil
	// HashPrismController_Create: end_implement
}

// Show runs the show action.
func (c *HashPrismController) Show(ctx *app.ShowHashPrismContext) error {
	// HashPrismController_Show: start_implement

	// Put your logic here

	res := &app.GoadesignExamplesHashprism{}
	return ctx.OK(res)
	// HashPrismController_Show: end_implement
}
