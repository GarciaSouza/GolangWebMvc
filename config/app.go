package config

import (
	"crypto/sha256"
	"finance/config/db"
	applog "finance/config/log"
	"hash"
	"io/ioutil"
	"log"
	"os"
	"time"
)

//Env Environment
var Env string

//Port Port
var Port string

//ApplicationSecretKey The key used to generate the secret
var ApplicationSecretKey string

//ApplicationSecretHash The hash algorithm used to generate the secret
var ApplicationSecretHash func() hash.Hash

func init() {
	ApplicationSecretKey = "mysecretkey"
	ApplicationSecretHash = sha256.New

	SessionCookieName = "session"
	SessionTimeOut = time.Minute * 15

	logflags := log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile

	applog.Trace = log.New(ioutil.Discard, "TRACE: ", logflags)
	applog.Debug = log.New(os.Stdout, "DEBUG: ", logflags)
	applog.Info = log.New(os.Stdout, "INFO: ", logflags)
	applog.Warning = log.New(os.Stdout, "WARNING: ", logflags)
	applog.Error = log.New(os.Stderr, "ERROR: ", logflags)

	db.Host = "localhost"
	db.DbName = "bookstore"
	db.Open()
}
