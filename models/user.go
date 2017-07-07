package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"golang-webmvc/config"
	"golang-webmvc/config/db"
	"io"

	"gopkg.in/mgo.v2/bson"
)

//User A user
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string        `bson:"username"`
	First    string        `bson:"first"`
	Last     string        `bson:"last"`
	Password string        `bson:"password"`
	Email    string        `bson:"email"`
	Role     string        `bson:"role"`
}

// Business

//OneUserByID Find one user by ID
func OneUserByID(id bson.ObjectId) (*User, error) {
	return getUserByID(id)
}

//OneUserByUsername Find one user by Username
func OneUserByUsername(username string) (*User, error) {
	return getUserByUsername(username)
}

//LoginValidate Validate the login
func LoginValidate(username string, password string) (*User, error) {
	user, err := OneUserByUsername(username)
	if err != nil {
		return nil, err
	}

	h := hmac.New(sha256.New, []byte(config.ApplicationSecretKey))
	io.WriteString(h, password)
	secret := fmt.Sprintf("%x", h.Sum(nil))

	if secret == user.Password {
		return user, nil
	}

	return nil, errors.New("Invalid user")
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

func getUserByUsername(username string) (*User, error) {
	var user *User
	err := db.Users.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
