package usecase

import (
	"context"
	"fmt"
	"mikke-server/internal/database"
	"mikke-server/internal/domain"
)

type PostUsecase struct {
	repo domain.PostRepositry
	db   database.Interface
}

func NewPostUsecase(repo domain.PostRepositry, db database.Interface) *PostUsecase {
	return &PostUsecase{
		repo: repo,
		db:   db,
	}
}

func (p *PostUsecase) SendPost(ctx context.Context, user_id int, title, comment string) error {
	err := p.repo.SendPost(ctx, p.db, user_id, title, comment)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostUsecase) ListPosts(ctx context.Context) ([]*domain.Post, error) {
	posts, err := p.repo.ListPosts(ctx, p.db)
	if err != nil {
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}
	return posts, nil
}

func (p *PostUsecase) GetPost(ctx context.Context, post_id int) (*domain.Post, error) {
	post, err := p.repo.GetPost(ctx, p.db, post_id)
	if err != nil {
		return nil, fmt.Errorf("failed to get post: %w", err)
	}
	return post, nil
}
