package appOnlineSub

import (
	"apiGo/internal/onlineSub/model/structs"
	"apiGo/internal/onlineSub/service"
	"apiGo/pkg/errors"
	"apiGo/pkg/header"
	"apiGo/pkg/requests"
	"log/slog"

	"context"
	"fmt"
	"net/http"
	"time"
)

type OnlineSubHandler struct {
	svc *service.OnlineSubService
	l   *slog.Logger
}

func New(svc *service.OnlineSubService, l *slog.Logger) *OnlineSubHandler {
	return &OnlineSubHandler{
		svc: svc,
		l:   l,
	}
}

// AddingARecord apiGo
// @Summary Добавление новой подписки
// @Description Добавляет запись о новой подписке пользователя
// @Tags Subscriptions
// @Accept json
// @Produce plain
// @Param subscription body structs.Subscription2 true "Данные подписки"
// @Success 200 {string} string "Успешное добавление записи"
// @Failure 400 {string} string "Ошибка в зависимости от контекста"
// @Router /add/ [post]
func (s *OnlineSubHandler) AddingARecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.HandleError(s.l, w, fmt.Errorf("Неверный метод"), http.StatusBadRequest)
		return
	}
	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if err := s.svc.AddSubscriptionLogic(ctx, a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	header.HeaderWithText(s.l, w, []byte("Успешное добавление записи"))
}

// ConclusionARecord apiGo
// @Summary Получение информации о стоимости всех подписок за выбранный период с фильтрацией по id пользователя и названию подписки
// @Description Возвращает стоимость всех подписок
// @Tags Subscriptions
// @Accept json
// @Produce json
// @Param subscription body structs.Subscription true "Данные для поиска подписок"
// @Success 200 {array} structs.Sum
// @Failure 400 {string} string "Ошибка в зависимости от контекста"
// @Router /sum/ [post]
func (s *OnlineSubHandler) ConclusionARecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errors.HandleError(s.l, w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AmountOfSubscriptionsLogic(ctx, a)
	if err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	req, err := requests.NewMarshall(w, &res)
	if err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	header.HeaderWithSub(s.l, w, req)
}

// AllSubscriptions apiGo
// @Summary Получение всех подписок
// @Description Возвращает список всех подписок в системе
// @Tags Subscriptions
// @Produce json
// @Success 200 {array} structs.Subscription2
// @Failure 400 {string} string "Ошибка в зависимости от контекста"
// @Router /all/ [get]
func (s *OnlineSubHandler) AllSubscriptions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errors.HandleError(s.l, w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	res, err := s.svc.AllSubscriptionsLogic(ctx)
	if err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	req, err := requests.NewMarshall(w, &res)
	if err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	header.HeaderWithSub(s.l, w, req)
}

// UpdateSubscriptionRecord apiGo
// @Summary Обновление подписки
// @Description Обновляет информацию о существующей подписке
// @Tags Subscriptions
// @Accept json
// @Produce plain
// @Param subscription body structs.Subscription2 true "Обновленные данные подписки"
// @Success 200 {string} string "Успешная перезапись подписки"
// @Failure 400 {string} string "Ошибка в зависимости от контекста"
// @Router /update/ [put]
func (s *OnlineSubHandler) UpdateSubscriptionRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		errors.HandleError(s.l, w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	if err := s.svc.UpdateSubscriptionRecordLogic(ctx, a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	header.HeaderWithText(s.l, w, []byte("Успешная перезапись подписки"))
}

// DeleteSubscriptionRecord apiGo
// @Summary Удаление подписки
// @Description Удаляет информацию о подписке
// @Tags Subscriptions
// @Accept json
// @Produce plain
// @Param subscription body structs.Subscription2 true "Данные подписки для удаления"
// @Success 200 {string} string "Успешное удаление подписки"
// @Failure 400 {string} string "Ошибка в зависимости от контекста"
// @Router /delete/ [delete]
func (s *OnlineSubHandler) DeleteSubscriptionRecord(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		errors.HandleError(s.l, w, fmt.Errorf("Не верный метод"), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var a structs.Subscription

	if err := requests.NewDec(r, &a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	if err := s.svc.DeleteSubscriptionRecordLogic(ctx, a); err != nil {
		errors.HandleError(s.l, w, err, http.StatusBadRequest)
		return
	}

	header.HeaderWithText(s.l, w, []byte("Успешное удаление подписки"))
}
