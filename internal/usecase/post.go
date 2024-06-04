package usecase

import (
	"context"
	"fmt"
	"mikke-server/internal/database"
	domain2 "mikke-server/internal/domain"
)

type PostUsecase struct {
	Repo PostAdder
	DB   database.Execer
}

func (p *PostUsecase) SendPost(ctx context.Context, user_id int, title string, comment string) (*domain2.Post, error) {
	post := &domain2.Post{
		UserID:  domain2.UserID(user_id),
		Title:   title,
		Comment: comment,
	}
	err := p.Repo.SendPost(ctx, p.DB, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

type ListPostsUsecase struct {
	Repo PostLister
	DB   database.Queryer
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
	Repo PostGeter
	DB   database.Queryer
}
