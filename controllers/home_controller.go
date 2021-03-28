package controllers

import (
	"finance/config"
	"finance/models"
	modelSession "finance/models/session"
	modelUser "finance/models/user"
	"net/http"
	"time"
)

//HomeIndex GET /
func HomeIndex(res http.ResponseWriter, req *http.Request) {
	view(res, req, tplhome([]string{"index"}), nil, nil)
}

//HomeLogin GET /login
func HomeLogin(res http.ResponseWriter, req *http.Request) {
	view(res, req, tplhome([]string{"login"}), nil, nil)
}

//HomeLoginSubmit POST /login
func HomeLoginSubmit(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if return500(res, err) {
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	user, err := modelUser.LoginValidate(username, password)

	if err != nil {
		return500(res, err)
	} else if user != nil {
		dologin(res, req, *user)
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

	session, err := modelSession.OneSessionByKey(ssCookie.Value)
	if session != nil {
		ferr := modelSession.DeleteSession(*session)
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

	view(res, req, tplhome([]string{"index"}), nil, nil)
}

//HomeSignup GET /signup
func HomeSignup(res http.ResponseWriter, req *http.Request) {
	view(res, req, tplhome([]string{"signup"}), modelUser.NewUser(), nil)
}

//HomeSignupSubmit POST /signup
func HomeSignupSubmit(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if return500(res, err) {
		return
	}

	pass := req.FormValue("password")
	repass := req.FormValue("repassword")

	user := modelUser.NewUser()

	user.Username = req.FormValue("username")
	user.Firstname = req.FormValue("firstname")
	user.Lastname = req.FormValue("lastname")
	user.Email = req.FormValue("email")
	user.Role = "General"

	if pass != repass {
		ferr := []models.FieldError{
			{FieldName: "", Err: modelUser.ErrorUserPassRepass},
		}
		view(res, req, tplhome([]string{"signup"}), user, ferr)
		return
	}

	user.Password = modelUser.EncryptPass(req.FormValue("password"))

	user, ferr := modelUser.PutUser(user)
	if len(ferr) > 0 {
		view(res, req, tplhome([]string{"signup"}), user, ferr)
		return
	}

	dologin(res, req, user)
}

func dologin(res http.ResponseWriter, req *http.Request, user modelUser.User) {

	ssCookie, err := req.Cookie(config.SessionCookieName)
	if return500(res, err) {
		return
	}

	newsession := modelSession.NewSession(ssCookie.Value, user.ID)
	_, ferr := modelSession.PutSession(newsession)
	if len(ferr) > 0 {
		return500(res, ferr[0].Err)
		return
	}

	view(res, req, tplhome([]string{"index"}), nil, nil)
}
