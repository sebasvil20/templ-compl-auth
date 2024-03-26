package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sebasvil20/templ-compl-auth/internal/common"
	"github.com/sebasvil20/templ-compl-auth/internal/model"
	"github.com/sebasvil20/templ-compl-auth/internal/service"
	"github.com/sebasvil20/templ-compl-auth/internal/view/page"
)

type IAccountHandler interface {
	RenderAccountPage(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type AccountHandler struct {
	BaseHandler
	accountSrv service.IAccountService
}

func NewAccountHandler(v *validator.Validate, decoder *schema.Decoder, sm *sessions.CookieStore, accSrv service.IAccountService) IAccountHandler {
	return &AccountHandler{
		BaseHandler: BaseHandler{
			validator:      v,
			decoder:        decoder,
			sessionManager: sm,
		},
		accountSrv: accSrv,
	}
}

func (h *AccountHandler) RenderAccountPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	page.Accounts().Render(r.Context(), w)
}

func (h *AccountHandler) Signup(w http.ResponseWriter, r *http.Request) {
	reqUser := model.User{}
	err := common.DecodeAndValidateMultipartForm(h.decoder, h.validator, &reqUser, r)
	if err != nil {
		common.HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	token, err := h.accountSrv.Signup(r.Context(), reqUser)
	if err != nil {
		common.HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	session, _ := h.sessionManager.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["jwt-token"] = *token
	err = session.Save(r, w)
	if err != nil {
		common.HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}
	common.HandleResponse(w, nil, nil, http.StatusCreated)
}

func (h *AccountHandler) Login(w http.ResponseWriter, r *http.Request) {
	reqUser := model.Credentials{}
	err := common.DecodeAndValidateMultipartForm(h.decoder, h.validator, &reqUser, r)
	if err != nil {
		common.HandleResponse(w, nil, err, http.StatusBadRequest)
		return
	}

	token, err := h.accountSrv.Login(r.Context(), reqUser)
	if err != nil {
		common.HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}

	session, _ := h.sessionManager.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["jwt-token"] = *token
	err = session.Save(r, w)
	if err != nil {
		common.HandleResponse(w, nil, err, http.StatusInternalServerError)
		return
	}
	common.HandleResponse(w, nil, nil, http.StatusOK)
}

func (h *AccountHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := h.sessionManager.Get(r, "session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
