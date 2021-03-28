package controllers

import (
	"finance/models"
	"net/http"
	"path"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//BookIndex GET /books
func BookIndex(res http.ResponseWriter, req *http.Request) {
	bks, err := models.AllBooks()

	if return500(res, err) {
		return
	}

	view(res, req, tplbooks([]string{"index"}), bks, nil)
}

//BookShow GET /books/:id
func BookShow(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 2 {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if return500(res, err) {
		return
	}

	view(res, req, tplbooks([]string{"show"}), bk, nil)
}

//BookNew GET /books/new
func BookNew(res http.ResponseWriter, req *http.Request) {
	if !isUserAuthorized(res, req, []string{"Admin"}) {
		return
	}

	view(res, req, tplbooks([]string{"new", "form"}), models.NewBook(), nil)
}

//BookCreate POST /books
func BookCreate(res http.ResponseWriter, req *http.Request) {
	var ferr []models.FieldError

	bk := models.NewBook()
	tpladdr := tplbooks([]string{"new", "form"})

	if bk, ferr = parsebook(bk, req); ferr != nil && len(ferr) > 0 {
		view(res, req, tpladdr, bk, ferr)
		return
	}

	if bk, ferr = models.PutBook(bk); ferr != nil && len(ferr) > 0 {
		view(res, req, tpladdr, bk, ferr)
		return
	}

	tpladdr = tplbooks([]string{"show"})

	view(res, req, tpladdr, bk, nil)
}

//BookEdit GET /books/:id/edit
func BookEdit(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 3 || paths[2] != "edit" {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if return500(res, err) {
		return
	}

	view(res, req, tplbooks([]string{"edit", "form"}), bk, nil)
}

//BookUpdate POST /books/:id
func BookUpdate(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 2 {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpladdr := tplbooks([]string{"edit", "form"})

	if newbk, ferr := parsebook(*bk, req); ferr != nil && len(ferr) > 0 {
		view(res, req, tpladdr, newbk, ferr)
		return
	}

	if newbk, ferr := models.UpdateBook(*bk); ferr != nil && len(ferr) > 0 {
		view(res, req, tpladdr, newbk, ferr)
		return
	}

	tpladdr = tplbooks([]string{"show"})

	view(res, req, tpladdr, bk, nil)
}

//BookDelete GET /books/:id/delete
func BookDelete(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 3 || paths[2] != "delete" {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if return500(res, err) {
		return
	}

	view(res, req, tplbooks([]string{"delete"}), bk, nil)
}

//BookDeleteConfirm POST /books/:id/delete
func BookDeleteConfirm(res http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 3 || paths[2] != "delete" {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	var ferr []models.FieldError

	if ferr = models.DeleteBook(*bk); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "delete.gohtml"),
		}
		view(res, req, tpladdr, bk, ferr)
		return
	}

	bks, err := models.AllBooks()
	if return500(res, err) {
		return
	}

	view(res, req, tplbooks([]string{"index"}), bks, nil)
}
