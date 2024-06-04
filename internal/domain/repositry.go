package domain

import (
	"context"
	"mikke-server/internal/database"
)

type PostRepositry interface {
	SendPost(ctx context.Context, db database.Execer, user_id int, title string, comment string) error
	ListPosts(ctx context.Context, db database.Queryer) ([]*Post, error)
	GetPost(ctx context.Context, db database.Queryer, postId int) (*Post, error)
}

type ReplyAdder interface {
	SendReply(ctx context.Context, db database.Execer, p *Reply) error
}

type ReplyLister interface {
	ListReplies(ctx context.Context, db database.Queryer, postId PostID) (Replies, error)
}
