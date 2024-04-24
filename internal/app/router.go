package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sebasvil20/templ-compl-auth/internal/handler"
)

func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}

func registerRoutes(r *chi.Mux, indexHandler handler.IIndexHandler, account handler.IAccountHandler, user handler.IUserHandler) {

	// Handle static files
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", indexHandler.RenderIndexPage)
	r.Get("/account", account.RenderAccountPage)
	r.Get("/profile", user.RenderUserProfilePage)

	r.Route("/v1/api", func(r chi.Router) {
		r.Post("/signup", account.Signup)
		r.Post("/login", account.Login)
		r.Post("/logout", account.Logout)
	})
}
