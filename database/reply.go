package database

import (
	"context"
	"mikke-server/domain"
)

func (r Repository) ListReplies(ctx context.Context, db Queryer, postId domain.PostID) (domain.Replies, error) {
	replies := domain.Replies{}
	sql := `SELECT reply_id, post_id, user_id, title, comment, created, modified FROM reply WHERE post_id = ? ORDER BY reply_id ASC`
	if err := db.SelectContext(ctx, &replies, sql, int(postId)); err != nil {
		return nil, err
	}
	return replies, nil
}
