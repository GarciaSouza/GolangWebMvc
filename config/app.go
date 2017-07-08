package config

//ApplicationSecretKey The secret key used to generate secret hash's
var ApplicationSecretKey string

func init() {
	ApplicationSecretKey = "mysecretkey"
}
