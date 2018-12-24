// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "URL shortner API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner/design
// --out=$(GOPATH)/src/github.com/nikhil-thomas/technik-dojo/goa-exp/goa-url-shortner
// --version=v1.3.1

package client

import (
	"net/http"
)

// Url analytics (default view)
//
// Identifier: application/vnd.analytics+json; view=default
type Analytics struct {
	Hits *int `form:"hits,omitempty" json:"hits,omitempty" yaml:"hits,omitempty" xml:"hits,omitempty"`
}

// DecodeAnalytics decodes the Analytics instance encoded in resp body.
func (c *Client) DecodeAnalytics(resp *http.Response) (*Analytics, error) {
	var decoded Analytics
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A URL (default view)
//
// Identifier: application/vnd.url+json; view=default
type URL struct {
	// URL ID
	ID *int `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// URL path key
	Path *string `form:"path,omitempty" json:"path,omitempty" yaml:"path,omitempty" xml:"path,omitempty"`
	// External URL
	URL *string `form:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty" xml:"url,omitempty"`
}

// DecodeURL decodes the URL instance encoded in resp body.
func (c *Client) DecodeURL(resp *http.Response) (*URL, error) {
	var decoded URL
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}