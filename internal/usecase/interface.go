package usecase

import (
	"context"
	"mikke-server/internal/database"
	domain2 "mikke-server/internal/domain"
)

type PostAdder interface {
	SendPost(ctx context.Context, db database.Execer, p *domain2.Post) error
}

type PostLister interface {
	ListPosts(ctx context.Context, db database.Queryer) ([]*domain2.Post, error)
}

type PostGeter interface {
	GetPost(ctx context.Context, db database.Queryer, postId domain2.PostID) (domain2.Post, error)
}

type ReplyAdder interface {
	SendReply(ctx context.Context, db database.Execer, p *domain2.Reply) error
}

type ReplyLister interface {
	ListReplies(ctx context.Context, db database.Queryer, postId domain2.PostID) (domain2.Replies, error)
}
