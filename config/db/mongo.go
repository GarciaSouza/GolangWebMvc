package db

import (
	"fmt"
	"golang-webmvc/config/log"
	"strings"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

var db *mgo.Database

//Host The host
var Host string

//DbName The name of the MongoDB database
var DbName string

//Books Books MongoDB Collection
var Books *mgo.Collection

//Users Users MongoDB Collection
var Users *mgo.Collection

//Sessions Sessions MongoDB Collection
var Sessions *mgo.Collection

//Open Try to stablish connection with MongoDB
func Open() {
	var err error

	session, err = mgo.Dial(fmt.Sprintf("mongodb://%s/%s", Host, DbName))
	if err != nil {
		log.Error.Fatalln(err)
	}

	if err = session.Ping(); err != nil {
		log.Error.Fatalln(err)
	}

	db = session.DB(DbName)
	Books = db.C("books")
	Users = db.C("users")
	Sessions = db.C("sessions")

	cols, err := db.CollectionNames()
	if err != nil {
		log.Error.Fatalln(err)
	}

	colsj := strings.Join(cols, ",")

	if !strings.Contains(colsj, "books") {
		if err = Books.Create(&mgo.CollectionInfo{}); err != nil {
			log.Error.Fatalln(err)
		}
	}

	if !strings.Contains(colsj, "users") {
		if err = Users.Create(&mgo.CollectionInfo{}); err != nil {
			log.Error.Fatalln(err)
		}
	}

	if !strings.Contains(colsj, "sessions") {
		if err = Sessions.Create(&mgo.CollectionInfo{}); err != nil {
			log.Error.Fatalln(err)
		}
	}

	if err = Users.EnsureIndexKey("username"); err != nil {
		log.Error.Fatalln(err)
	}

	if err = Sessions.EnsureIndexKey("key"); err != nil {
		log.Error.Fatalln(err)
	}

	if err = Sessions.EnsureIndexKey("userid"); err != nil {
		log.Error.Fatalln(err)
	}

	log.Info.Println("Connected to MongoDB")
}
