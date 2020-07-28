package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	mgo "gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"runtime"
)

func main() {
	dbDriver := "mgo"
	dbConnectionString := "host=192.168.87.210"

	// init database with appropriate driver
	db, err := InitDB(dbDriver, dbConnectionString, runtime.NumCPU()*4)
	if err != nil {
		log.Fatal(err)
	}

	db.GetOneRandomWorld()

	defer db.Close()
}

// Mongo struct
type Mongo struct {
	session *mgo.Session
	// collections
	database *mgo.Database
	fortunes *mgo.Collection
	worlds   *mgo.Collection
}

// Connect create connection and ping db
func (mongo *Mongo) Connect(dbConnectionString string, maxConnectionsInPool int) error {
	session, err := mgo.Dial(dbConnectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	mongo.database = session.DB("hello_world")
	mongo.worlds = mongo.database.C("world")
	mongo.fortunes = mongo.database.C("fortune")

	return nil
}

// Close connect to db
func (mongo *Mongo) Close() {
	mongo.session.Close()
}

func (mongo Mongo) GetOneRandomWorld() error {
	var err error
	err = nil
	//queryID := rand.Intn(worldsCount) + 1
	//if err = mongo.worlds.Find(bson.M{"_id": queryID}).One(&w); err != nil {
	//	err = fmt.Errorf("error finding world with ID %d: %s", queryID, err)
	//}
	return err
}

// NewMongoDB creates new connection to mongodb with mgo driver
func NewMongoDB(dbConnectionString string, maxConnectionsInPool int) (*Mongo, error) {
	var mongo Mongo
	if err := mongo.Connect(dbConnectionString, maxConnectionsInPool); err != nil {
		return nil, err
	}
	return &mongo, nil
}

// InitDB with appropriate driver
func InitDB(dbDriver, dbConnectionString string, maxConnectionCount int) (*Mongo, error) {
	var err error
	var db *Mongo

	db, err = NewMongoDB(dbConnectionString, maxConnectionCount)
	if err != nil {
		return nil, fmt.Errorf("Error opening mongo database with mgo driver: %s", err)
	}
	// } else if dbDriver == "pq" {
	// 	db, err = NewPqDB(
	// 		dbConnectionString, maxConnectionCount)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("Error opening postgresql database with pq driver: %s", err)
	// 	}

	return db, nil
}
