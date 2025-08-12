package interfaces

import (
	"apiGo/internal/onlineSub/model/structs"

	"context"
	"net/http"
)

type HandlersOnlineSub interface {
	AddARecord(w http.ResponseWriter, r *http.Request)
	ConclusionARecord(w http.ResponseWriter, r *http.Request)
}

type OnlineSubRepo interface {
	AddSubscriptionLogic(ctx context.Context, str structs.Subscription) error
	AmountOfSubscriptionsLogic(ctx context.Context, str structs.Subscription) ([]structs.Subscription, error)
}
