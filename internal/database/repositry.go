package database

import (
	"context"
	"mikke-server/internal/domain"
)

type PostRepository interface {
	SendPost(ctx context.Context, db Execer, user_id int, title string, comment string) error
	ListPosts(ctx context.Context, db Queryer) ([]*domain.Post, error)
	GetPost(ctx context.Context, db Queryer, postId int) (*domain.Post, error)
}

type ReplyAdder interface {
	SendReply(ctx context.Context, db Execer, p *domain.Reply) error
}

type ReplyLister interface {
	ListReplies(ctx context.Context, db Queryer, postId domain.PostID) (domain.Replies, error)
}
