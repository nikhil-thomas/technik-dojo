package main

import (
	"fmt"
	"net/http"
)

// PlyerStore interface design player behavior
type PlyerStore interface {
	GetPlayerScore(string) int
	RecordWin(string)
}

// PlayerServer manages a player server
type PlayerServer struct {
	store PlyerStore
}

// ServeHTTP implements http handler interface
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
