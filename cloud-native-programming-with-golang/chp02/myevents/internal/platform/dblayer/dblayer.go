package dblayer

import (
	"errors"

	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/platform/dblayer/mongostore"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/store"
)

// DBTYPE defines supported database type
type DBTYPE string

const (
	// MONGODB represents mongodb
	MONGODB DBTYPE = "mongodb"

	// DYNAMODB represents dynamodb
	DYNAMODB DBTYPE = "dynamodb"
)

// NewStore creates a new store (db layer)
func NewStore(options DBTYPE, connAddr string) (store.EventReadWriter, error) {
	switch options {
	case MONGODB:
		return mongostore.NewStore(connAddr)
	}
	err := errors.New("error in creating new store")
	return nil, err
}
