package database

import (
	"context"
	"mikke-server/domain"
)

func (r Repository) SendReply(ctx context.Context, db Execer, p *domain.Reply) error {
	// TODO: 認証機能を実装後変更
	// 現在はすべて7777として登録
	p.Created = r.Clocker.Now()
	p.Modified = p.Created
	sql := `INSERT INTO reply (post_id, user_id, title, comment, created, modified) VALUES (?, ?, ?, ?, ? ,?)`
	result, err := db.ExecContext(ctx, sql, int(p.PostID), int(p.UserID), p.Title, p.Comment, p.Created, p.Modified)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ReplyID = domain.ReplyID(id)
	return nil
}

func (r Repository) ListReplies(ctx context.Context, db Queryer, postId domain.PostID) (domain.Replies, error) {
	replies := domain.Replies{}
	sql := `SELECT reply_id, post_id, user_id, title, comment, created, modified FROM reply WHERE post_id = ? ORDER BY reply_id ASC`
	if err := db.SelectContext(ctx, &replies, sql, int(postId)); err != nil {
		return nil, err
	}
	return replies, nil
}
