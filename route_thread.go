package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anuragdhingra/lets-chat/data"
	"github.com/julienschmidt/httprouter"
)

// ThreadInfoPublic handles public data of User
type ThreadInfoPublic struct {
	Thread    data.Thread
	CreatedBy data.User
	Posts     []PostInfoPublic
}

// ThreadInfoPrivate handles private data of User
type ThreadInfoPrivate struct {
	Thread    data.Thread
	CreatedBy data.User
	User      data.User
	Posts     []PostInfoPublic
}

// NewThread function creates a new thread
func NewThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess, err := session(w, r)
	if err == nil {
		loggedInUser, err := sess.User()
		data := ThreadsInfoPrivate{nil, loggedInUser}
		if err != nil {
			log.Print(err)
			return
		}
		generateHTML(w, data, "layout", "private.navbar", "new.thread")
	} else {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

// CreateThread function requests a new thread from NewThread function
func CreateThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)
	sess, err := session(w, r)
	throwError(err)
	user, err := sess.User()
	throwError(err)

	createThreadRequest := data.CreateThreadRequest{
		r.PostFormValue("topic"),
		user.Id,
	}
	threadID, err := createThreadRequest.Create()
	log.Print(threadId)
	throwError(err)
	url := "/threads/" + strconv.Itoa(threadId)
	log.Print(url)
	http.Redirect(w, r, url, http.StatusFound)
}

// FindThread finds the thread by Id
func FindThread(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	threadID := p.ByName("id")
	thread, err := data.ThreadByID(threadId)
	if err != nil {
		log.Print(err)
		return
	}
	user, err := data.UserById(thread.UserId)
	throwError(err)
	posts, err := data.PostsByThreadId(thread.Id)
	throwError(err)
	postList := CreatePostList(posts)
	sess, err := session(w, r)
	if err != nil {
		data := ThreadInfoPublic{thread, user, postList}
		generateHTML(w, data, "layout", "public.navbar", "public.thread")
	} else {
		loggedInUser, err := sess.User()
		throwError(err)
		data := ThreadInfoPrivate{thread, user, loggedInUser, postList}
		generateHTML(w, data, "layout", "private.navbar", "private.thread")
	}
}

// CreateThreadList creates a list of threads
func CreateThreadList(threads []data.Thread) (threadListPublic []ThreadInfoPublic) {
	for _, thread := range threads {
		threadUserID := thread.UserId
		user, err := data.UserById(threadUserId)
		throwError(err)
		threadInfoPublic := ThreadInfoPublic{thread, user, nil}
		threadListPublic = append(threadListPublic, threadInfoPublic)
	}
	return
}
