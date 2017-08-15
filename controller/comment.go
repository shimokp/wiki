package controller

import (
	"net/http"
	"strconv"
	"fmt"
	"database/sql"
	"github.com/shimokp/wiki/sessions"
	"log"
)

//type Comment struct {
//	userID int64
//	body string
//}

type Comment struct {
	DB *sql.DB
}

var comments map[int64] Comment

func (t *Article) CommentSave(w http.ResponseWriter, r *http.Request) error {

	id := r.PostFormValue("id")
	aid, _ := strconv.ParseInt(id, 10, 64)
	body := r.PostFormValue("body")

	//log.Println(body)
	//log.Println("body", body)
	//log.Println("aid", aid)

	comments = make(map[int64]Comment)

	sess, _ := sessions.Get(r, "user")

	//log.Printf("%#v", sess)
	//log.Printf("%#v", sess.Values["id"])

	//comments[aid] = Comment{sess.Values["id"].(int64), body}

	//log.Printf("comments: %#v", comments)

	http.Redirect(w, r, fmt.Sprintf("/article/%s", id), 301)


	return nil
}