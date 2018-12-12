package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("types", func() {
	Description("Fake API ti showcase the types of DSL")
	Host("localhost")
})
