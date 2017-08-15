package model

import "database/sql"

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