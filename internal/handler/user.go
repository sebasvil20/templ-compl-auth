package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-compl-auth/internal/common"
	"github.com/sebasvil20/templ-compl-auth/internal/config"
	"github.com/sebasvil20/templ-compl-auth/internal/model"
	"github.com/sebasvil20/templ-compl-auth/internal/service"
	"github.com/sebasvil20/templ-compl-auth/internal/view/page"
)

type IUserHandler interface {
	RenderUserProfilePage(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	BaseHandler
	userSrv service.IUserService
}

func NewUserHandler(v *validator.Validate, decoder *schema.Decoder, sm *sessions.CookieStore, uSrv service.IUserService) IUserHandler {
	return &UserHandler{
		BaseHandler: BaseHandler{
			validator:      v,
			decoder:        decoder,
			sessionManager: sm,
		},
		userSrv: uSrv,
	}
}

func (h *UserHandler) RenderUserProfilePage(w http.ResponseWriter, r *http.Request) {
	claims, err := GetJWTClaimsFromSession(h.sessionManager, r)
	if err != nil {
		common.RenderTemplate(w, r, page.ErrorPage(model.HTTPError{StatusCode: "500", Message: err.Error()}))
		return
	}

	user, err := h.userSrv.GetUserBy(r.Context(), claims["email"].(string))
	if err != nil {
		common.RenderTemplate(w, r, page.ErrorPage(model.HTTPError{StatusCode: "500", Message: err.Error()}))
		return
	}
	page.Profile(*user).Render(r.Context(), w)
}

func GetJWTClaimsFromSession(sessionManager *sessions.CookieStore, r *http.Request) (jwt.MapClaims, error) {
	session, err := sessionManager.Get(r, "session")
	if err != nil {
		return nil, err
	}

	jwtToken, ok := session.Values["jwt-token"].(string)
	if !ok {
		return nil, fmt.Errorf("no jwt token found in session")
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unexpected claims type: %T", token.Claims)
	}

	email, ok := claims["email"].(string)
	if !ok || email == "" {
		return nil, fmt.Errorf("no email claim found in jwt token")
	}

	id, ok := claims["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("no id claim found in jwt token")
	}

	return claims, nil
}
