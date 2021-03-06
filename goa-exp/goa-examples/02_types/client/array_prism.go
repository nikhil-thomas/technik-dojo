// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "types": array_prism Resource Client
//
// Command:
// $ goagen
// --design=github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types/design
// --out=$(GOPATH)/src/github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/02_types
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	uuid "github.com/goadesign/goa/uuid"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

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

// CreateArrayPrismPath computes a request path to the create action of array_prism.
func CreateArrayPrismPath() string {

	return fmt.Sprintf("/array")
}

// Action create accepts a payload with one array member of each primitive type
func (c *Client) CreateArrayPrism(ctx context.Context, path string, payload *CreateArrayPrismPayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateArrayPrismRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateArrayPrismRequest create the request corresponding to the create action endpoint of the array_prism resource.
func (c *Client) NewCreateArrayPrismRequest(ctx context.Context, path string, payload *CreateArrayPrismPayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowArrayPrismPath computes a request path to the show action of array_prism.
func ShowArrayPrismPath() string {

	return fmt.Sprintf("/array")
}

// Action accepts one queerystring array parameter for each primitive type.Description
// Array parameters are constructed using all the values given to the same querstring key. e.g:
//
// GET /array?uint_param=1&int_param=2
//
func (c *Client) ShowArrayPrism(ctx context.Context, path string, anyArray []interface{}, boolArray []bool, dateTimeArray []time.Time, intArray []int, numArray []float64, stringArray []string, uuidArray []uuid.UUID) (*http.Response, error) {
	req, err := c.NewShowArrayPrismRequest(ctx, path, anyArray, boolArray, dateTimeArray, intArray, numArray, stringArray, uuidArray)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowArrayPrismRequest create the request corresponding to the show action endpoint of the array_prism resource.
func (c *Client) NewShowArrayPrismRequest(ctx context.Context, path string, anyArray []interface{}, boolArray []bool, dateTimeArray []time.Time, intArray []int, numArray []float64, stringArray []string, uuidArray []uuid.UUID) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	for _, p := range anyArray {
		tmp17 := fmt.Sprintf("%v", p)
		values.Add("any_array", tmp17)
	}
	for _, p := range boolArray {
		tmp18 := strconv.FormatBool(p)
		values.Add("bool_array", tmp18)
	}
	for _, p := range dateTimeArray {
		tmp19 := p.Format(time.RFC3339)
		values.Add("date_time_array", tmp19)
	}
	for _, p := range intArray {
		tmp20 := strconv.Itoa(p)
		values.Add("int_array", tmp20)
	}
	for _, p := range numArray {
		tmp21 := strconv.FormatFloat(p, 'f', -1, 64)
		values.Add("num_array", tmp21)
	}
	for _, p := range stringArray {
		tmp22 := p
		values.Add("string_array", tmp22)
	}
	for _, p := range uuidArray {
		tmp23 := p.String()
		values.Add("uuid_array", tmp23)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
