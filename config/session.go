package config

import "time"

//SessionCookieName Name of the session cookie
var SessionCookieName string

//SessionTimeOut The maximum time to session
var SessionTimeOut time.Duration

//ApplicationSecretKey The secret key used to generate secret hash's
var ApplicationSecretKey string

func init() {
	SessionCookieName = "session"
	SessionTimeOut = time.Minute * 15
	ApplicationSecretKey = "mysecretkey"
}
