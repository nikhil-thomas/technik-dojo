package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Player defines a player
type Player struct {
	Name string
	Wins int
}

// PlayerStore interface design player behavior
type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(string)
	GetLeague() []Player
}

// PlayerServer manages a player server
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer creates a new player server
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))
	p.Handler = router
	return p
}

// ServeHTTP implements http handler interface
// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	p.router.ServeHTTP(w, r)
// }

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.WriteHeader(http.StatusOK)
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	leagueTable := p.store.GetLeague()
// 	return leagueTable
// }

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, player)
	case http.MethodPost:
		p.processWins(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "%d", score)
}

func (p *PlayerServer) processWins(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// // PlayerServer manages playersw
// func PlayerServer(w http.ResponseWriter, r *http.Request) {
//
// 	fmt.Fprintf(w, getPlayerScore(player))
// }

func getPlayerScore(p string) string {
	switch p {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	default:
		return ""
	}
}
