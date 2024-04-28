package handler

import (
	"context"
	"github.com/go-chi/chi"
	"mikke-server/domain"
	"mikke-server/usecase"
	"net/http"
	"strconv"
)

func (p ListReplies) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postIDStr := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDStr)
	posts, err := p.Usecase.ListReplies(ctx, domain.PostID(postID))
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []domain.Reply{}
	for _, ps := range posts {
		rsp = append(rsp, domain.Reply{
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

type ListReplies struct {
	Usecase usecase.ListRepliesUsecase
}

type ListRepliesUsecase interface {
	ListReplies(ctx context.Context, postID domain.PostID) (domain.Replies, error)
}
