package usecase

import (
	"context"
	"fmt"
	"mikke-server/database"
	"mikke-server/domain"
)

func (p *ListRepliesUsecase) ListReplies(ctx context.Context, postID domain.PostID) (domain.Replies, error) {
	replies, err := p.Repo.ListReplies(ctx, p.DB, postID)
	if err != nil {
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return replies, nil
}

type ListRepliesUsecase struct {
	Repo ReplyLister
	DB   database.Queryer
}
