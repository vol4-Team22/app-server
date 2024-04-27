package database

import (
	"context"
	"mikke-server/domain"
)

func (r Repository) ListPosts(ctx context.Context, db Queryer) (domain.Posts, error) {
	posts := domain.Posts{}
	sql := `SELECT post_id, title, created FROM post ORDER BY created DESC`
	if err := db.SelectContext(ctx, &posts, sql); err != nil {
		return nil, err
	}
	return posts, nil
}
