package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("array_prism", func() {
	BasePath("/array")
	Description("resource array_prism contains artificial actions that show use array in goa")
	Action("show", func() {
		Description(`Action accepts one queerystring array parameter for each primitive type.Description
Array parameters are constructed using all the values given to the same querstring key. e.g:

    GET /array?uint_param=1&int_param=2
`)
		Routing(GET(""))
		Params(func() {
			Param("bool_array", ArrayOf(Boolean), "Boolean array parameter")
			Param("int_array", ArrayOf(Integer), "Integer array parameter")
			Param("num_array", ArrayOf(Number), "Number array parameter")
			Param("string_array", ArrayOf(String), "String array parameter")
			Param("date_time_array", ArrayOf(DateTime), "DateTime array parameter")
			Param("uuid_array", ArrayOf(UUID), "UUID array parameter")
			Param("any_array", ArrayOf(Any), "Any array parameter")
		})
		Response(OK)
	})
})
