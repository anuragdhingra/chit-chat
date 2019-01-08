package main

import (
	_ "chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
	"net/http"
)

func main() {

	mux := httprouter.New()
	mux.ServeFiles("/public/*filepath", http.Dir("public"))

	mux.GET("/", Index)
	mux.GET("/threads/:id", FindThread)
	mux.GET("/signup", Signup)
	mux.POST("/signup_account", SignupAccount)
	mux.GET("/login", Login)
	mux.POST("/authenticate", Authenticate)
	mux.GET("/logout", Logout)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	http2.ConfigureServer(server, &http2.Server{})
	server.ListenAndServe()
}
