package service

import (
	"apiGo/internal/onlineSub/database/postgreSQL"
	"apiGo/internal/onlineSub/model/structs"

	"context"
)

type OnlineSubService struct {
	repo *postgreSQL.DBService
}

func New(repo *postgreSQL.DBService) *OnlineSubService {
	return &OnlineSubService{
		repo: repo,
	}
}

func (o *OnlineSubService) AddSubscriptionLogic(ctx context.Context, str structs.Subscription) error {
	return o.repo.AddARecordSQL(ctx, str)
}

func (o *OnlineSubService) AmountOfSubscriptionsLogic(ctx context.Context, str structs.Subscription) (structs.Sum, error) {
	return o.repo.ConclusionARecordSQL(ctx, str)
}

func (o *OnlineSubService) AllSubscriptionsLogic(ctx context.Context) ([]structs.Subscription, error) {
	return o.repo.AllSubscriptionsSQL(ctx)
}

func (o *OnlineSubService) UpdateSubscriptionRecordLogic(ctx context.Context, a structs.Subscription) error {
	return o.repo.UpdateSubscriptionRecordSQL(ctx, a)
}

func (o *OnlineSubService) DeleteSubscriptionRecordLogic(ctx context.Context, a structs.Subscription) error {
	return o.repo.DeleteSubscriptionRecordSQL(ctx, a)
}
