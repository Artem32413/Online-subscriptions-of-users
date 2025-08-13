package header

import (
	"apiGo/pkg/errors"

	"fmt"
	"net/http"
)

func HeaderWithText(w http.ResponseWriter, strErr []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(strErr); err != nil {
		errors.HandleError(w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}

func HeaderWithSub(w http.ResponseWriter, str []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(str); err != nil {
		errors.HandleError(w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}
