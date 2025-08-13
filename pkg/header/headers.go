package header

import (
	"apiGo/pkg/errors"
	"log/slog"

	"fmt"
	"net/http"
)

func HeaderWithText(log *slog.Logger, w http.ResponseWriter, strErr []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(strErr); err != nil {
		errors.HandleError(log, w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}

func HeaderWithSub(log *slog.Logger, w http.ResponseWriter, str []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(str); err != nil {
		errors.HandleError(log, w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}
