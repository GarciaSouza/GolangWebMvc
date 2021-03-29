package controllers

import (
	"errors"
	"finance/config"
	"finance/models"
	modelBook "finance/models/book"
	modelSession "finance/models/session"
	modelUser "finance/models/user"
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
	Session *modelSession.Session
	User    *modelUser.User
}

//View View
func View(res http.ResponseWriter, req *http.Request, tpladdr []string, data interface{}, errors []models.FieldError) {
	var tmpl *template.Template
	var err error

	if tmpl, err = template.New("").ParseGlob(path.Join("views", "*.gohtml")); err != nil {
		Return500(res, err)
		return
	}

	for _, tpl := range tpladdr {
		if tmpl, err = tmpl.ParseFiles(tpl); err != nil {
			Return500(res, err)
			return
		}
	}

	vr := GetViewResult(req, data, errors)

	if err = tmpl.ExecuteTemplate(res, "master", vr); err != nil {
		Return500(res, err)
	}
}

//Return500 Return HTTP 500
func Return500(res http.ResponseWriter, err error) bool {
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

//Return401 Return401
func Return401(res http.ResponseWriter) {
	http.Error(res, http.StatusText(401), http.StatusUnauthorized)
}

//IsUserAuthorized IsUserAuthorized
func IsUserAuthorized(res http.ResponseWriter, req *http.Request, roles []string) bool {
	session := GetSession(req)
	if session == nil {
		Return401(res)
		return false
	}

	user, err := modelUser.OneUserByID(session.UserID)
	if err != nil {
		Return401(res)
		return false
	}

	ar := strings.Join(roles, ",")
	if !strings.Contains(ar, user.Role) {
		Return401(res)
		return false
	}

	return true
}

//GetSession GetSession
func GetSession(req *http.Request) *modelSession.Session {
	ssCookie, err := req.Cookie(config.SessionCookieName)
	if err != nil {
		return nil
	}

	session, err := modelSession.OneSessionByKey(ssCookie.Value)
	if err != nil {
		return nil
	}

	if time.Now().Sub(session.LastActivity) > config.SessionTimeOut {
		ferr := modelSession.DeleteSession(*session)
		if len(ferr) > 0 {
			//TODO: add log
		}

		return nil
	}

	session.LastActivity = time.Now()

	newsession, ferr := modelSession.UpdateSession(*session)
	if len(ferr) > 0 {
		//TODO: add log
	} else {
		session = &newsession
	}

	return session
}

//GetViewResult GetViewResult
func GetViewResult(req *http.Request, data interface{}, errors []models.FieldError) ViewResult {
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

	session := GetSession(req)

	if session != nil {
		user, err := modelUser.OneUserByID(session.UserID)
		if err != nil {
			return vr
		}

		vr.Session = session
		vr.User = user
	}

	return vr
}

//TplHome TplHome
func TplHome(tpls []string) []string {
	newtpls := []string{}
	for _, tpl := range tpls {
		newtpls = append(newtpls, path.Join("views", "home", tpl+".gohtml"))
	}
	return newtpls
}

//TplBooks TplBooks
func TplBooks(tpls []string) []string {
	newtpls := []string{}
	for _, tpl := range tpls {
		newtpls = append(newtpls, path.Join("views", "books", tpl+".gohtml"))
	}
	return newtpls
}

//ParseBook ParseBook
func ParseBook(bk modelBook.Book, req *http.Request) (modelBook.Book, []models.FieldError) {
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
