package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("hash_prism", func() {
	BasePath("/hash")
	Description("Resource hash_prism contains artificial actions that showcase the use of the hash type in DSL")
	Action("show", func() {
		Description("Action show returns a media type with one hash memmerb per primitive type, each using the primitive type as value and as a key")
		Routing(GET(""))
		Response(OK, HashPrismMedia)
	})
	Action("create", func() {
		Description("Action create acepts a payload with one hash member per primitive type, each using the type as value and String as key")
		Routing(POST(""))
		Payload(func() {
			Member("bool_val_hash", HashOf(String, Boolean), "Hash with Boolean value member")
			Member("int_val_hash", HashOf(String, Integer), "Hash with Integer value member")
			Member("num_val_hash", HashOf(String, Number), "Hash with Number value member")
			Member("string_val_hash", HashOf(String, String), "Hash with String value member")
			Member("data_time_val_hash", HashOf(String, DateTime), "Hash with DateTime value member")
			Member("uuid_val_hash", HashOf(String, UUID), "Hash with UUID value member")
			Member("any_val_hash", HashOf(String, Any), "Hash with Any value member")
		})
	})
})

var HashPrismMedia = MediaType("application/vnd.goadesign.examples.hashprism", func() {
	Description("HashPrismMedia is a media type with one hash member per primitive type: type as value, and string as key")
	Attributes(func() {
		Attribute("bool_val_hash", HashOf(String, Boolean), "Hash with Boolean value member")
		Attribute("int_val_hash", HashOf(String, Integer), "Hash with Integer value member")
		Attribute("num_val_hash", HashOf(String, Number), "Hash with Number value member")
		Attribute("string_val_hash", HashOf(String, String), "Hash with String value member")
		Attribute("date_time_val_hash", HashOf(String, DateTime), "Hash with DateTime value member")
		Attribute("uuid_val_hash", HashOf(String, UUID), "Hash with UUID value member")
		Attribute("any_val_hash", HashOf(String, Any), "Hash with Any value member")
	})
	View("default", func() {
		Attribute("bool_value_hash")
		Attribute("int_value_hash")
		Attribute("num_value_hash")
		Attribute("string_value_hash")
		Attribute("date_time_value_hash")
		Attribute("uuid_value_hash")
		Attribute("any_value_hash")
	})
})
