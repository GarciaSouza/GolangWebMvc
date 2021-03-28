package user

import (
	"crypto/hmac"
	"finance/config"
	"finance/models"
	"fmt"
	"io"

	"gopkg.in/mgo.v2/bson"
)

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
func PutUser(user User) (User, []models.FieldError) {
	var err error

	fe := validateSaveUser(user)
	if len(fe) > 0 {
		return user, fe
	}

	user, err = createNewUser(user)

	fe = []models.FieldError{}
	if err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return user, fe
}

//UpdateUser Update a existing user
func UpdateUser(user User) (User, []models.FieldError) {
	var err error

	fe := validateEditUser(user)
	if len(fe) > 0 {
		return user, fe
	}

	if user, err = updateUser(user); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return user, fe
}

//DeleteUser Delete a existing user
func DeleteUser(user User) []models.FieldError {
	fe := validateRemoveUser(user)

	if err := deleteUser(user); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
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
