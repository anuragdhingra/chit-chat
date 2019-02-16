package main

import (
	"log"
	"net/http"

	"github.com/anuragdhingra/lets-chat/data"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

// Signup is the handler which validates url
func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkInvalidRequests(w, r)
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

// SignupAccount creates the user from the form data
func SignupAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)

	user := data.User{
		Username:    r.PostFormValue("username"),
		Email:       r.PostFormValue("email"),
		Password:    encryptPassword(r.PostFormValue("password")),
		HasPassword: true,
	}

	err = user.Create()
	throwError(err)
	http.Redirect(w, r, "/login", http.StatusFound)
}

// Login function handles the login url
func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkInvalidRequests(w, r)
	generateHTML(w, nil, "layout", "public.navbar", "login")
}

// Authenticate handler authenticates the input form data
func Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)

	emailOrUsername := r.PostFormValue("emailOrUsername")
	pass := r.PostFormValue("password")

	user, err := data.UserByEmailOrUsername(emailOrUsername)
	throwError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	session, err := user.CreateSession()
	throwError(err)

	cookie := http.Cookie{
		Name:     "_cookie",
		Value:    session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	log.Print("User successfully logged in")
	http.Redirect(w, r, "/", http.StatusFound)
}

// Logout closes the current session of the user
func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		sess := data.Session{Uuid: cookie.Value}
		err = sess.DeleteByUUID()
		if err != nil {
			log.Print(err)
			return
		}
		cookie := http.Cookie{
			Name:   "_cookie",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		log.Print("Invalid request")
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
