package models

import (
	"golang-webmvc/config/db"

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

//NewBook Create a new book
func NewBook() Book {
	return Book{
		ID: bson.NewObjectId(),
	}
}

//AllBooks Get all books
func AllBooks() ([]Book, error) {
	return getAllBook()
}

//OneBookByIsbn Find one book by Isbn
func OneBookByIsbn(isbn string) (*Book, error) {
	return getBookByIsbn(isbn)
}

//OneBookByID Find one book by ID
func OneBookByID(id bson.ObjectId) (*Book, error) {
	return getBookByID(id)
}

//PutBook Insert a new book
func PutBook(book Book) (Book, []FieldError) {
	var err error

	fe := validateSaveBook(book)
	if len(fe) > 0 {
		return book, fe
	}

	book, err = createNewBook(book)

	if err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return book, fe
}

//UpdateBook Update a existing book
func UpdateBook(book Book) (Book, []FieldError) {
	var err error

	fe := validateEditBook(book)
	if len(fe) > 0 {
		return book, fe
	}

	book, err = updateBook(book)
	if err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return book, fe
}

//DeleteBook Delete a existing book
func DeleteBook(book Book) []FieldError {
	fe := validateRemoveBook(book)

	if err := deleteBook(book); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return fe
}

// CRUD

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

// Validators

func validateSaveBook(book Book) []FieldError {
	fe := []FieldError{}
	//fe = append(fe, FieldError{FieldName: "Title", Err: errors.New("Choose a better Title")})
	return fe
}

func validateEditBook(book Book) []FieldError {
	fe := []FieldError{}
	return fe
}

func validateRemoveBook(book Book) []FieldError {
	fe := []FieldError{}
	return fe
}
