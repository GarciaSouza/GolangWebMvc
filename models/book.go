package models

import (
	"golang-webmvc/config"

	"gopkg.in/mgo.v2/bson"
)

//Book A book
type Book struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Isbn   string        `json:"isbn" bson:"isbn"`
	Title  string        `json:"title" bson:"title"`
	Author string        `json:"author" bson:"author"`
	Price  float32       `json:"price" bson:"price"`
}

// Business

//AllBooks Get all books
func AllBooks() ([]Book, error) {
	return getAll()
}

//OneBookByIsbn Find one book by Isbn
func OneBookByIsbn(isbn string) (Book, error) {
	return getByIsbn(isbn)
}

//OneBookByID Find one book by ID
func OneBookByID(id bson.ObjectId) (Book, error) {
	return getByID(id)
}

// CRUD

func getAll() ([]Book, error) {
	bks := []Book{}
	err := config.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

func getByIsbn(isbn string) (Book, error) {
	bk := Book{}
	err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func getByID(id bson.ObjectId) (Book, error) {
	bk := Book{}
	err := config.Books.Find(bson.M{"_id": id}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func createNew(book Book) (Book, error) {
	err := config.Books.Insert(book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func update(book Book) (Book, error) {
	err := config.Books.Update(bson.M{"isbn": book.Isbn}, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func delete(book Book) error {
	return config.Books.Remove(bson.M{"isbn": book.Isbn})
}

// Validators

/*
func validateSaveBook(book Book) []error {
}

func validateEditBook(book Book) []error {
}

func validateRemoveBook(book Book) []error {
}
*/
