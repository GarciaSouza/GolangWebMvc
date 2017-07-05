package db

import (
	"golang-webmvc/config/log"

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

	log.Info.Println("You connected to your mongo database")
}
