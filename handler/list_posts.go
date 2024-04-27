package handler

import (
	"context"
	"mikke-server/domain"
	"mikke-server/usecase"
	"net/http"
	"time"
)

type post struct {
	PostID  domain.PostID `json:"post_id"`
	Title   string        `json:"title"`
	Created time.Time     `json:"created"`
}

func (p ListPosts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	posts, err := p.Usecase.ListPosts(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := []post{}
	for _, ps := range posts {
		rsp = append(rsp, post{
			PostID:  ps.PostID,
			Title:   ps.Title,
			Created: ps.Created,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

type ListPosts struct {
	Usecase usecase.ListPostsUsecase
}

type ListPostsUsecase interface {
	ListPosts(ctx context.Context) (domain.Posts, error)
}
