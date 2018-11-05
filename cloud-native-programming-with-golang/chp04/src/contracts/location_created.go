package contracts

import "github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp04/src/lib/persistence"

type LocationCreatedEvent struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Address string             `json:"address"`
	Country string             `json:"country"`
	Halls   []persistence.Hall `json:"halls"`
}

func (c *LocationCreatedEvent) EventName() string {
	return "locationCreated"
}
