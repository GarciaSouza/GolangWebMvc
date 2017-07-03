package main

import (
	"golang-webmvc/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	http.ListenAndServe("localhost:8080", makeRouter())
}

func makeRouter() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.StripPrefix("/public", http.FileServer(http.Dir("./public")))

	// Book routes
	router.GET("/books", controllers.BookIndex)
	router.GET("/books/:id", controllers.BookShow)
	/*router.GET("/books/new", controllers.BookNew)
	router.POST("/books", controllers.BookCreate)
	router.GET("/books/:id/edit", controllers.BookEdit)
	router.POST("/books/:id", controllers.BookUpdate)
	router.GET("/books/:id/delete", controllers.BookDelete)
	router.POST("/books/:id/delete", controllers.BookDeleteConfirm)*/

	return router
}
