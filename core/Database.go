package core

import (
	"fmt"
	config "github.com/fathoniadi/go-urlshort/config"
	"gopkg.in/mgo.v2"
)

func DBConnect() *mgo.Session {
	db_config := config.DBConfig{}
	db_config = db_config.Get()
	db, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("error connect")
	}

	return db
}
