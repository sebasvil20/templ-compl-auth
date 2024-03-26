package common

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
)

func HandleResponse(w http.ResponseWriter, data any, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		errBody := map[string]string{"error": err.Error()}
		marshaledErr, _ := json.Marshal(errBody)
		w.Write(marshaledErr)
		return
	}

	if data == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(marshaledData)
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, page templ.Component) {
	w.Header().Set("Content-Type", "text/html")
	page.Render(r.Context(), w)
}
