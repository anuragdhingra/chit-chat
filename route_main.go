package main

import (
	"github.com/julienschmidt/httprouter"
	"lets-chat/data"
	"log"
	"net/http"
)

type ThreadsInfoPrivate struct {
	ThreadList []ThreadInfoPublic
	User data.User
}

type ThreadsInfoPublic struct {
	ThreadList []ThreadInfoPublic
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	threads, err := data.Threads()
	if err != nil {
		log.Print(err)
		return
	} else {
		sess, err := session(w, r)
		loggedInUser, err := sess.User()
		if err != nil {
			data := ThreadsInfoPublic{CreateThreadList(threads)}
			generateHTML(w, data, "layout","public.navbar", "index")
		} else {
			data := ThreadsInfoPrivate{CreateThreadList(threads), loggedInUser}
			generateHTML(w, data, "layout", "private.navbar","index")
		}
	}
	}