// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "adder": operands TestHelpers
//
// Command:
// $ goagen
// --design=github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/01_adder/design
// --out=$(GOPATH)/src/github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/01_adder
// --version=v1.3.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/nikhil-thomas/technik-dojo/goa-exp/goa-examples/01_adder/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// AddOperandsOK runs the method Add of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func AddOperandsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.OperandsController, left int, right int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/add/%v/%v", left, right),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["left"] = []string{fmt.Sprintf("%v", left)}
	prms["right"] = []string{fmt.Sprintf("%v", right)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "OperandsTest"), rw, req, prms)
	addCtx, _err := app.NewAddOperandsContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Add(addCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
