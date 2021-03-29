package books

import (
	"finance/controllers"
	"finance/models"
	modelBook "finance/models/book"
	"net/http"
	"path"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//BookIndex GET /books
func BookIndex(res http.ResponseWriter, req *http.Request) {
	bks, err := modelBook.AllBooks()

	if controllers.Return500(res, err) {
		return
	}

	controllers.View(res, req, controllers.TplBooks([]string{"index"}), bks, nil)
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

	bk, err := modelBook.OneBookByID(bson.ObjectIdHex(id))
	if controllers.Return500(res, err) {
		return
	}

	controllers.View(res, req, controllers.TplBooks([]string{"show"}), bk, nil)
}

//BookNew GET /books/new
func BookNew(res http.ResponseWriter, req *http.Request) {
	if !controllers.IsUserAuthorized(res, req, []string{"Admin"}) {
		return
	}

	controllers.View(res, req, controllers.TplBooks([]string{"new", "form"}), modelBook.NewBook(), nil)
}

//BookCreate POST /books
func BookCreate(res http.ResponseWriter, req *http.Request) {
	var ferr []models.FieldError

	bk := modelBook.NewBook()
	tpladdr := controllers.TplBooks([]string{"new", "form"})

	if bk, ferr = controllers.ParseBook(bk, req); ferr != nil && len(ferr) > 0 {
		controllers.View(res, req, tpladdr, bk, ferr)
		return
	}

	if bk, ferr = modelBook.PutBook(bk); ferr != nil && len(ferr) > 0 {
		controllers.View(res, req, tpladdr, bk, ferr)
		return
	}

	tpladdr = controllers.TplBooks([]string{"show"})

	controllers.View(res, req, tpladdr, bk, nil)
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

	bk, err := modelBook.OneBookByID(bson.ObjectIdHex(id))
	if controllers.Return500(res, err) {
		return
	}

	controllers.View(res, req, controllers.TplBooks([]string{"edit", "form"}), bk, nil)
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

	bk, err := modelBook.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	tpladdr := controllers.TplBooks([]string{"edit", "form"})

	if newbk, ferr := controllers.ParseBook(*bk, req); ferr != nil && len(ferr) > 0 {
		controllers.View(res, req, tpladdr, newbk, ferr)
		return
	}

	if newbk, ferr := modelBook.UpdateBook(*bk); ferr != nil && len(ferr) > 0 {
		controllers.View(res, req, tpladdr, newbk, ferr)
		return
	}

	tpladdr = controllers.TplBooks([]string{"show"})

	controllers.View(res, req, tpladdr, bk, nil)
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

	bk, err := modelBook.OneBookByID(bson.ObjectIdHex(id))
	if controllers.Return500(res, err) {
		return
	}

	controllers.View(res, req, controllers.TplBooks([]string{"delete"}), bk, nil)
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

	bk, err := modelBook.OneBookByID(bson.ObjectIdHex(id))
	if err != nil {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	var ferr []models.FieldError

	if ferr = modelBook.DeleteBook(*bk); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("controllers.views", "books", "delete.gohtml"),
		}
		controllers.View(res, req, tpladdr, bk, ferr)
		return
	}

	bks, err := modelBook.AllBooks()
	if controllers.Return500(res, err) {
		return
	}

	controllers.View(res, req, controllers.TplBooks([]string{"index"}), bks, nil)
}
