package controllers

import (
	"golang-webmvc/config"
	"golang-webmvc/models"
	"net/http"
	"time"
)

//HomeIndex GET /
func HomeIndex(res http.ResponseWriter, req *http.Request) {
	err := view(res, req, tplhome([]string{"index"}), nil, nil)
	return500(res, err)
}

//HomeLogin GET /login
func HomeLogin(res http.ResponseWriter, req *http.Request) {
	err := view(res, req, tplhome([]string{"login"}), nil, nil)
	return500(res, err)
}

//HomeLoginSubmit POST /login
func HomeLoginSubmit(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if return500(res, err) {
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	user, err := models.LoginValidate(username, password)

	if err != nil {
		return500(res, err)
	} else if user != nil {
		ssCookie, err := req.Cookie(config.SessionCookieName)
		if return500(res, err) {
			return
		}

		newsession := models.NewSession(ssCookie.Value, user.ID)
		_, ferr := models.PutSession(newsession)
		if len(ferr) > 0 {
			return500(res, ferr[0].Err)
			return
		}

		err = view(res, req, tplhome([]string{"index"}), nil, nil)
		return500(res, err)
	} else {
		http.Error(res, http.StatusText(500), http.StatusInternalServerError)
	}
}

//HomeLogout GET /logout
func HomeLogout(res http.ResponseWriter, req *http.Request) {
	ssCookie, err := req.Cookie(config.SessionCookieName)
	if return500(res, err) {
		return
	}

	session, err := models.OneSessionByKey(ssCookie.Value)
	if session != nil {
		ferr := models.DeleteSession(*session)
		if len(ferr) > 0 {
			return500(res, ferr[0].Err)
			return
		}

		cookie, err := req.Cookie(config.SessionCookieName)
		if return500(res, err) {
			return
		}

		cookie.Expires = time.Date(1970, 1, 1, 9, 9, 9, 9, time.UTC)
		cookie.Value = ""

		http.SetCookie(res, cookie)
	}

	err = view(res, req, tplhome([]string{"index"}), nil, nil)
	return500(res, err)
}

//HomeSignup GET /signup
func HomeSignup(res http.ResponseWriter, req *http.Request) {
	err := view(res, req, tplhome([]string{"signup"}), nil, nil)
	return500(res, err)
}

//HomeSignupSubmit POST /signup
func HomeSignupSubmit(res http.ResponseWriter, req *http.Request) {
}
