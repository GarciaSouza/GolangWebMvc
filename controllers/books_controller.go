package controllers

import (
	"fmt"
	"golang-webmvc/models"
	"html/template"
	"io"
	"net/http"
	"path"

	"gopkg.in/mgo.v2/bson"

	"github.com/julienschmidt/httprouter"
)

//BookIndex GET /books
func BookIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	bks, err := models.AllBooks()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	ct := path.Join("views", "books", "index.gohtml")
	view(w, ct, bks)
}

//BookShow GET /books/:id
func BookShow(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	bk, err := models.OneBookByID(bson.ObjectIdHex(params.ByName("id")))
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}

	ct := path.Join("views", "books", "show.gohtml")
	view(w, ct, bk)
}

//BookNew GET /books/new
func BookNew(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

//BookCreate POST /books
func BookCreate(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

//BookEdit GET /books/:id/edit
func BookEdit(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

//BookUpdate POST /books/:id
func BookUpdate(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

//BookDelete GET /books/:id/delete
func BookDelete(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

//BookDeleteConfirm POST /books/:id/delete
func BookDeleteConfirm(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
}

// Mapper req to model

func view(w io.Writer, tpladdr string, data interface{}) error {
	var tmpl *template.Template
	var err error

	var fm = template.FuncMap{
		"hexstr": hexstr,
	}

	if tmpl, err = template.New("").Funcs(fm).ParseGlob(path.Join("views", "*.gohtml")); err != nil {
		return err
	}

	if tmpl, err = tmpl.ParseFiles(tpladdr); err != nil {
		return err
	}

	if err = tmpl.ExecuteTemplate(w, "master", data); err != nil {
		return err
	}

	return nil
}

func hexstr(id bson.ObjectId) string {
	return id.Hex()
}
