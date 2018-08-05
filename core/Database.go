package core

import (
	config "github.com/fathoniadi/go-urlshort/config"
	"gopkg.in/mgo.v2"
	"strconv"
	"time"
)

func DBConnect() *mgo.Session {
	db_config := config.Load()

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{db_config.Database_host + ":" + strconv.Itoa(db_config.Database_port)},
		Timeout:  60 * time.Second,
		Database: db_config.Database_name,
		Username: db_config.Database_username,
		Password: db_config.Database_password,
	}

	db, err := mgo.DialWithInfo(mongoDBDialInfo)

	if err != nil {
		panic(err)
	}

	return db
}
