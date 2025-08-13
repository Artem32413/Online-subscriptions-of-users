package interfaces

import (
	"net/http"
)

type HandlersOnlineSub interface {
	AddARecord(w http.ResponseWriter, r *http.Request)
	ConclusionARecord(w http.ResponseWriter, r *http.Request)
	AllSubscriptions(w http.ResponseWriter, r *http.Request)
	UpdateSubscriptionRecord(w http.ResponseWriter, r *http.Request)
	DeleteSubscriptionRecord(w http.ResponseWriter, r *http.Request)
}
