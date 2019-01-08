package main

import (
	"chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type Data struct {
	Thread []data.Thread
	User data.User
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	threads, err := data.Threads()
	if err != nil {
		log.Print(err)
		return
	} else {
		sess, err := session(w, r)
		loggedInUser, err := sess.User()
		data := Data{ threads, loggedInUser}

		if err != nil {
			generateHTML(w, data, "layout","public.navbar", "index")
		} else {
			generateHTML(w, data, "layout", "private.navbar","index")
		}
	}
	}

func FindThread(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	threadId := p.ByName("id")
	thread, err := data.ThreadByID(threadId)
	throwError(err)
	if err != nil {
		log.Print(err)
		return
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, thread, "layout","public.navbar", "thread")
		} else {
			generateHTML(w, thread, "layout", "private.navbar","thread")
		}
	}
}