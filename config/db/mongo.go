package db

import (
	"golang-webmvc/config/log"
	"strings"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

//Books Books MongoDB Collection
var Books *mgo.Collection

//Users Users MongoDB Collection
var Users *mgo.Collection

//Sessions Sessions MongoDB Collection
var Sessions *mgo.Collection

func init() {
	Open()
}

//Open Try to stablish connection with MongoDB
func Open() {
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	db = s.DB("bookstore")
	Books = db.C("books")
	Users = db.C("users")
	Sessions = db.C("sessions")

	cols, err := db.CollectionNames()
	if err != nil {
		panic(err)
	}

	colsj := strings.Join(cols, ",")

	if !strings.Contains(colsj, "books") {
		if err = Books.Create(&mgo.CollectionInfo{}); err != nil {
			panic(err)
		}
	}

	if !strings.Contains(colsj, "users") {
		if err = Users.Create(&mgo.CollectionInfo{}); err != nil {
			panic(err)
		}
	}

	if !strings.Contains(colsj, "sessions") {
		if err = Sessions.Create(&mgo.CollectionInfo{}); err != nil {
			panic(err)
		}
	}

	if err = Users.EnsureIndexKey("username"); err != nil {
		panic(err)
	}

	if err = Sessions.EnsureIndexKey("key"); err != nil {
		panic(err)
	}

	if err = Sessions.EnsureIndexKey("userid"); err != nil {
		panic(err)
	}

	log.Info("You connected to your mongo database")
}
