package controllers

import (
	"errors"
	"golang-webmvc/config/log"
	"golang-webmvc/models"
	"html/template"
	"io"
	"net/http"
	"path"
	"strconv"
)

//ViewResult A book view result
type ViewResult struct {
	Errors map[string][]error
	Data   interface{}
}

// Controller's functions

func view(w io.Writer, tpladdr []string, data interface{}) error {
	var tmpl *template.Template
	var err error

	if tmpl, err = template.New("").ParseGlob(path.Join("views", "*.gohtml")); err != nil {
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

func feonmap(afielderror []models.FieldError, mapOf map[string][]error) {
	if afielderror == nil || mapOf == nil {
		return
	}

	for _, b := range afielderror {
		if _, ok := mapOf[b.FieldName]; !ok {
			mapOf[b.FieldName] = []error{}
		}
		mapOf[b.FieldName] = append(mapOf[b.FieldName], b.Err)
	}
}

func parsebook(bk models.Book, req *http.Request) (models.Book, []models.FieldError) {
	ferr := []models.FieldError{}

	req.ParseForm()

	bk.Isbn = req.FormValue("Isbn")
	bk.Title = req.FormValue("Title")
	bk.Author = req.FormValue("Author")

	p := req.FormValue("Price")
	f64, err := strconv.ParseFloat(p, 64)
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

func return500(res http.ResponseWriter, err error) bool {
	if err != nil {
		log.Error.Println(err)
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		return true
	}

	return false
}
