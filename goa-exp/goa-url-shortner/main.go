//go:generate goagen bootstrap -d github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner/app"
)

func main() {
	// Create service
	service := goa.New("URL shortner API")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "shortner" controller
	c := NewShortnerController(service)
	app.MountShortnerController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
