package app

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-compl-auth/internal/handler"
	"github.com/sebasvil20/templ-compl-auth/internal/repository"
	"github.com/sebasvil20/templ-compl-auth/internal/service"
)

func Run() {
	r := newRouter()

	v := validator.New()
	decoder := newDecoder()
	sessionManager := sessions.NewCookieStore([]byte("test"))

	db, err := initializeDatabase()
	if err != nil {
		panic(err)
	}

	indexHandler := handler.NewIndexHandler()

	userRepo := repository.NewUserRepository(db)
	accountService := service.NewAccountService(userRepo)
	accountHandler := handler.NewAccountHandler(v, decoder, sessionManager, accountService)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(v, decoder, sessionManager, userService)

	registerRoutes(r, indexHandler, accountHandler, userHandler)

	http.ListenAndServe(":3000", r)
}
