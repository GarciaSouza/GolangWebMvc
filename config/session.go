package config

import "time"

//SessionCookieName Name of the session cookie
var SessionCookieName string

//SessionTimeOut The maximum time to session
var SessionTimeOut time.Duration

func init() {
	SessionCookieName = "session"
	SessionTimeOut = time.Minute * 15
}
