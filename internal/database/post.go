package database

import (
	"context"
	"fmt"
	"mikke-server/internal/domain"
	"time"

	"github.com/Masterminds/squirrel"
)

func (r Repository) SendPost(ctx context.Context, db Execer, user_id int, title, comment string) error {
	// TODO: 認証機能を実装後変更
	// 現在はすべて7777として登録
	var created, modified time.Time
	created = r.Clocker.Now()
	modified = created
	query, param, err := squirrel.
		Insert("post").
		Columns(
			"user_id",
			"title",
			"comment",
			"created",
			"modified",
		).
		Values(user_id, title, comment, created, modified).
		ToSql()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, query, param...)
	if err != nil {
		return err
	}
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

func (r Repository) GetPost(ctx context.Context, db Queryer, postId int) (*domain.Post, error) {
	var post Post
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
		return nil, fmt.Errorf("error in ToSql")
	}
	if err := db.GetContext(ctx, &post, query, params...); err != nil {
		return nil, err
	}
	rsp := post.postToDomain()
	return rsp, nil
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
