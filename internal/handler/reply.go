package handler

import (
	"context"
	"encoding/json"
	domain2 "mikke-server/internal/domain"
	"mikke-server/internal/usecase"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

func (p SendReply) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		PostID  int    `json:"post_id" validate:"required"`
		Title   string `json:"title" validate:"required"`
		Comment string `json:"comment" validate:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	if err := p.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}
	UserID := 7777
	reply := &domain2.Reply{
		PostID:  domain2.PostID(b.PostID),
		UserID:  domain2.UserID(UserID),
		Title:   b.Title,
		Comment: b.Comment,
	}
	_, err := p.Usecase.SendReply(ctx, reply)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, nil, http.StatusOK)
}

func (p ListReplies) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postIDStr := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDStr)
	posts, err := p.Usecase.ListReplies(ctx, domain2.PostID(postID))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []domain2.Reply{}
	for _, ps := range posts {
		rsp = append(rsp, domain2.Reply{
			ReplyID:  ps.ReplyID,
			PostID:   ps.PostID,
			UserID:   ps.UserID,
			Title:    ps.Title,
			Comment:  ps.Comment,
			Created:  ps.Created,
			Modified: ps.Modified,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

type SendReply struct {
	Usecase   usecase.SendReplyUsecase
	Validator *validator.Validate
}

type SendReplyUsecase interface {
	SendReply(ctx context.Context, reply *domain2.Reply) (*domain2.Reply, error)
}

type ListReplies struct {
	Usecase usecase.ListRepliesUsecase
}

type ListRepliesUsecase interface {
	ListReplies(ctx context.Context, postID domain2.PostID) (domain2.Replies, error)
}
