package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type URL struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	URL        string        `json:"url" bson:"url"`
	Alias      string        `json:"alias" bson:"alias"`
	Created_at time.Time     `json:"created_at" bson:"created_at"`
	visits     int           `json:"visits" bson:"visits"`
}
