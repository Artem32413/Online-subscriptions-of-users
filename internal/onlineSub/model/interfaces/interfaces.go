package interfaces

import (
	"apiGo/internal/onlineSub/model/structs"

	"context"
)

type HandlersOnlineSub interface {
	AddSubscriptionLogic(ctx context.Context, str structs.Subscription) error
	AmountOfSubscriptionsLogic(ctx context.Context, str structs.Subscription) (structs.Sum, error)
	AllSubscriptionsLogic(ctx context.Context) ([]structs.Subscription, error) 
	UpdateSubscriptionRecordLogic(ctx context.Context, a structs.Subscription) error
	DeleteSubscriptionRecordLogic(ctx context.Context, a structs.Subscription) error 
}
