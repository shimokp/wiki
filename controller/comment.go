package controller

import (
	"net/http"
	"strconv"
	"fmt"
	"database/sql"
	"github.com/shimokp/wiki/httputil"
	"github.com/shimokp/wiki/sessions"
	"github.com/shimokp/wiki/model"
	"log"
)

type Comment struct {
	DB *sql.DB
}

//var comments map[int64] Comment

func (c *Comment) CommentSave(w http.ResponseWriter, r *http.Request) error {

	id := r.PostFormValue("id")
	aid, _ := strconv.ParseInt(id, 10, 64)
	body := r.PostFormValue("body")

	//log.Println(body)
	log.Println("body", body)
	log.Println("aid", aid)

	//comments = make(map[int64]Comment)

	sess, _ := sessions.Get(r, "user")

	log.Printf("%#v", sess)
	//log.Printf("%#v", sess.Values["id"])

	//comments[aid] = Comment{sess.Values["id"].(int64), body}

	//log.Printf("comments: %#v", comments)

	var comment model.Comment
	comment.ArticleID = aid
	comment.Body = body
	comment.UserID =sess.Values["id"].(int64)

	if err := TXHandler(c.DB, func(tx *sql.Tx) error {
		_, err := comment.Insert(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		//id, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	http.Redirect(w, r, fmt.Sprintf("/article/%s", id), 301)

	return nil
}

func (c Comment) CommentDelete(w http.ResponseWriter, r *http.Request) error {
	var comment model.Comment
	id := r.PostFormValue("id")
	if id == "" {
		return &httputil.HTTPError{Status: http.StatusBadRequest}
	}
	articleId := r.PostFormValue("article_id")
	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	comment.ID = aid
	if err := TXHandler(c.DB, func(tx *sql.Tx) error {
		if _, err := comment.Delete(tx); err != nil {
			return err
		}
		return tx.Commit()
	}); err != nil {
		return err
	}

	if articleId == ""  {
		http.Redirect(w, r, fmt.Sprintf("/article/%v", articleId), 301)
		return nil
	}
	http.Redirect(w, r,"/", 301)
	return nil
}