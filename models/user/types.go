package user

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

//User A user
type User struct {
	ID        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"username"`
	Firstname string        `bson:"firstname"`
	Lastname  string        `bson:"lastname"`
	Password  string        `bson:"password"`
	Email     string        `bson:"email"`
	Role      string        `bson:"role"`
}

var (
	//ErrorUserInvalidCredentials Invalid user credentials error
	ErrorUserInvalidCredentials = errors.New("Invalid credentials")
	//ErrorUserPassRepass Password and Repassword are different
	ErrorUserPassRepass = errors.New("Password and Repassword are different")
)
