package models

import "time"

//Session A session
type Session struct {
	Key          string
	LastActivity time.Time
}
