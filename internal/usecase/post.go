package usecase

import (
	"context"
	"fmt"
	"mikke-server/internal/database"
	"mikke-server/internal/domain"
)

type PostUsecase struct {
	repo domain.PostRepositry
	db   database.Execer
}

func NewPostUsecase(repo domain.PostRepositry, db database.Execer) *PostUsecase {
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

func (p *ListPostsUsecase) ListPosts(ctx context.Context) (domain2.Posts, error) {
	posts, err := p.Repo.ListPosts(ctx, p.DB)

	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return domain2.Posts(posts), nil
}

func (u GetPostUsecase) GetPost(ctx context.Context, postId int) (domain2.Post, error) {
	post, err := u.Repo.GetPost(ctx, u.DB, domain2.PostID(postId))
	if err != nil {
		return domain2.Post{}, fmt.Errorf("failed to get: %w", err)
	}
	return post, nil
}

type GetPostUsecase struct {
	Repo domain.PostGeter
	DB   database.Queryer
}
