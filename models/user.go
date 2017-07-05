package models

//User A user
type User struct {
	Username string
	First    string
	Last     string
	Password []byte
	Email    string
	Role     string
}
