package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"mikke-server/domain"
	"mikke-server/usecase"
	"net/http"
	"strconv"
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

func (p GetPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	postIDStr := chi.URLParam(r, "post_id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	post, err := p.Usecase.GetPost(ctx, postID)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := post_detail{
		PostID:   post.PostID,
		UserID:   post.UserID,
		Title:    post.Title,
		Comment:  post.Comment,
		Created:  post.Created,
		Modified: post.Modified,
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

type post struct {
	PostID  domain.PostID `json:"post_id"`
	Title   string        `json:"title"`
	Created time.Time     `json:"created"`
}

type post_detail struct {
	PostID   domain.PostID `json:"post_id"`
	UserID   domain.UserID `json:"user_ID"`
	Title    string        `json:"title"`
	Comment  string        `json:"comment"`
	Created  time.Time     `json:"created"`
	Modified time.Time     `json:"modified"`
}

type ListPosts struct {
	Usecase usecase.ListPostsUsecase
}

type ListPostsUsecase interface {
	ListPosts(ctx context.Context) (domain.Posts, error)
}

type GetPost struct {
	Usecase usecase.GetPostUsecase
}

type GetPostUsecase interface {
	GetPost(ctx context.Context, postId int) (*domain.Post, error)
}

type PostQuestion struct {
	Usecase   usecase.PostUsecase
	Validator *validator.Validate
}

type PostQuestionsUsecace interface {
	PostQuestion(ctx context.Context, user_id int, title string, comment string) (*domain.Post, error)
}
