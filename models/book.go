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

//PutBook Insert a new book
func PutBook(book Book) (Book, []FieldError) {
	var err error

	fe := validateSaveBook(book)
	if len(fe) > 0 {
		return book, fe
	}

	book, err = createNew(book)

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

	book, err = update(book)
	if err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return book, fe
}

//DeleteBook Delete a existing book
func DeleteBook(book Book) []FieldError {
	fe := validateRemoveBook(book)

	if err := delete(book); err != nil {
		fe = append(fe, FieldError{Err: err, FieldName: ""})
	}

	return fe
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
	err := config.Books.Update(bson.M{"_id": book.ID}, &book)
	if err != nil {
		return book, err
	}
	return book, nil
}

func delete(book Book) error {
	return config.Books.Remove(bson.M{"_id": book.ID})
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
