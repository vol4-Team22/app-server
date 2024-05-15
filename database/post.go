package database

import (
	"context"
	"fmt"
	"mikke-server/domain"
	"time"

	"github.com/Masterminds/squirrel"
)

func (r Repository) SendPost(ctx context.Context, db Execer, p *domain.Post) error {
	// TODO: 認証機能を実装後変更
	// 現在はすべて7777として登録
	p.Created = r.Clocker.Now()
	p.Modified = p.Created
	query, param, err := squirrel.
		Insert("post").
		Columns(
			"user_id",
			"title",
			"comment",
			"created",
			"modified",
		).
		Values(int(p.UserID), p.Title, p.Comment, p.Created, p.Modified).
		ToSql()
	if err != nil {
		return err
	}
	result, err := db.ExecContext(ctx, query, param...)
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

func (r Repository) ListPosts(ctx context.Context, db Queryer) ([]*domain.Post, error) {
	q := squirrel.
		Select(
			"post_id",
			"user_id",
			"title",
			"comment",
			"created",
			"modified",
		).
		From("post").
		OrderBy("created DESC")
	query, params, err := q.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error in ToSql")
	}
	var rows []*Post
	if err := db.SelectContext(ctx, &rows, query, params...); err != nil {
		return nil, err
	}
	posts := make([]*domain.Post, 0, len(rows))
	for _, p := range rows {
		posts = append(posts, p.postToDomain())
	}

	return posts, nil
}

func (r Repository) GetPost(ctx context.Context, db Queryer, postId domain.PostID) (domain.Post, error) {
	post := domain.Post{}
	q := squirrel.
		Select(
			"post_id",
			"user_id",
			"title",
			"comment",
			"created",
			"modified",
		).
		From("post as p").
		Where(squirrel.Eq{"p.post_id": postId})
	query, params, err := q.ToSql()
	if err != nil {
		return domain.Post{}, fmt.Errorf("error in ToSql")
	}
	if err := db.GetContext(ctx, &post, query, params...); err != nil {
		return domain.Post{}, err
	}
	return post, nil
}

type Post struct {
	PostID   int       `db:"post_id"`
	UserID   int       `db:"user_id"`
	Title    string    `db:"title"`
	Comment  string    `db:"comment"`
	Created  time.Time `db:"created"`
	Modified time.Time `db:"modified"`
}

func (p Post) postToDomain() *domain.Post {
	return &domain.Post{
		PostID:   domain.PostID(p.PostID),
		UserID:   7777,
		Title:    p.Title,
		Comment:  p.Title,
		Created:  p.Created,
		Modified: p.Modified,
	}
}
