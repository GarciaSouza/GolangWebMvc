package session

import (
	"finance/config/db"

	"gopkg.in/mgo.v2/bson"
)

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
