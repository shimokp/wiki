//go:generate scaneo $GOFILE

package model

import "time"

// User returns model object for user.
type User struct {
	ID      int64      `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Salt    string     `json:"salt"`
	Salted  string     `json:"salted"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

// Article returns model object for article.
type Article struct {
	ID      int64      `json:"id"`
	Title   string     `json:"title"`
	Body    string     `json:"body"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

type Comment struct {
	ID int64			`json:"id"`
	UserID int64		`json:"user_id"`
	ArticleID int64		`json:"article_id"`
	Body string			`json:"body"`
	Created *time.Time  `json:"created"`
	Updated *time.Time  `json:"updated"`
}
