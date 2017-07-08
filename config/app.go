package config

import (
	"crypto/sha256"
	"hash"
)

//ApplicationSecretKey The key used to generate the secret
var ApplicationSecretKey string

//ApplicationSecretHash The hash algorithm used to generate the secret
var ApplicationSecretHash func() hash.Hash

func init() {
	ApplicationSecretKey = "mysecretkey"
	ApplicationSecretHash = sha256.New
}
