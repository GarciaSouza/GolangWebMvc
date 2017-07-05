package controllers

import (
	"golang-webmvc/models"
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

	tpladdr := []string{
		path.Join("views", "books", "index.gohtml"),
	}

	err = view(res, tpladdr, ViewResult{Data: bks})
	if return500(res, err) {
		return
	}
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

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	err = view(res, tpladdr, ViewResult{Data: bk})
	if return500(res, err) {
		return
	}
}

//BookNew GET /books/new
func BookNew(res http.ResponseWriter, req *http.Request) {
	tpladdr := []string{
		path.Join("views", "books", "new.gohtml"),
		path.Join("views", "books", "errors.gohtml"),
		path.Join("views", "books", "form.gohtml"),
	}

	err := view(res, tpladdr, ViewResult{Data: models.Book{}})
	if return500(res, err) {
		return
	}
}

//BookCreate POST /books
func BookCreate(res http.ResponseWriter, req *http.Request) {
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
		view(res, tpladdr, vr) //ViewResult{Data: bk, Errors: ferr})
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
		view(res, tpladdr, vr) //ViewResult{Data: bk, Errors: ferr})
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	err := view(res, tpladdr, ViewResult{Data: bk})
	if return500(res, err) {
		return
	}
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

	tpladdr := []string{
		path.Join("views", "books", "edit.gohtml"),
		path.Join("views", "books", "errors.gohtml"),
		path.Join("views", "books", "form.gohtml"),
	}

	err = view(res, tpladdr, ViewResult{Data: bk})
	if return500(res, err) {
		return
	}
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
		view(res, tpladdr, vr)
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
		view(res, tpladdr, vr)
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	err = view(res, tpladdr, ViewResult{Data: bk})
	if return500(res, err) {
		return
	}
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

	tpladdr := []string{
		path.Join("views", "books", "delete.gohtml"),
		path.Join("views", "books", "errors.gohtml"),
	}

	err = view(res, tpladdr, ViewResult{Data: bk})
	if return500(res, err) {
		return
	}
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

	vr := ViewResult{Errors: make(map[string][]error)}
	var ferr []models.FieldError

	if ferr = models.DeleteBook(bk); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "delete.gohtml"),
			path.Join("views", "books", "errors.gohtml"),
		}
		vr.Data = bk
		feonmap(ferr, vr.Errors)
		view(res, tpladdr, vr)
		return
	}

	bks, err := models.AllBooks()
	if return500(res, err) {
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "index.gohtml"),
	}

	err = view(res, tpladdr, ViewResult{Data: bks})
	if return500(res, err) {
		return
	}
}
