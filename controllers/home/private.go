package home

import (
	"finance/config"
	"finance/controllers"
	modelSession "finance/models/session"
	modelUser "finance/models/user"
	"net/http"
)

func dologin(res http.ResponseWriter, req *http.Request, user modelUser.User) {

	ssCookie, err := req.Cookie(config.SessionCookieName)
	if controllers.Return500(res, err) {
		return
	}

	newsession := modelSession.NewSession(ssCookie.Value, user.ID)
	_, ferr := modelSession.PutSession(newsession)
	if len(ferr) > 0 {
		controllers.Return500(res, ferr[0].Err)
		return
	}

	controllers.View(res, req, controllers.TplHome([]string{"index"}), nil, nil)
}
