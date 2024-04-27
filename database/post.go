package database

import (
	"context"
	"mikke-server/domain"
)

func (r Repository) PostQuestion(ctx context.Context, db Execer, p *domain.Post) error {
	// TODO: 認証機能を実装後変更
	// 現在はすべて7777として登録
	p.Created = r.Clocker.Now()
	p.Modified = p.Created
	sql := `INSERT INTO post (user_id, title, comment, created, modified) VALUES (?, ?, ?, ? ,?)`
	result, err := db.ExecContext(ctx, sql, int(p.UserID), p.Title, p.Comment, p.Created, p.Modified)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.PostID = domain.PostID(id)
	return nil
}
