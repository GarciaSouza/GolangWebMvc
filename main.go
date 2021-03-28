package main

import (
	"finance/config"
	"finance/config/db"
	"finance/config/log"
	"finance/controllers"
	"finance/models"

	"flag"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	config.Port = *flag.String("port", "8182", "port")
	config.Env = *flag.String("env", "dev", "environment")

	flag.Parse()
	tail := flag.Args()

	if len(tail) > 0 {
		if len(tail) == 1 && tail[0] == "initdb" {
			initdb()
			return
		}

		log.Error.Fatalln("Invalid Arguments")
	}

	log.Info.Println("Running On", config.Port)

	// Serve the public directory statically
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/books", books)
	http.HandleFunc("/books/", booksID)
	http.HandleFunc("/", home)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signup)

	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Error.Fatalln(err)
	}
}

func initdb() {
	var err error

	log.Info.Println("InitDB - Start")

	_, err = db.Users.RemoveAll(bson.M{})
	if err != nil {
		log.Error.Fatalln(err)
	}

	_, err = db.Books.RemoveAll(bson.M{})
	if err != nil {
		log.Error.Fatalln(err)
	}

	_, err = db.Sessions.RemoveAll(bson.M{})
	if err != nil {
		log.Error.Fatalln(err)
	}

	admin := models.User{
		ID:        bson.NewObjectId(),
		Username:  "admin",
		Firstname: "admin",
		Lastname:  "",
		Email:     "admin@localhost",
		Role:      "Admin",
		Password:  models.EncryptPass("admin#123"),
	}

	err = db.Users.Insert(admin)
	if err != nil {
		log.Error.Fatalln(err)
	}

	book1 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "8575420275",
		Title:  "O Poder do Agora",
		Author: "Tolle, Eckhart",
		Price:  20.30,
	}

	book2 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "9788539004119",
		Title:  "O Poder do Hábito - Por Que Fazemos o Que Fazemos na Vida e Nos Negócios",
		Author: "Duhigg, Charles",
		Price:  37,
	}

	book3 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "8575422391",
		Title:  "Os Segredos da Mente Milionária - Aprenda a Enriquecer Mudando seus Conceitos Sobre o Dinheiro",
		Author: "Eker, T. Harv",
		Price:  19.10,
	}

	book4 := models.Book{
		ID:     bson.NewObjectId(),
		Isbn:   "9788535206234",
		Title:  "Pai Rico Pai Pobre",
		Author: "Kiyosaki, Robert T. / Kiyosaki, Robert T.",
		Price:  48.90,
	}

	err = db.Books.Insert(book1)
	if err != nil {
		log.Error.Fatalln(err)
	}

	err = db.Books.Insert(book2)
	if err != nil {
		log.Error.Fatalln(err)
	}
	err = db.Books.Insert(book3)
	if err != nil {
		log.Error.Fatalln(err)
	}

	err = db.Books.Insert(book4)
	if err != nil {
		log.Error.Fatalln(err)
	}

	log.Info.Println("InitDB - Done")
}

func books(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)
	if req.Method == http.MethodGet {
		// GET /books
		controllers.BookIndex(res, req)
	} else if req.Method == http.MethodPost {
		// POST /books
		controllers.BookCreate(res, req)
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func booksID(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)

	paths := strings.Split(req.URL.Path, "/")
	paths = paths[1:]

	if len(paths) < 2 && len(paths) > 3 {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}

	id := paths[1]

	if len(id) <= 0 {
		http.Error(res, http.StatusText(400), http.StatusBadRequest)
	}

	if req.Method == http.MethodGet {
		if len(paths) == 2 {
			if id == "new" {
				// GET /books/new
				controllers.BookNew(res, req)
			} else if bson.IsObjectIdHex(id) {
				// GET /books/:id
				controllers.BookShow(res, req)
			} else {
				http.Error(res, http.StatusText(400), http.StatusBadRequest)
			}
		} else if len(paths) == 3 {
			if bson.IsObjectIdHex(id) {
				action := paths[2]

				if action == "edit" {
					// GET /books/:id/edit
					controllers.BookEdit(res, req)
				} else if action == "delete" {
					// GET /books/:id/delete
					controllers.BookDelete(res, req)
				} else {
					http.Error(res, http.StatusText(400), http.StatusBadRequest)
				}
			} else {
				http.Error(res, http.StatusText(400), http.StatusBadRequest)
			}
		} else {
			http.Error(res, http.StatusText(404), http.StatusNotFound)
		}
	} else if req.Method == http.MethodPost {
		if bson.IsObjectIdHex(id) {
			if len(paths) == 2 {
				// POST /books/:id
				controllers.BookUpdate(res, req)
			} else if len(paths) == 3 && paths[2] == "delete" {
				// POST /books/:id/delete
				controllers.BookDeleteConfirm(res, req)
			} else {
				http.Error(res, http.StatusText(400), http.StatusBadRequest)
			}
		} else {
			http.Error(res, http.StatusText(400), http.StatusBadRequest)
		}
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)
	if req.Method == http.MethodGet {
		controllers.HomeIndex(res, req)
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func logout(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)
	if req.Method == http.MethodGet {
		controllers.HomeLogout(res, req)
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func login(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)
	if req.Method == http.MethodGet {
		controllers.HomeLogin(res, req)
	} else if req.Method == http.MethodPost {
		controllers.HomeLoginSubmit(res, req)
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func signup(res http.ResponseWriter, req *http.Request) {
	setSessionCookie(res, req)
	if req.Method == http.MethodGet {
		controllers.HomeSignup(res, req)
	} else if req.Method == http.MethodPost {
		controllers.HomeSignupSubmit(res, req)
	} else {
		http.Error(res, http.StatusText(404), http.StatusNotFound)
	}
}

func setSessionCookie(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		return
	}

	cookie, err := req.Cookie(config.SessionCookieName)

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     config.SessionCookieName,
			Value:    uuid.NewV4().String(),
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(res, cookie)
	}
}
