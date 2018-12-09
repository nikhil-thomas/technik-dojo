//go:generate goagen bootstrap -d github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/app"
)

func main() {
	// Create service
	service := goa.New("types")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "array_prism" controller
	c := NewArrayPrismController(service)
	app.MountArrayPrismController(service, c)
	// Mount "hash_prism" controller
	c2 := NewHashPrismController(service)
	app.MountHashPrismController(service, c2)
	// Mount "prism" controller
	c3 := NewPrismController(service)
	app.MountPrismController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
