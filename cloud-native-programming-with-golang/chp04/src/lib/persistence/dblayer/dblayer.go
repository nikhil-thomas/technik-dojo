package dblayer

import (
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp04/src/lib/persistence"
	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp04/src/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB    DBTYPE = "mongodb"
	DOCUMENTDB DBTYPE = "documentdb"
	DYNAMODB   DBTYPE = "dynamodb"
)

func NewPresistencelayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
