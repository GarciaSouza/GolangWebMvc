package controllers

import (
	"errors"
	"fmt"
	"golang-webmvc/models"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//BookResult A book view result
type BookResult struct {
	Errors []models.FieldError
	One    models.Book
	Data   []models.Book
}

//BookIndex GET /books
func BookIndex(w http.ResponseWriter, req *http.Request) {
	bks, err := models.AllBooks()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	tpladdr := []string{
		path.Join("views", "books", "index.gohtml"),
	}

	view(w, tpladdr, bks)
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

	view(w, tpladdr, bk)
}

//BookNew GET /books/new
func BookNew(w http.ResponseWriter, req *http.Request) {
}

//BookCreate POST /books
func BookCreate(w http.ResponseWriter, req *http.Request) {
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
		path.Join("views", "books", "form.gohtml"),
	}

	view(w, tpladdr, BookResult{One: bk})
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

	var ferr []models.FieldError

	if bk, ferr = parse(bk, req); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "edit.gohtml"),
			path.Join("views", "books", "form.gohtml"),
		}
		view(w, tpladdr, BookResult{One: bk, Errors: ferr})
		return
	}

	if bk, ferr = models.UpdateBook(bk); ferr != nil && len(ferr) > 0 {
		tpladdr := []string{
			path.Join("views", "books", "edit.gohtml"),
			path.Join("views", "books", "form.gohtml"),
		}
		view(w, tpladdr, BookResult{One: bk, Errors: ferr})
		return
	}

	tpladdr := []string{
		path.Join("views", "books", "show.gohtml"),
	}

	view(w, tpladdr, bk)
}

//BookDelete GET /books/:id/delete
func BookDelete(w http.ResponseWriter, req *http.Request) {
}

//BookDeleteConfirm POST /books/:id/delete
func BookDeleteConfirm(w http.ResponseWriter, req *http.Request) {
}

// Mapper req to model

func view(w io.Writer, tpladdr []string, data interface{}) error {
	var tmpl *template.Template
	var err error

	var fm = template.FuncMap{
		"hexstr": hexstr,
	}

	if tmpl, err = template.New("").Funcs(fm).ParseGlob(path.Join("views", "*.gohtml")); err != nil {
		return err
	}

	for _, tpl := range tpladdr {
		if tmpl, err = tmpl.ParseFiles(tpl); err != nil {
			return err
		}
	}

	if err = tmpl.ExecuteTemplate(w, "master", data); err != nil {
		return err
	}

	return nil
}

func hexstr(id bson.ObjectId) string {
	return id.Hex()
}

func parse(bk models.Book, req *http.Request) (models.Book, []models.FieldError) {
	ferr := []models.FieldError{}

	req.ParseForm()

	bk.Isbn = req.FormValue("Isbn")
	bk.Title = req.FormValue("Title")
	bk.Author = req.FormValue("Author")

	p := req.FormValue("Price")
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		f := models.FieldError{
			FieldName: "Price",
			Err:       errors.New("Must be a number"),
		}
		ferr = append(ferr, f)
	} else {
		bk.Price = float32(f64)
	}

	return bk, ferr
}
