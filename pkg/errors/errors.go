package errors

import (
	"log/slog"
	"net/http"
)

func HandleError(log *slog.Logger, w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	http.Error(w, err.Error(), status)
	log.Error(err.Error())
}
