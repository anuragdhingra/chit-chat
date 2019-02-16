package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/anuragdhingra/lets-chat/data"
	"github.com/julienschmidt/httprouter"
)

// PostInfoPublic organises public post data
type PostInfoPublic struct {
	Post      data.Post
	CreatedBy data.User
}

// CreatePost functions creates post from form data
func CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess, err := session(w, r)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	} else {
		user, _ := sess.User()

		err = r.ParseForm()
		throwError(err)
		postBody := r.PostFormValue("body")
		threadIDString := r.PostFormValue("id")
		threadID, err := strconv.Atoi(threadIdString)
		throwError(err)

		postRequest := data.PostRequest{postBody, user.Id, threadId}
		_, err = postRequest.CreatePost()
		throwError(err)
		http.Redirect(w, r, "/threads/"+threadIdString, http.StatusFound)
	}
}

// CreatePostList function creates a list of posts of User
func CreatePostList(posts []data.Post) (postListPublic []PostInfoPublic) {
	for _, post := range posts {
		postUserID := post.UserId
		user, err := data.UserById(postUserId)
		throwError(err)
		postInfoPublic := PostInfoPublic{post, user}
		postListPublic = append(postListPublic, postInfoPublic)
	}
	return
}
