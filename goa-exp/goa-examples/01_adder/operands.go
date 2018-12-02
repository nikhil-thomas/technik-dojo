package main

import (
	"strconv"

	"github.com/goadesign/goa"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/01_adder/app"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	// OperandsController_Add: start_implement

	// Put your logic here

	return ctx.OK([]byte(strconv.Itoa(ctx.Left + ctx.Right)))
}
