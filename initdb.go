package main

import (
	"golang-webmvc/config"
	"golang-webmvc/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	config.Users.RemoveAll(bson.M{})
	config.Books.RemoveAll(bson.M{})
	config.Sessions.RemoveAll(bson.M{})

	adminPass := "admin#123"

	bs, err := bcrypt.GenerateFromPassword([]byte(adminPass), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}

	admin := models.User{
		Username: "admin",
		First:    "admin",
		Last:     "",
		Email:    "admin@localhost",
		Role:     "Admin",
		Password: bs,
	}

	err = config.Users.Insert(admin)
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

	err = config.Books.Insert(book1)
	if err != nil {
		log.Fatalln(err)
	}

	err = config.Books.Insert(book2)
	if err != nil {
		log.Fatalln(err)
	}
	err = config.Books.Insert(book3)
	if err != nil {
		log.Fatalln(err)
	}

	err = config.Books.Insert(book4)
	if err != nil {
		log.Fatalln(err)
	}
}
