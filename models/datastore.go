package models

import "gopkg.in/mgo.v2"

// TODO: config values
const (
	Host     = "mongodb://3.3.3.3:30017"
	Username = "root"
	Password = ""
	Database = "kaupa"
)

// DataStore manages mongo connection
type DataStore struct {
	session *mgo.Session
}

// NewDataStore creates a new DataStore
func NewDataStore() *DataStore {
	return &DataStore{session: getSession()}
}

func getSession() *mgo.Session {
	// Connect to our local mongo
	session, err := mgo.Dial(Host) // TODO: config values

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return session
}