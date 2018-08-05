package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Value     int           `json:"value" bson:"value"`
	Parameter string        `json:"parameter" bson:"parameter"`
}
