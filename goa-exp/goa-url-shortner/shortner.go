package main

import (
	"github.com/goadesign/goa"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner/app"
)

// ShortnerController implements the shortner resource.
type ShortnerController struct {
	*goa.Controller
}

// NewShortnerController creates a shortner controller.
func NewShortnerController(service *goa.Service) *ShortnerController {
	return &ShortnerController{Controller: service.NewController("ShortnerController")}
}

// Analytics runs the analytics action.
func (c *ShortnerController) Analytics(ctx *app.AnalyticsShortnerContext) error {
	// ShortnerController_Analytics: start_implement

	// Put your logic here

	res := &app.Analytics{}
	return ctx.OK(res)
	// ShortnerController_Analytics: end_implement
}

// Create runs the create action.
func (c *ShortnerController) Create(ctx *app.CreateShortnerContext) error {
	// ShortnerController_Create: start_implement

	// Put your logic here

	return nil
	// ShortnerController_Create: end_implement
}

// Delete runs the delete action.
func (c *ShortnerController) Delete(ctx *app.DeleteShortnerContext) error {
	// ShortnerController_Delete: start_implement

	// Put your logic here

	return nil
	// ShortnerController_Delete: end_implement
}

// Get runs the get action.
func (c *ShortnerController) Get(ctx *app.GetShortnerContext) error {
	// ShortnerController_Get: start_implement

	// Put your logic here

	res := &app.URL{}
	return ctx.OK(res)
	// ShortnerController_Get: end_implement
}

// Update runs the update action.
func (c *ShortnerController) Update(ctx *app.UpdateShortnerContext) error {
	// ShortnerController_Update: start_implement

	// Put your logic here

	return nil
	// ShortnerController_Update: end_implement
}
