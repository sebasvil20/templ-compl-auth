package handler

import (
	"net/http"

	"github.com/sebasvil20/templ-compl-auth/internal/common"
	"github.com/sebasvil20/templ-compl-auth/internal/view/page"
)

type IIndexHandler interface {
	RenderIndexPage(w http.ResponseWriter, r *http.Request)
}

type IndexHandler struct {
}

func NewIndexHandler() IIndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) RenderIndexPage(w http.ResponseWriter, r *http.Request) {
	common.RenderTemplate(w, r, page.Index())
}
