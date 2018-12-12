package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckStatusOK(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/status", nil)
	if err != nil {
		t.Error(err)
	}
	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(checkStatusOk)
	handler.ServeHTTP(resp, req)

	status := resp.Code
	if status != http.StatusOK {
		t.Errorf("expected %v, got %v\n", http.StatusOK, status)
	}

	expect := "Status: OK"

	if resp.Body.String() != expect {
		t.Errorf("expected %v, got %v\n", expect, resp.Body.String())
	}
}

func TestNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/notfound", nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp := httptest.NewRecorder()

	handler := http.HandlerFunc(statusNotFound)
	handler.ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusNotFound {
		t.Errorf("expected %v, got %v\n", http.StatusOK, status)
	}
}
