package usecase

import (
	"context"
	"fmt"
	"mikke-server/database"
	"mikke-server/domain"
)

type PostUsecase struct {
	Repo PostAdder
	DB   database.Execer
}

func (p *PostUsecase) SendPost(ctx context.Context, user_id int, title string, comment string) (*domain.Post, error) {
	post := &domain.Post{
		UserID:  domain.UserID(user_id),
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

func (p *ListPostsUsecase) ListPosts(ctx context.Context) (domain.Posts, error) {
	posts, err := p.Repo.ListPosts(ctx, p.DB)

	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return domain.Posts(posts), nil
}

func (u GetPostUsecase) GetPost(ctx context.Context, postId int) (domain.Post, error) {
	post, err := u.Repo.GetPost(ctx, u.DB, domain.PostID(postId))
	if err != nil {
		return domain.Post{}, fmt.Errorf("failed to get: %w", err)
	}
	return post, nil
}

type GetPostUsecase struct {
	Repo PostGeter
	DB   database.Queryer
}
