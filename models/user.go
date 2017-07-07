package models

import (
	"golang-webmvc/config/db"

	"gopkg.in/mgo.v2/bson"
)

//User A user
type User struct {
	Username string
	First    string
	Last     string
	Password []byte
	Email    string
	Role     string
}

// Business

//OneUserByID Find one user by ID
func OneUserByID(id bson.ObjectId) (*User, error) {
	return getUserByID(id)
}

// CRUD

func getUserByID(id bson.ObjectId) (*User, error) {
	var user *User
	err := db.Users.Find(bson.M{"_id": id}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
