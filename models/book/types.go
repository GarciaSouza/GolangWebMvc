package book

import (
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
