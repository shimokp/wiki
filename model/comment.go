package model

import (
	"database/sql"
)

func (c *Comment) Insert(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`
	insert into comments (user_id, article_id, body)
	values(?, ?, ?)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.UserID, c.ArticleID, c.Body)
}

func (c *Comment) Delete(tx *sql.Tx) (sql.Result, error) {
	stmt, err := tx.Prepare(`delete from comments where comment_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.ID)
}

// ArticlesAll returns all articles.
func CommentsAll(db *sql.DB, id int64) ([]Comment, error) {
	rows, err := db.Query(`select * from comments where article_id=?`, id)
	if err != nil {
		return nil, err
	}
	return ScanComments(rows)
}