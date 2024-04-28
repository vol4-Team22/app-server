package usecase

import (
	"context"
	"fmt"
	"mikke-server/database"
	"mikke-server/domain"
)

type SendReplyUsecase struct {
	Repo ReplyAdder
	DB   database.Execer
}

func (p *SendReplyUsecase) SendReply(ctx context.Context, reply *domain.Reply) (*domain.Reply, error) {
	err := p.Repo.SendReply(ctx, p.DB, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

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
