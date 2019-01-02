package main

import (
	"chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

func SignupAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)

	user := data.User{
		Username:r.PostFormValue("username"),
		Email:r.PostFormValue("email"),
		Password: encryptPassword(r.PostFormValue("password")),
	}

	err = user.Create()
	throwError(err)
	http.Redirect(w, r, "/", 302)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	generateHTML(w, nil, "layout", "public.navbar", "login")
}

func Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)

	emailOrUsername := r.PostFormValue("emailOrUsername")
	pass := r.PostFormValue("password")
	user, err := data.UserByEmailOrUsername(emailOrUsername)
	throwError(err)
	log.Print(user.Password + "/" + pass)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", 302)
		return
	} else {
		log.Print("User successfully logged in")
		http.Redirect(w, r, "/", 302)
	}
}