package store

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// User struct defines a user
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	First    string
	Last     string
	Age      int
	Bookings []Booking
}

// String prints user details as a string
func (u *User) String() string {
	return fmt.Sprintf(
		"id: %s\nFirst: %s\nLast: %s\nAge: %d\nBookings: %v\n",
		u.ID,
		u.First,
		u.Last,
		u.Age,
		u.Bookings,
	)
}

// Booking struct describes a booking
type Booking struct {
	Date    int64
	EventID []byte
	Seats   int
}

// Event struct describes an event
type Event struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Duration  int
	StartDate int64
	EndDate   int64
	Location  Location
}

// Location struct describes a location
type Location struct {
	ID        bson.ObjectId `bson:"_id"`
	Name      string
	Address   string
	Country   string
	OpenTime  int
	CloseTime int
	Halls     []Hall
}

// Hall struct describes a Hall
type Hall struct {
	Name     string `json:"name"`
	Location string `json:"location,omitempty"`
	Capacity int    `json:"capacity"`
}
