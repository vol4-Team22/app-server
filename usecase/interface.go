package usecase

import (
	"context"
	"mikke-server/database"
	"mikke-server/domain"
)

type PostAdder interface {
	PostQuestion(ctx context.Context, db database.Execer, p *domain.Post) error
}

type PostLister interface {
	ListPosts(ctx context.Context, db database.Queryer) (domain.Posts, error)
}
