package core

import (
	"fmt"
	config "github.com/fathoniadi/go-urlshort/config"
	"gopkg.in/mgo.v2"
)

func DBConnect() *mgo.Session {
	db_config := config.Load()
	fmt.Println(db_config)
	db, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("error connect")
	}

	return db
}
