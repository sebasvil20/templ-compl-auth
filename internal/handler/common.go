package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

type BaseHandler struct {
	validator      *validator.Validate
	decoder        *schema.Decoder
	sessionManager *sessions.CookieStore
}
