package handler

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"mikke-server/domain"
	"mikke-server/usecase"
	"net/http"
	"time"
)

func (p PostQuestion) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
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
	_, err := p.Usecase.PostQuestion(ctx, UserID, b.Title, b.Comment)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	RespondJSON(ctx, w, nil, http.StatusOK)
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

type post struct {
	PostID  domain.PostID `json:"post_id"`
	Title   string        `json:"title"`
	Created time.Time     `json:"created"`
}

type ListPosts struct {
	Usecase usecase.ListPostsUsecase
}

type ListPostsUsecase interface {
	ListPosts(ctx context.Context) (domain.Posts, error)
}

type PostQuestion struct {
	Usecase   usecase.PostUsecase
	Validator *validator.Validate
}

type PostQuestionsUsecace interface {
	PostQuestion(ctx context.Context, user_id int, title string, comment string) (*domain.Post, error)
}
