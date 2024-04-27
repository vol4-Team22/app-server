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

func (r Repository) ListPosts(ctx context.Context, db Queryer) (domain.Posts, error) {
	posts := domain.Posts{}
	sql := `SELECT post_id, title, created FROM post ORDER BY created DESC`
	if err := db.SelectContext(ctx, &posts, sql); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r Repository) GetPost(ctx context.Context, db Queryer, postId domain.PostID) (domain.Post, error) {
	post := domain.Post{}
	sql := `SELECT post_id, user_id, title, comment, created, modified FROM post WHERE post_id = ?`
	if err := db.GetContext(ctx, &post, sql, int(postId)); err != nil {
		return domain.Post{}, err
	}
	return post, nil
}
