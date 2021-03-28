package models

import (
	"crypto/hmac"
	"errors"
	"finance/config"
	"finance/config/db"
	"fmt"
	"io"

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

// Business

//NewUser Create a new user
func NewUser() User {
	return User{
		ID: bson.NewObjectId(),
	}
}

//OneUserByID Find one user by ID
func OneUserByID(id bson.ObjectId) (*User, error) {
	return getUserByID(id)
}

//OneUserByUsername Find one user by Username
func OneUserByUsername(username string) (*User, error) {
	return getUserByUsername(username)
}

//PutUser Insert a new user
func PutUser(user User) (User, []FieldError) {
	var err error

	fe := validateSaveUser(user)
	if len(fe) > 0 {
		return user, fe
	}

	user, err = createNewUser(user)

	fe = []FieldError{}
	if err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return user, fe
}

//UpdateUser Update a existing user
func UpdateUser(user User) (User, []FieldError) {
	var err error

	fe := validateEditUser(user)
	if len(fe) > 0 {
		return user, fe
	}

	if user, err = updateUser(user); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return user, fe
}

//DeleteUser Delete a existing user
func DeleteUser(user User) []FieldError {
	fe := validateRemoveUser(user)

	if err := deleteUser(user); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return fe
}

//LoginValidate Validate the login
func LoginValidate(username string, password string) (*User, error) {
	user, err := OneUserByUsername(username)
	if err != nil {
		return nil, ErrorUserInvalidCredentials
	}

	secret := EncryptPass(password)

	if secret != user.Password {
		return nil, ErrorUserInvalidCredentials
	}

	return user, nil
}

//EncryptPass Encrypt the password
func EncryptPass(password string) string {
	h := hmac.New(config.ApplicationSecretHash, []byte(config.ApplicationSecretKey))
	io.WriteString(h, password)
	secret := fmt.Sprintf("%x", h.Sum(nil))
	return secret
}

// CRUD

func getUserByID(id bson.ObjectId) (*User, error) {
	var user *User

	err := db.Users.Find(bson.M{"_id": id}).One(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func getUserByUsername(username string) (*User, error) {
	var user *User

	err := db.Users.Find(bson.M{"username": username}).One(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func createNewUser(user User) (User, error) {
	err := db.Users.Insert(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func updateUser(user User) (User, error) {
	err := db.Users.Update(bson.M{"_id": user.ID}, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func deleteUser(user User) error {
	return db.Users.Remove(bson.M{"_id": user.ID})
}

// Validators

func validateSaveUser(user User) []FieldError {
	fe := []FieldError{}
	return fe
}

func validateEditUser(user User) []FieldError {
	fe := []FieldError{}
	return fe
}

func validateRemoveUser(user User) []FieldError {
	fe := []FieldError{}
	return fe
}
