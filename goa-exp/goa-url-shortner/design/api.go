package design

import (
	"net/http"
	"os"

	// . "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("URL shortner API", func() {
	Title("The URL Shortner API")
	Description("An API for a URL shortner")
	Contact(func() {
		Name(os.Getenv("DEV_NAME"))
		Email(os.Getenv("DEV_EMAIL"))
	})
	Host(os.Getenv("HOST"))
	Scheme("http")
	BasePath("/api/")
	Origin("*", func() {
		Headers("Content-Type")
		Methods(
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
		)
	})
	Consumes("application/json")
	Produces("application/json")
})
