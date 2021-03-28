package session

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Session A session
type Session struct {
	ID           bson.ObjectId `bson:"_id"`
	Key          string        `bson:"key"`
	UserID       bson.ObjectId `bson:"userid"`
	LastActivity time.Time     `bson:"lastactivity"`
}
