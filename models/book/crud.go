package book

import (
	"finance/config/db"

	"gopkg.in/mgo.v2/bson"
)

func getAllBook() ([]Book, error) {
	bks := []Book{}
	err := db.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}

	return bks, nil
}

func getBookByIsbn(isbn string) (*Book, error) {
	var bk *Book
	err := db.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func getBookByID(id bson.ObjectId) (*Book, error) {
	var bk *Book
	err := db.Books.Find(bson.M{"_id": id}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func createNewBook(book Book) (Book, error) {
	err := db.Books.Insert(book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func updateBook(book Book) (Book, error) {
	err := db.Books.Update(bson.M{"_id": book.ID}, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func deleteBook(book Book) error {
	return db.Books.Remove(bson.M{"_id": book.ID})
}
