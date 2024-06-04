package handler

import (
	"context"
	"encoding/json"
	domain2 "mikke-server/internal/domain"
	"mikke-server/internal/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	post      *usecase.PostUsecase
	Validator *validator.Validate
}

func NewPostHandler(u *usecase.PostUsecase, v *validator.Validate) *PostHandler {
	return &PostHandler{
		post:      u,
		Validator: v,
	}
}

func (h PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	if err := h.Validator.Struct(b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
	}
	UserID := 7777
	_, err := h.post.Post(ctx, UserID, b.Title, b.Comment)
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
	PostID  domain2.PostID `json:"post_id"`
	Title   string         `json:"title"`
	Created time.Time      `json:"created"`
}

type post_detail struct {
	PostID   domain2.PostID `json:"post_id"`
	UserID   domain2.UserID `json:"user_ID"`
	Title    string         `json:"title"`
	Comment  string         `json:"comment"`
	Created  time.Time      `json:"created"`
	Modified time.Time      `json:"modified"`
}

type PostQuestionsUsecace interface {
	SendPost(ctx context.Context, user_id int, title string, comment string) (*domain2.Post, error)
}

type ListPosts struct {
	Usecase usecase.ListPostsUsecase
}

type ListPostsUsecase interface {
	ListPosts(ctx context.Context) (domain2.Posts, error)
}

type GetPost struct {
	Usecase usecase.GetPostUsecase
}

type GetPostUsecase interface {
	GetPost(ctx context.Context, postId int) (*domain2.Post, error)
}
