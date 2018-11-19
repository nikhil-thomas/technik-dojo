package main

// InMemoryPlayerStore implements Player store
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore returns score of a player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin returns score of a player
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetLeague returns score of a player
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

// NewInMemoryPlaystore initializes an empty playerstore
func NewInMemoryPlaystore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}
