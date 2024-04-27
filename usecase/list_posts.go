package usecase

import (
	"context"
	"fmt"
	"mikke-server/database"
	"mikke-server/domain"
)

type ListPostsUsecase struct {
	Repo PostLister
	DB   database.Queryer
}

func (p *ListPostsUsecase) ListPosts(ctx context.Context) (domain.Posts, error) {
	posts, err := p.Repo.ListPosts(ctx, p.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return posts, nil
}
