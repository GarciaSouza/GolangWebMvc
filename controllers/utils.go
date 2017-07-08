package controllers

import (
	"errors"
	"golang-webmvc/config"
	"golang-webmvc/models"
	"html/template"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"
)

//ViewResult A book view result
type ViewResult struct {
	Data    interface{}
	Errors  map[string][]error
	Session *models.Session
	User    *models.User
}

// Controller's helper functions

func view(res http.ResponseWriter, req *http.Request, tpladdr []string, data interface{}, errors []models.FieldError) {
	var tmpl *template.Template
	var err error

	if tmpl, err = template.New("").ParseGlob(path.Join("views", "*.gohtml")); err != nil {
		return500(res, err)
		return
	}

	for _, tpl := range tpladdr {
		if tmpl, err = tmpl.ParseFiles(tpl); err != nil {
			return500(res, err)
			return
		}
	}

	vr := getviewresult(req, data, errors)

	if err = tmpl.ExecuteTemplate(res, "master", vr); err != nil {
		return500(res, err)
	}
}

func return500(res http.ResponseWriter, err error) bool {
	if err != nil {
		if strings.Contains("dev,test", config.Env) {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(res, http.StatusText(500), http.StatusInternalServerError)
		}
		return true
	}

	return false
}

func getsession(req *http.Request) *models.Session {
	ssCookie, err := req.Cookie(config.SessionCookieName)
	if err != nil {
		return nil
	}

	session, err := models.OneSessionByKey(ssCookie.Value)
	if err != nil {
		return nil
	}

	if time.Now().Sub(session.LastActivity) > config.SessionTimeOut {
		ferr := models.DeleteSession(*session)
		if len(ferr) > 0 {
			//TODO: add log
		}

		return nil
	}

	session.LastActivity = time.Now()

	newsession, ferr := models.UpdateSession(*session)
	if len(ferr) > 0 {
		//TODO: add log
	} else {
		session = &newsession
	}

	return session
}

func getviewresult(req *http.Request, data interface{}, errors []models.FieldError) ViewResult {
	vr := ViewResult{
		Data:    nil,
		Errors:  make(map[string][]error),
		Session: nil,
		User:    nil,
	}

	if data != nil {
		vr.Data = data
	}

	if errors != nil {
		for _, b := range errors {
			if _, ok := vr.Errors[b.FieldName]; !ok {
				vr.Errors[b.FieldName] = []error{}
			}
			vr.Errors[b.FieldName] = append(vr.Errors[b.FieldName], b.Err)
		}
	}

	session := getsession(req)

	if session != nil {
		user, err := models.OneUserByID(session.UserID)
		if err != nil {
			return vr
		}

		vr.Session = session
		vr.User = user
	}

	return vr
}

func tplhome(tpls []string) []string {
	newtpls := []string{}
	for _, tpl := range tpls {
		newtpls = append(newtpls, path.Join("views", "home", tpl+".gohtml"))
	}
	return newtpls
}

func tplbooks(tpls []string) []string {
	newtpls := []string{}
	for _, tpl := range tpls {
		newtpls = append(newtpls, path.Join("views", "books", tpl+".gohtml"))
	}
	return newtpls
}

// Model's helper functions

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
