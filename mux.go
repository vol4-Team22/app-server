package main

import (
	"context"
	"mikke-server/config"
	"mikke-server/database"
	"mikke-server/handler"
	"mikke-server/tools/clock"
	"mikke-server/usecase"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	r := database.Repository{Clocker: clock.RealClocker{}}
	v := validator.New()
	db, cleanup, err := database.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	ps := &handler.SendPost{
		Usecase:   usecase.PostUsecase{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/post", ps.ServeHTTP)
	lp := &handler.ListPosts{
		Usecase: usecase.ListPostsUsecase{DB: db, Repo: &r},
	}
	mux.Get("/list", lp.ServeHTTP)
	lt := &handler.GetPost{
		Usecase: usecase.GetPostUsecase{DB: db, Repo: &r},
	}
	mux.Get("/post/{post_id}", lt.ServeHTTP)
	sr := &handler.SendReply{
		Usecase:   usecase.SendReplyUsecase{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/reply", sr.ServeHTTP)
	lr := &handler.ListReplies{
		Usecase: usecase.ListRepliesUsecase{DB: db, Repo: &r},
	}
	mux.Get("/reply/list/{post_id}", lr.ServeHTTP)
	return mux, cleanup, err
}
