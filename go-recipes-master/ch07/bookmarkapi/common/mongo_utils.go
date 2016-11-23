package common

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// GetSession returns a MongoDB Session
func getSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{AppConfig.MongoDBHost},
			Username: AppConfig.DBUser,
			Password: AppConfig.DBPwd,
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}
func createDBSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}

// Add indexes into MongoDB
func addIndexes() {
	var err error
	userIndex := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	// Add indexes into MongoDB
	session := getSession().Copy()
	defer session.Close()
	userCol := session.DB(AppConfig.Database).C("users")

	err = userCol.EnsureIndex(userIndex)
	if err != nil {
		log.Fatalf("[addIndexes]: %s\n", err)
	}
}

// DataStore for MongoDB
type DataStore struct {
	MongoSession *mgo.Session
}

// Close closes a mgo.Session value.
// Used to add defer statements for closing the copied session.
func (ds *DataStore) Close() {
	ds.MongoSession.Close()
}

// Collection returns mgo.collection for the given name
func (ds *DataStore) Collection(name string) *mgo.Collection {
	return ds.MongoSession.DB(AppConfig.Database).C(name)
}

// NewDataStore creates a new DataStore object to be used for each HTTP request.
func NewDataStore() *DataStore {
	session := getSession().Copy()
	dataStore := &DataStore{
		MongoSession: session,
	}
	return dataStore
}
