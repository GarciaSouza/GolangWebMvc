package models

import (
	"golang-webmvc/config/db"
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

// Business

//NewSession Create a new session
func NewSession(key string, userid bson.ObjectId) Session {
	return Session{
		ID:           bson.NewObjectId(),
		Key:          key,
		UserID:       userid,
		LastActivity: time.Now(),
	}
}

//OneSessionByKey Find one session by Key
func OneSessionByKey(key string) (*Session, error) {
	return getSessionByKey(key)
}

//PutSession Insert a new session
func PutSession(session Session) (Session, []FieldError) {
	var err error
	fe := []FieldError{}

	if session, err = createNewSession(session); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return session, fe
}

//UpdateSession Update a existing session
func UpdateSession(session Session) (Session, []FieldError) {
	var err error
	fe := []FieldError{}

	if session, err = updateSession(session); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return session, fe
}

//DeleteSession Delete a existing session
func DeleteSession(session Session) []FieldError {
	fe := []FieldError{}

	if err := deleteSession(session); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return fe
}

// CRUD

func getSessionByKey(key string) (*Session, error) {
	var session *Session

	err := db.Sessions.Find(bson.M{"key": key}).One(&session)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func createNewSession(session Session) (Session, error) {
	err := db.Sessions.Insert(session)

	if err != nil {
		return session, err
	}

	return session, nil
}

func updateSession(session Session) (Session, error) {
	err := db.Sessions.Update(bson.M{"_id": session.ID}, &session)

	if err != nil {
		return session, err
	}

	return session, nil
}

func deleteSession(session Session) error {
	return db.Sessions.Remove(bson.M{"_id": session.ID})
}
