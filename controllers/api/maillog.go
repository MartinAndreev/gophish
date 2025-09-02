package api

import (
	"encoding/json"
	"net/http"

	log "github.com/gophish/gophish/logger"

	"github.com/gophish/gophish/models"
)

func (as *Server) MailLog(w http.ResponseWriter, r *http.Request) {
	log.Info("Called with ", r.Method, " ", r.URL.Path)

	switch {
	case r.Method == "POST":
		c := models.MailLogRetryRequest{}
		// Put the request into a campaign
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: "Invalid JSON structure"}, http.StatusBadRequest)
			return
		}
		err = models.GenerateRetryMailLogs(c.Items)
		if err != nil {
			JSONResponse(w, models.Response{Success: false, Message: err.Error()}, http.StatusBadRequest)
			return
		}

		JSONResponse(w, models.Response{Success: true, Message: "Emails were scheduled for resending."}, http.StatusCreated)
	}
}
