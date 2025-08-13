package transport

import (
	"apiGo/internal/onlineSub/appOnlineSub"
	"apiGo/internal/onlineSub/config/databaseConfig"
	"apiGo/internal/onlineSub/database/postgreSQL"
	"apiGo/internal/onlineSub/service"
	swaggerpkg "apiGo/internal/onlineSub/transport/swaggerPkg"

	"context"
	"log/slog"
	"net/http"
)

func AllHandles(ctx context.Context, log *slog.Logger) *http.ServeMux {
	db, err := databaseConfig.ConstructorDB(ctx)
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	repo := postgreSQL.New(db)
	svc := service.New(repo)
	handlers := appOnlineSub.New(svc, log)

	mux := http.NewServeMux()

	swaggerpkg.AddSwaggerRoutes(mux)

	mux.HandleFunc("/add/", handlers.AddingARecord)
	mux.HandleFunc("/sum/", handlers.ConclusionARecord)
	mux.HandleFunc("/all/", handlers.AllSubscriptions)
	mux.HandleFunc("/update/", handlers.UpdateSubscriptionRecord)
	mux.HandleFunc("/delete/", handlers.DeleteSubscriptionRecord)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
