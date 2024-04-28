package usecase

import (
	"context"
	"mikke-server/database"
	"mikke-server/domain"
)

type PostAdder interface {
	SendPost(ctx context.Context, db database.Execer, p *domain.Post) error
}

type PostLister interface {
	ListPosts(ctx context.Context, db database.Queryer) (domain.Posts, error)
}

type PostGeter interface {
	GetPost(ctx context.Context, db database.Queryer, postId domain.PostID) (domain.Post, error)
}

type ReplyLister interface {
	ListReplies(ctx context.Context, db database.Queryer, postId domain.PostID) (domain.Replies, error)
}
