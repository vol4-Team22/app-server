package usecase

import (
	"context"
	"mikke-server/database"
	"mikke-server/domain"
)

type PostUsecase struct {
	Repo PostAdder
	DB   database.Execer
}

func (p *PostUsecase) PostQuestion(ctx context.Context, user_id int, title string, comment string) (*domain.Post, error) {
	post := &domain.Post{
		UserID:  domain.UserID(user_id),
		Title:   title,
		Comment: comment,
	}
	err := p.Repo.PostQuestion(ctx, p.DB, post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
