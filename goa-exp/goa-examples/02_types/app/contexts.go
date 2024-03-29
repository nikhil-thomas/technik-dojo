// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "types": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/design
// --out=$(GOPATH)/src/github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
)

// CreateArrayPrismContext provides the array_prism create action context.
type CreateArrayPrismContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateArrayPrismPayload
}

// NewCreateArrayPrismContext parses the incoming request URL and body, performs validations and creates the
// context used by the array_prism controller create action.
func NewCreateArrayPrismContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateArrayPrismContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateArrayPrismContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createArrayPrismPayload is the array_prism create action payload.
type createArrayPrismPayload struct {
	// Any array member
	AnyArray []interface{} `form:"any_array,omitempty" json:"any_array,omitempty" yaml:"any_array,omitempty" xml:"any_array,omitempty"`
	// Boolean array member
	BoolArray []bool `form:"bool_array,omitempty" json:"bool_array,omitempty" yaml:"bool_array,omitempty" xml:"bool_array,omitempty"`
	// DateTime array member
	DateTimeArray []time.Time `form:"date_time_array,omitempty" json:"date_time_array,omitempty" yaml:"date_time_array,omitempty" xml:"date_time_array,omitempty"`
	// Integer array member
	IntArray []int `form:"int_array,omitempty" json:"int_array,omitempty" yaml:"int_array,omitempty" xml:"int_array,omitempty"`
	// Number array member
	NumArray []float64 `form:"num_array,omitempty" json:"num_array,omitempty" yaml:"num_array,omitempty" xml:"num_array,omitempty"`
	// String array member
	StringArray []string `form:"string_array,omitempty" json:"string_array,omitempty" yaml:"string_array,omitempty" xml:"string_array,omitempty"`
	// UUID array member
	UUIDArray []uuid.UUID `form:"uuid_array,omitempty" json:"uuid_array,omitempty" yaml:"uuid_array,omitempty" xml:"uuid_array,omitempty"`
}

// Publicize creates CreateArrayPrismPayload from createArrayPrismPayload
func (payload *createArrayPrismPayload) Publicize() *CreateArrayPrismPayload {
	var pub CreateArrayPrismPayload
	if payload.AnyArray != nil {
		pub.AnyArray = payload.AnyArray
	}
	if payload.BoolArray != nil {
		pub.BoolArray = payload.BoolArray
	}
	if payload.DateTimeArray != nil {
		pub.DateTimeArray = payload.DateTimeArray
	}
	if payload.IntArray != nil {
		pub.IntArray = payload.IntArray
	}
	if payload.NumArray != nil {
		pub.NumArray = payload.NumArray
	}
	if payload.StringArray != nil {
		pub.StringArray = payload.StringArray
	}
	if payload.UUIDArray != nil {
		pub.UUIDArray = payload.UUIDArray
	}
	return &pub
}

// CreateArrayPrismPayload is the array_prism create action payload.
type CreateArrayPrismPayload struct {
	// Any array member
	AnyArray []interface{} `form:"any_array,omitempty" json:"any_array,omitempty" yaml:"any_array,omitempty" xml:"any_array,omitempty"`
	// Boolean array member
	BoolArray []bool `form:"bool_array,omitempty" json:"bool_array,omitempty" yaml:"bool_array,omitempty" xml:"bool_array,omitempty"`
	// DateTime array member
	DateTimeArray []time.Time `form:"date_time_array,omitempty" json:"date_time_array,omitempty" yaml:"date_time_array,omitempty" xml:"date_time_array,omitempty"`
	// Integer array member
	IntArray []int `form:"int_array,omitempty" json:"int_array,omitempty" yaml:"int_array,omitempty" xml:"int_array,omitempty"`
	// Number array member
	NumArray []float64 `form:"num_array,omitempty" json:"num_array,omitempty" yaml:"num_array,omitempty" xml:"num_array,omitempty"`
	// String array member
	StringArray []string `form:"string_array,omitempty" json:"string_array,omitempty" yaml:"string_array,omitempty" xml:"string_array,omitempty"`
	// UUID array member
	UUIDArray []uuid.UUID `form:"uuid_array,omitempty" json:"uuid_array,omitempty" yaml:"uuid_array,omitempty" xml:"uuid_array,omitempty"`
}

// ShowArrayPrismContext provides the array_prism show action context.
type ShowArrayPrismContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AnyArray      []interface{}
	BoolArray     []bool
	DateTimeArray []time.Time
	IntArray      []int
	NumArray      []float64
	StringArray   []string
	UUIDArray     []uuid.UUID
}

// NewShowArrayPrismContext parses the incoming request URL and body, performs validations and creates the
// context used by the array_prism controller show action.
func NewShowArrayPrismContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowArrayPrismContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowArrayPrismContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAnyArray := req.Params["any_array"]
	if len(paramAnyArray) > 0 {
		params := make([]interface{}, len(paramAnyArray))
		for i, rawAnyArray := range paramAnyArray {
			params[i] = rawAnyArray
		}
		rctx.AnyArray = params
	}
	paramBoolArray := req.Params["bool_array"]
	if len(paramBoolArray) > 0 {
		params := make([]bool, len(paramBoolArray))
		for i, rawBoolArray := range paramBoolArray {
			if boolArray, err2 := strconv.ParseBool(rawBoolArray); err2 == nil {
				params[i] = boolArray
			} else {
				err = goa.MergeErrors(err, goa.InvalidParamTypeError("bool_array", rawBoolArray, "boolean"))
			}
		}
		rctx.BoolArray = params
	}
	paramDateTimeArray := req.Params["date_time_array"]
	if len(paramDateTimeArray) > 0 {
		params := make([]time.Time, len(paramDateTimeArray))
		for i, rawDateTimeArray := range paramDateTimeArray {
			if dateTimeArray, err2 := time.Parse(time.RFC3339, rawDateTimeArray); err2 == nil {
				params[i] = dateTimeArray
			} else {
				err = goa.MergeErrors(err, goa.InvalidParamTypeError("date_time_array", rawDateTimeArray, "datetime"))
			}
		}
		rctx.DateTimeArray = params
	}
	paramIntArray := req.Params["int_array"]
	if len(paramIntArray) > 0 {
		params := make([]int, len(paramIntArray))
		for i, rawIntArray := range paramIntArray {
			if intArray, err2 := strconv.Atoi(rawIntArray); err2 == nil {
				params[i] = intArray
			} else {
				err = goa.MergeErrors(err, goa.InvalidParamTypeError("int_array", rawIntArray, "integer"))
			}
		}
		rctx.IntArray = params
	}
	paramNumArray := req.Params["num_array"]
	if len(paramNumArray) > 0 {
		params := make([]float64, len(paramNumArray))
		for i, rawNumArray := range paramNumArray {
			if numArray, err2 := strconv.ParseFloat(rawNumArray, 64); err2 == nil {
				params[i] = numArray
			} else {
				err = goa.MergeErrors(err, goa.InvalidParamTypeError("num_array", rawNumArray, "number"))
			}
		}
		rctx.NumArray = params
	}
	paramStringArray := req.Params["string_array"]
	if len(paramStringArray) > 0 {
		params := paramStringArray
		rctx.StringArray = params
	}
	paramUUIDArray := req.Params["uuid_array"]
	if len(paramUUIDArray) > 0 {
		params := make([]uuid.UUID, len(paramUUIDArray))
		for i, rawUUIDArray := range paramUUIDArray {
			if uuidArray, err2 := uuid.FromString(rawUUIDArray); err2 == nil {
				params[i] = uuidArray
			} else {
				err = goa.MergeErrors(err, goa.InvalidParamTypeError("uuid_array", rawUUIDArray, "uuid"))
			}
		}
		rctx.UUIDArray = params
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowArrayPrismContext) OK(r *GoadesignExamplesArrayprism) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goadesign.examples.arrayprism")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreateHashPrismContext provides the hash_prism create action context.
type CreateHashPrismContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateHashPrismPayload
}

// NewCreateHashPrismContext parses the incoming request URL and body, performs validations and creates the
// context used by the hash_prism controller create action.
func NewCreateHashPrismContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateHashPrismContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateHashPrismContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createHashPrismPayload is the hash_prism create action payload.
type createHashPrismPayload struct {
	// Hash with Any value member
	AnyValHash map[string]interface{} `form:"any_val_hash,omitempty" json:"any_val_hash,omitempty" yaml:"any_val_hash,omitempty" xml:"any_val_hash,omitempty"`
	// Hash with Boolean value member
	BoolValHash map[string]bool `form:"bool_val_hash,omitempty" json:"bool_val_hash,omitempty" yaml:"bool_val_hash,omitempty" xml:"bool_val_hash,omitempty"`
	// Hash with DateTime value member
	DataTimeValHash map[string]time.Time `form:"data_time_val_hash,omitempty" json:"data_time_val_hash,omitempty" yaml:"data_time_val_hash,omitempty" xml:"data_time_val_hash,omitempty"`
	// Hash with Integer value member
	IntValHash map[string]int `form:"int_val_hash,omitempty" json:"int_val_hash,omitempty" yaml:"int_val_hash,omitempty" xml:"int_val_hash,omitempty"`
	// Hash with Number value member
	NumValHash map[string]float64 `form:"num_val_hash,omitempty" json:"num_val_hash,omitempty" yaml:"num_val_hash,omitempty" xml:"num_val_hash,omitempty"`
	// Hash with String value member
	StringValHash map[string]string `form:"string_val_hash,omitempty" json:"string_val_hash,omitempty" yaml:"string_val_hash,omitempty" xml:"string_val_hash,omitempty"`
	// Hash with UUID value member
	UUIDValHash map[string]uuid.UUID `form:"uuid_val_hash,omitempty" json:"uuid_val_hash,omitempty" yaml:"uuid_val_hash,omitempty" xml:"uuid_val_hash,omitempty"`
}

// Publicize creates CreateHashPrismPayload from createHashPrismPayload
func (payload *createHashPrismPayload) Publicize() *CreateHashPrismPayload {
	var pub CreateHashPrismPayload
	if payload.AnyValHash != nil {
		pub.AnyValHash = payload.AnyValHash
	}
	if payload.BoolValHash != nil {
		pub.BoolValHash = payload.BoolValHash
	}
	if payload.DataTimeValHash != nil {
		pub.DataTimeValHash = payload.DataTimeValHash
	}
	if payload.IntValHash != nil {
		pub.IntValHash = payload.IntValHash
	}
	if payload.NumValHash != nil {
		pub.NumValHash = payload.NumValHash
	}
	if payload.StringValHash != nil {
		pub.StringValHash = payload.StringValHash
	}
	if payload.UUIDValHash != nil {
		pub.UUIDValHash = payload.UUIDValHash
	}
	return &pub
}

// CreateHashPrismPayload is the hash_prism create action payload.
type CreateHashPrismPayload struct {
	// Hash with Any value member
	AnyValHash map[string]interface{} `form:"any_val_hash,omitempty" json:"any_val_hash,omitempty" yaml:"any_val_hash,omitempty" xml:"any_val_hash,omitempty"`
	// Hash with Boolean value member
	BoolValHash map[string]bool `form:"bool_val_hash,omitempty" json:"bool_val_hash,omitempty" yaml:"bool_val_hash,omitempty" xml:"bool_val_hash,omitempty"`
	// Hash with DateTime value member
	DataTimeValHash map[string]time.Time `form:"data_time_val_hash,omitempty" json:"data_time_val_hash,omitempty" yaml:"data_time_val_hash,omitempty" xml:"data_time_val_hash,omitempty"`
	// Hash with Integer value member
	IntValHash map[string]int `form:"int_val_hash,omitempty" json:"int_val_hash,omitempty" yaml:"int_val_hash,omitempty" xml:"int_val_hash,omitempty"`
	// Hash with Number value member
	NumValHash map[string]float64 `form:"num_val_hash,omitempty" json:"num_val_hash,omitempty" yaml:"num_val_hash,omitempty" xml:"num_val_hash,omitempty"`
	// Hash with String value member
	StringValHash map[string]string `form:"string_val_hash,omitempty" json:"string_val_hash,omitempty" yaml:"string_val_hash,omitempty" xml:"string_val_hash,omitempty"`
	// Hash with UUID value member
	UUIDValHash map[string]uuid.UUID `form:"uuid_val_hash,omitempty" json:"uuid_val_hash,omitempty" yaml:"uuid_val_hash,omitempty" xml:"uuid_val_hash,omitempty"`
}

// ShowHashPrismContext provides the hash_prism show action context.
type ShowHashPrismContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowHashPrismContext parses the incoming request URL and body, performs validations and creates the
// context used by the hash_prism controller show action.
func NewShowHashPrismContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowHashPrismContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowHashPrismContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowHashPrismContext) OK(r *GoadesignExamplesHashprism) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goadesign.examples.hashprism")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// CreatePrismContext provides the prism create action context.
type CreatePrismContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreatePrismPayload
}

// NewCreatePrismContext parses the incoming request URL and body, performs validations and creates the
// context used by the prism controller create action.
func NewCreatePrismContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreatePrismContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreatePrismContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createPrismPayload is the prism create action payload.
type createPrismPayload struct {
	// Any member
	AnyMember interface{} `form:"any_member,omitempty" json:"any_member,omitempty" yaml:"any_member,omitempty" xml:"any_member,omitempty"`
	// Boolean member
	BoolMember *bool `form:"bool_member,omitempty" json:"bool_member,omitempty" yaml:"bool_member,omitempty" xml:"bool_member,omitempty"`
	// DateTime member
	DateTimeMember *time.Time `form:"date_time_member,omitempty" json:"date_time_member,omitempty" yaml:"date_time_member,omitempty" xml:"date_time_member,omitempty"`
	// Integer member
	IntMember *int `form:"int_member,omitempty" json:"int_member,omitempty" yaml:"int_member,omitempty" xml:"int_member,omitempty"`
	// Number member
	NumMember *float64 `form:"num_member,omitempty" json:"num_member,omitempty" yaml:"num_member,omitempty" xml:"num_member,omitempty"`
	// String member
	StringMember *string `form:"string_member,omitempty" json:"string_member,omitempty" yaml:"string_member,omitempty" xml:"string_member,omitempty"`
	// UUID member
	UUIDMember *uuid.UUID `form:"uuid_member,omitempty" json:"uuid_member,omitempty" yaml:"uuid_member,omitempty" xml:"uuid_member,omitempty"`
}

// Publicize creates CreatePrismPayload from createPrismPayload
func (payload *createPrismPayload) Publicize() *CreatePrismPayload {
	var pub CreatePrismPayload
	if payload.AnyMember != nil {
		pub.AnyMember = payload.AnyMember
	}
	if payload.BoolMember != nil {
		pub.BoolMember = payload.BoolMember
	}
	if payload.DateTimeMember != nil {
		pub.DateTimeMember = payload.DateTimeMember
	}
	if payload.IntMember != nil {
		pub.IntMember = payload.IntMember
	}
	if payload.NumMember != nil {
		pub.NumMember = payload.NumMember
	}
	if payload.StringMember != nil {
		pub.StringMember = payload.StringMember
	}
	if payload.UUIDMember != nil {
		pub.UUIDMember = payload.UUIDMember
	}
	return &pub
}

// CreatePrismPayload is the prism create action payload.
type CreatePrismPayload struct {
	// Any member
	AnyMember interface{} `form:"any_member,omitempty" json:"any_member,omitempty" yaml:"any_member,omitempty" xml:"any_member,omitempty"`
	// Boolean member
	BoolMember *bool `form:"bool_member,omitempty" json:"bool_member,omitempty" yaml:"bool_member,omitempty" xml:"bool_member,omitempty"`
	// DateTime member
	DateTimeMember *time.Time `form:"date_time_member,omitempty" json:"date_time_member,omitempty" yaml:"date_time_member,omitempty" xml:"date_time_member,omitempty"`
	// Integer member
	IntMember *int `form:"int_member,omitempty" json:"int_member,omitempty" yaml:"int_member,omitempty" xml:"int_member,omitempty"`
	// Number member
	NumMember *float64 `form:"num_member,omitempty" json:"num_member,omitempty" yaml:"num_member,omitempty" xml:"num_member,omitempty"`
	// String member
	StringMember *string `form:"string_member,omitempty" json:"string_member,omitempty" yaml:"string_member,omitempty" xml:"string_member,omitempty"`
	// UUID member
	UUIDMember *uuid.UUID `form:"uuid_member,omitempty" json:"uuid_member,omitempty" yaml:"uuid_member,omitempty" xml:"uuid_member,omitempty"`
}

// ShowPrismContext provides the prism show action context.
type ShowPrismContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AnyParam      interface{}
	BoolParam     *bool
	DateTimeParam *time.Time
	IntParam      *int
	NumParam      *float64
	StringParam   *string
	UUIDParam     *uuid.UUID
}

// NewShowPrismContext parses the incoming request URL and body, performs validations and creates the
// context used by the prism controller show action.
func NewShowPrismContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowPrismContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowPrismContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAnyParam := req.Params["any_param"]
	if len(paramAnyParam) > 0 {
		rawAnyParam := paramAnyParam[0]
		rctx.AnyParam = rawAnyParam
	}
	paramBoolParam := req.Params["bool_param"]
	if len(paramBoolParam) > 0 {
		rawBoolParam := paramBoolParam[0]
		if boolParam, err2 := strconv.ParseBool(rawBoolParam); err2 == nil {
			tmp6 := &boolParam
			rctx.BoolParam = tmp6
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bool_param", rawBoolParam, "boolean"))
		}
	}
	paramDateTimeParam := req.Params["date_time_param"]
	if len(paramDateTimeParam) > 0 {
		rawDateTimeParam := paramDateTimeParam[0]
		if dateTimeParam, err2 := time.Parse(time.RFC3339, rawDateTimeParam); err2 == nil {
			tmp7 := &dateTimeParam
			rctx.DateTimeParam = tmp7
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("date_time_param", rawDateTimeParam, "datetime"))
		}
	}
	paramIntParam := req.Params["int_param"]
	if len(paramIntParam) > 0 {
		rawIntParam := paramIntParam[0]
		if intParam, err2 := strconv.Atoi(rawIntParam); err2 == nil {
			tmp9 := intParam
			tmp8 := &tmp9
			rctx.IntParam = tmp8
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("int_param", rawIntParam, "integer"))
		}
	}
	paramNumParam := req.Params["num_param"]
	if len(paramNumParam) > 0 {
		rawNumParam := paramNumParam[0]
		if numParam, err2 := strconv.ParseFloat(rawNumParam, 64); err2 == nil {
			tmp10 := &numParam
			rctx.NumParam = tmp10
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("num_param", rawNumParam, "number"))
		}
	}
	paramStringParam := req.Params["string_param"]
	if len(paramStringParam) > 0 {
		rawStringParam := paramStringParam[0]
		rctx.StringParam = &rawStringParam
	}
	paramUUIDParam := req.Params["uuid_param"]
	if len(paramUUIDParam) > 0 {
		rawUUIDParam := paramUUIDParam[0]
		if uuidParam, err2 := uuid.FromString(rawUUIDParam); err2 == nil {
			tmp11 := &uuidParam
			rctx.UUIDParam = tmp11
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("uuid_param", rawUUIDParam, "uuid"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowPrismContext) OK(r *GoadesignExamplesPrism) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goadesign.examples.prism")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
