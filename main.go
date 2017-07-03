package main

import (
	"golang-webmvc/controllers"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/books", books)
	http.HandleFunc("/books/", booksID)
	http.ListenAndServe(":8080", nil)
}

func books(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		// GET /books
		controllers.BookIndex(w, req)
	} else if req.Method == http.MethodPost {
		// POST /books
		controllers.BookCreate(w, req)
	} else {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
}

func booksID(w http.ResponseWriter, req *http.Request) {
	/*
		router.POST("/books/:id", controllers.BookUpdate)
		router.POST("/books/:id/delete", controllers.BookDeleteConfirm)
	*/
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]

	if len(paths) < 2 && len(paths) > 3 {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}

	id := paths[1]

	if len(id) <= 0 {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
	}

	if req.Method == http.MethodGet {
		if len(paths) == 2 {
			if id == "new" {
				// GET /books/new
				controllers.BookNew(w, req)
			} else if bson.IsObjectIdHex(id) {
				// GET /books/:id
				controllers.BookShow(w, req)
			} else {
				http.Error(w, http.StatusText(400), http.StatusBadRequest)
			}
		} else if len(paths) == 3 {
			if bson.IsObjectIdHex(id) {
				action := paths[2]

				if action == "edit" {
					// GET /books/:id/edit
					controllers.BookEdit(w, req)
				} else if action == "delete" {
					// GET /books/:id/delete
					controllers.BookDelete(w, req)
				} else {
					http.Error(w, http.StatusText(400), http.StatusBadRequest)
				}
			} else {
				http.Error(w, http.StatusText(400), http.StatusBadRequest)
			}
		} else {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
		}
	} else if req.Method == http.MethodPost {
		if bson.IsObjectIdHex(id) {
			if len(paths) == 2 {
				// POST /books/:id
				controllers.BookUpdate(w, req)
			} else if len(paths) == 3 && paths[2] == "delete" {
				// POST /books/:id/delete
				controllers.BookDeleteConfirm(w, req)
			} else {
				http.Error(w, http.StatusText(400), http.StatusBadRequest)
			}
		} else {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
		}
	} else {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
}
