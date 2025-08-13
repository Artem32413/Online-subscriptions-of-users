package appOnlineSub

import (
	"apiGo/internal/onlineSub/model/structs"
	"apiGo/internal/onlineSub/service"
	"apiGo/pkg/errors"
	"apiGo/pkg/requests"

	"context"
	"fmt"
	"net/http"
	"time"
)

type OnlineSubHandler struct {
	svc *service.OnlineSubService
}

func New(svc *service.OnlineSubService) *OnlineSubHandler {
	return &OnlineSubHandler{
		svc: svc,
	}
}

func (s *OnlineSubHandler) AddingARecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.HandleError(w, fmt.Errorf("Неверный метод"), http.StatusBadRequest)
		return
	}
	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := s.svc.AddSubscriptionLogic(ctx, a); err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("Успешное добавление записи")); err != nil {
		errors.HandleError(w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}

func (s *OnlineSubHandler) ConclusionARecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.HandleError(w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AmountOfSubscriptionsLogic(ctx, a)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	req, err := requests.NewMarshall(w, &res)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(req); err != nil {
		errors.HandleError(w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}

func (s *OnlineSubHandler) AllSubscriptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errors.HandleError(w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AllSubscriptionsLogic(ctx)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	req, err := requests.NewMarshall(w, &res)
	if err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(req); err != nil {
		errors.HandleError(w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}

func (s *OnlineSubHandler) UpdateSubscriptionRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		errors.HandleError(w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	if err := s.svc.UpdateSubscriptionRecordLogic(ctx, a); err != nil {
		errors.HandleError(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("Успешная перезапись подписки")); err != nil {
		errors.HandleError(w, fmt.Errorf("Ошибка в выводе данных: %v", err), http.StatusBadRequest)
		return
	}
}

