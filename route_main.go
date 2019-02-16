package main

import (
	"log"
	"net/http"

	"github.com/anuragdhingra/lets-chat/data"
	"github.com/julienschmidt/httprouter"
)

// ThreadsInfoPrivate represents a list of private threads
type ThreadsInfoPrivate struct {
	ThreadList []ThreadInfoPublic
	User       data.User
}

// ThreadsInfoPublic represents a list of public threads
type ThreadsInfoPublic struct {
	ThreadList []ThreadInfoPublic
}

// Index function creates index of thread
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	threads, err := data.Threads()
	if err != nil {
		log.Print(err)
		return
	}
	sess, err := session(w, r)
	loggedInUser, err := sess.User()
	if err != nil {
		data := ThreadsInfoPublic{CreateThreadList(threads)}
		generateHTML(w, data, "layout", "public.navbar", "index")
	} else {
		data := ThreadsInfoPrivate{CreateThreadList(threads), loggedInUser}
		generateHTML(w, data, "layout", "private.navbar", "index")
	}
}
