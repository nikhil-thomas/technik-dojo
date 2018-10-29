package mongostore

import (
	"log"

	"github.com/nikhil-thomas/technik-dojo/cloud-native-programming-with-golang/chp02/myevents/internal/store"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	// DB constant represents db name
	DB = "myevents"

	// USERS constant represents db users collection
	USERS = "users"

	// EVENTS constant represents db events collection
	EVENTS = "events"
)

// MongoDBStore describes a mongodb store
type MongoDBStore struct {
	session *mgo.Session
}

// NewStore creates a new mongodb store
func NewStore(connAddr string) (store.EventReadWriter, error) {
	s, err := mgo.Dial(connAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &MongoDBStore{
		session: s,
	}, nil
}

func (mgoStr *MongoDBStore) getSession() *mgo.Session {
	return mgoStr.session.Copy()
}

// AddEvent add an event to mongoDB
func (mgoStr *MongoDBStore) AddEvent(evt store.Event) ([]byte, error) {
	sess := mgoStr.getSession()
	defer sess.Close()

	if !evt.ID.Valid() {
		evt.ID = bson.NewObjectId()
	}

	if !evt.Location.ID.Valid() {
		evt.Location.ID = bson.NewObjectId()
	}

	err := sess.DB(DB).C(EVENTS).Insert(evt)
	return []byte(evt.ID.String()), err
}

// FindEvent finds an event with id
func (mgoStr *MongoDBStore) FindEvent(id []byte) (store.Event, error) {
	sess := mgoStr.getSession()
	defer sess.Close()

	evt := store.Event{}

	err := sess.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&evt)
	return evt, err
}

// FindEventByName finds an event with name
func (mgoStr *MongoDBStore) FindEventByName(name string) (store.Event, error) {
	sess := mgoStr.getSession()
	defer sess.Close()

	evt := store.Event{}

	err := sess.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&evt)
	return evt, err
}

// FindAllEvents returns all events
func (mgoStr *MongoDBStore) FindAllEvents() ([]store.Event, error) {
	sess := mgoStr.getSession()
	defer sess.Close()

	evts := []store.Event{}

	err := sess.DB(DB).C(EVENTS).Find(nil).All(&evts)
	return evts, err
}
