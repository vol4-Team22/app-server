package server

import (
	"context"
	"mikke-server/internal/config"
	"mikke-server/internal/database"
	"mikke-server/internal/handler"
	"mikke-server/internal/usecase"
	"mikke-server/tools/clock"
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
	u := usecase.NewPostUsecase(r, db)
	hd := handler.NewPostHandler(u, v)
	if err != nil {
		return nil, cleanup, err
	}
	mux.Post("/post", hd.SendPost)
	mux.Get("/list", hd.ListPosts)
	mux.Get("/post/{post_id}", hd.GetPost)
	//sr := &handler.SendReply{
	//	Usecase:   usecase.SendReplyUsecase{DB: db, Repo: &r},
	//	Validator: v,
	//}
	//mux.Post("/reply", sr.ServeHTTP)
	//lr := &handler.ListReplies{
	//	Usecase: usecase.ListRepliesUsecase{DB: db, Repo: &r},
	//}
	//mux.Get("/reply/list/{post_id}", lr.ServeHTTP)
	return mux, cleanup, err
}
