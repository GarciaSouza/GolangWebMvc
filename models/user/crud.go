package user

import (
	"finance/config/db"

	"gopkg.in/mgo.v2/bson"
)

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
