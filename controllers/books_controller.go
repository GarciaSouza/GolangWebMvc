package controllers

import (
	"golang-webmvc/models"
	"net/http"
	"path"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//BookIndex GET /books
func BookIndex(w http.ResponseWriter, req *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	tpladdr := []string{
		path.Join("views", "books", "index.gohtml"),
	}

	view(w, tpladdr, ViewResult{Data: bks})
}

//BookShow GET /books/:id
func BookShow(w http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 2 {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	view(w, tpladdr, ViewResult{Data: bk})
}

//BookNew GET /books/new
func BookNew(w http.ResponseWriter, req *http.Request) {
	tpladdr := []string{
		path.Join("views", "books", "new.gohtml"),
		path.Join("views", "books", "errors.gohtml"),
		path.Join("views", "books", "form.gohtml"),
	}

	view(w, tpladdr, ViewResult{Data: models.Book{}})
}

//BookCreate POST /books
func BookCreate(w http.ResponseWriter, req *http.Request) {
	var ferr []models.FieldError
	bk := models.Book{ID: bson.NewObjectId()}
	vr := ViewResult{Errors: make(map[string][]error)}

	if bk, ferr = parsebook(bk, req); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "new.gohtml"),
			path.Join("views", "books", "errors.gohtml"),
			path.Join("views", "books", "form.gohtml"),
		}
		vr.Data = bk
		feonmap(ferr, vr.Errors)
		view(w, tpladdr, vr) //ViewResult{Data: bk, Errors: ferr})
		return
	}

	if bk, ferr = models.PutBook(bk); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "new.gohtml"),
			path.Join("views", "books", "errors.gohtml"),
			path.Join("views", "books", "form.gohtml"),
		}
		vr.Data = bk
		feonmap(ferr, vr.Errors)
		view(w, tpladdr, vr) //ViewResult{Data: bk, Errors: ferr})
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	view(w, tpladdr, ViewResult{Data: bk})
}

//BookEdit GET /books/:id/edit
func BookEdit(w http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 3 || paths[2] != "edit" {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "edit.gohtml"),
		path.Join("views", "books", "errors.gohtml"),
		path.Join("views", "books", "form.gohtml"),
	}

	view(w, tpladdr, ViewResult{Data: bk})
}

//BookUpdate POST /books/:id
func BookUpdate(w http.ResponseWriter, req *http.Request) {
	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]
	if len(paths) != 2 {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	id := paths[1]
	if !bson.IsObjectIdHex(id) {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	bk, err := models.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	vr := ViewResult{Errors: make(map[string][]error)}
	var ferr []models.FieldError

	if bk, ferr = parsebook(bk, req); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "edit.gohtml"),
			path.Join("views", "books", "errors.gohtml"),
			path.Join("views", "books", "form.gohtml"),
		}
		vr.Data = bk
		feonmap(ferr, vr.Errors)
		view(w, tpladdr, vr)
		return
	}

	if bk, ferr = models.UpdateBook(bk); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "edit.gohtml"),
			path.Join("views", "books", "errors.gohtml"),
			path.Join("views", "books", "form.gohtml"),
		}
		vr.Data = bk
		feonmap(ferr, vr.Errors)
		view(w, tpladdr, vr)
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	view(w, tpladdr, ViewResult{Data: bk})
}

//BookDelete GET /books/:id/delete
func BookDelete(w http.ResponseWriter, req *http.Request) {
}

//BookDeleteConfirm POST /books/:id/delete
func BookDeleteConfirm(w http.ResponseWriter, req *http.Request) {
}
