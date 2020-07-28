package wrappers

import (
	"github.com/globalsign/mgo"
	"errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	if err := initMongoDB(); err != nil {
		log.Error("Init mongo err", err)
	}
}

const (
	defaultMongoAddr = "192.168.87.210:27017"
	defaultMongoDB   = "ptest"
)

var MongoDB *mgo.Database

func initMongoDB() error {
	dbInfo := &mgo.DialInfo{}

	addr := defaultMongoAddr
	//if v := ServerConfig.MongoAddr(); v != "" {
	//	addr = v
	//}
	dbInfo.Addrs = append(dbInfo.Addrs, addr)

	dbInfo.Database = defaultMongoDB
	//if v := ServerConfig.MongoDB(); v != "" {
	//	dbInfo.Database = v
	//}

	//if v := ServerConfig.MongoUser(); v != "" {
	//	dbInfo.Username = v
	//}
	//if v := ServerConfig.MongoPw(); v != "" {
	//	dbInfo.Password = v
	//}

	mongoSess, err := mgo.DialWithInfo(dbInfo)
	if err != nil {
		return err
	}

	MongoDB = mongoSess.DB(dbInfo.Database)
	if MongoDB == nil {
		return errors.New("no database")
	}

	return nil
}
