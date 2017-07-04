package main

import (
	"golang-webmvc/config"
	"golang-webmvc/controllers"
	"net/http"
	"strings"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/books", books)
	http.HandleFunc("/books/", booksID)
	http.ListenAndServe(":8080", nil)
}

func books(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)
	if req.Method == http.MethodGet {
		// GET /books
		controllers.BookIndex(res, req)
	} else if req.Method == http.MethodPost {
		// POST /books
		controllers.BookCreate(res, req)
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func booksID(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)

	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]

	if len(paths) < 2 && len(paths) > 3 {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}

	id := paths[1]

	if len(id) <= 0 {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
	}

	if req.Method == http.MethodGet {
		if len(paths) == 2 {
			if id == "new" {
				// GET /books/new
				controllers.BookNew(res, req)
			} else if bson.IsObjectIdHex(id) {
				// GET /books/:id
				controllers.BookShow(res, req)
			} else {
				http.Error(res, http.StatusText(400), http.StatusBadRequest)
			}
		} else if len(paths) == 3 {
			if bson.IsObjectIdHex(id) {
				action := paths[2]

				if action == "edit" {
					// GET /books/:id/edit
					controllers.BookEdit(res, req)
				} else if action == "delete" {
					// GET /books/:id/delete
					controllers.BookDelete(res, req)
				} else {
					http.Error(res, http.StatusText(400), http.StatusBadRequest)
				}
			} else {
				http.Error(res, http.StatusText(400), http.StatusBadRequest)
			}
		} else {
			http.Error(res, http.StatusText(404), http.StatusNotFound)
		}
	} else if req.Method == http.MethodPost {
		if bson.IsObjectIdHex(id) {
			if len(paths) == 2 {
				// POST /books/:id
				controllers.BookUpdate(res, req)
			} else if len(paths) == 3 && paths[2] == "delete" {
				// POST /books/:id/delete
				controllers.BookDeleteConfirm(res, req)
			} else {
				http.Error(res, http.StatusText(400), http.StatusBadRequest)
			}
		} else {
			http.Error(res, http.StatusText(400), http.StatusBadRequest)
		}
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func setSessionCookie(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		return
	}

	cookie, err := req.Cookie(config.SessionCookieName)

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     config.SessionCookieName,
			Value:    uuid.NewV4().String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(res, cookie)
	}
}
