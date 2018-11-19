package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlaystore()
	server := NewPlayerServer(store)
	player := "Pepper"

	response := httptest.NewRecorder()

	postReq := newPostWinRequest(player)
	getReq := newGetScoreRequest(player)

	server.ServeHTTP(response, postReq)
	server.ServeHTTP(response, postReq)
	server.ServeHTTP(response, postReq)

	t.Run("get score", func(t *testing.T) {
		response = httptest.NewRecorder()
		server.ServeHTTP(response, getReq)
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response = httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

}
