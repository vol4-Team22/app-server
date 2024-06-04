package database

import (
	"context"
	domain2 "mikke-server/internal/domain"

	"github.com/Masterminds/squirrel"
)

func (r Repository) SendReply(ctx context.Context, db Execer, p *domain2.Reply) error {
	// TODO: 認証機能を実装後変更
	// 現在はすべて7777として登録
	p.Created = r.Clocker.Now()
	p.Modified = p.Created
	query, param, err := squirrel.
		Insert("reply").
		Columns(
			"post_id",
			"user_id",
			"title",
			"comment",
			"created",
			"modified",
		).
		Values(int(p.PostID), int(p.UserID), p.Title, p.Comment, p.Created, p.Modified).
		ToSql()
	result, err := db.ExecContext(ctx, query, param...)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ReplyID = domain2.ReplyID(id)
	return nil
}

func (r Repository) ListReplies(ctx context.Context, db Queryer, postId domain2.PostID) (domain2.Replies, error) {
	replies := domain2.Replies{}
	q := squirrel.
		Select(
			"reply_id",
			"post_id",
			"user_id",
			"title",
			"comment",
			"created",
			"modified",
		).
		From("reply as r").
		Where(squirrel.Eq{"r.post_id": postId}).
		OrderBy("created ASC")
	query, params, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	if err := db.SelectContext(ctx, &replies, query, params...); err != nil {
		return nil, err
	}
	return replies, nil
}
