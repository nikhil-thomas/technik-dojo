package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlaystore()
	server := PlayerServer{store}
	player := "Pepper"

	response := httptest.NewRecorder()

	postReq := newPostWinRequest(player)
	getReq := newGetScoreRequest(player)

	server.ServeHTTP(response, postReq)
	server.ServeHTTP(response, postReq)
	server.ServeHTTP(response, postReq)

	assertStatus(t, response.Code, http.StatusAccepted)

	response = httptest.NewRecorder()
	server.ServeHTTP(response, getReq)

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")

}
