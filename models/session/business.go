package session

import (
	"finance/models"
	"time"

	"gopkg.in/mgo.v2/bson"
)

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
func PutSession(session Session) (Session, []models.FieldError) {
	var err error
	fe := []models.FieldError{}

	if session, err = createNewSession(session); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return session, fe
}

//UpdateSession Update a existing session
func UpdateSession(session Session) (Session, []models.FieldError) {
	var err error
	fe := []models.FieldError{}

	if session, err = updateSession(session); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return session, fe
}

//DeleteSession Delete a existing session
func DeleteSession(session Session) []models.FieldError {
	fe := []models.FieldError{}

	if err := deleteSession(session); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return fe
}
