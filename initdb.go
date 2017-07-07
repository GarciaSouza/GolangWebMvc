package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"golang-webmvc/config"
	"golang-webmvc/config/db"
	"golang-webmvc/models"
	"io"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	var err error
	db.Users.RemoveAll(bson.M{})
	db.Books.RemoveAll(bson.M{})
	db.Sessions.RemoveAll(bson.M{})

	adminPass := "admin#123"

	h := hmac.New(sha256.New, []byte(config.ApplicationSecretKey))
	_, err = io.WriteString(h, adminPass)
	if err != nil {
		log.Fatalln(err)
	}

	secret := fmt.Sprintf("%x", h.Sum(nil))

	admin := models.User{
		ID:       bson.NewObjectId(),
		Username: "admin",
		First:    "admin",
		Last:     "",
		Email:    "admin@localhost",
		Role:     "Admin",
		Password: secret,
	}

	err = db.Users.Insert(admin)
	if err != nil {
		log.Fatalln(err)
	}

	book1 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "8575420275",
		Title:  "O Poder do Agora",
		Author: "Tolle, Eckhart",
		Price:  20.30,
	}

	book2 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "9788539004119",
		Title:  "O Poder do Hábito - Por Que Fazemos o Que Fazemos na Vida e Nos Negócios",
		Author: "Duhigg, Charles",
		Price:  37,
	}

	book3 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "8575422391",
		Title:  "Os Segredos da Mente Milionária - Aprenda a Enriquecer Mudando seus Conceitos Sobre o Dinheiro",
		Author: "Eker, T. Harv",
		Price:  19.10,
	}

	book4 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "9788535206234",
		Title:  "Pai Rico Pai Pobre",
		Author: "Kiyosaki, Robert T. / Kiyosaki, Robert T.",
		Price:  48.90,
	}

	err = db.Books.Insert(book1)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Books.Insert(book2)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Books.Insert(book3)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Books.Insert(book4)
	if err != nil {
		log.Fatalln(err)
	}
}
