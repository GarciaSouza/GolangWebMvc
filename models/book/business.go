package book

import (
	"finance/models"

	"gopkg.in/mgo.v2/bson"
)

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
func PutBook(book Book) (Book, []models.FieldError) {
	var err error

	fe := validateSaveBook(book)
	if len(fe) > 0 {
		return book, fe
	}

	if book, err = createNewBook(book); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return book, fe
}

//UpdateBook Update a existing book
func UpdateBook(book Book) (Book, []models.FieldError) {
	var err error

	fe := validateEditBook(book)
	if len(fe) > 0 {
		return book, fe
	}

	if book, err = updateBook(book); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return book, fe
}

//DeleteBook Delete a existing book
func DeleteBook(book Book) []models.FieldError {
	fe := validateRemoveBook(book)

	if err := deleteBook(book); err != nil {
		fe = append(fe, models.FieldError{Err: err, FieldName: ""})
	}

	return fe
}
