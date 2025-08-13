package app

import (
	"apiGo/internal/onlineSub/transport"

	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func StartMain(ctx context.Context, log *slog.Logger) error {

	log.Info("Сервер запущен")

	mux := transport.AllHandles(ctx, log)

	s := http.Server{
		Addr:    ":8080",
		Handler: LoggingMiddleware(mux),
	}

	go func() {
		<-ctx.Done()
		log.Info("Сервер завершен")
		s.Shutdown(ctx)
	}()

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

type keyRequestID struct{}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		newUuid := r.Header.Get("x-request-id")

		newUuid = uuid.New().String()

		slog := logrus.WithField("request_id", newUuid)

		ctx := context.WithValue(r.Context(), keyRequestID{}, newUuid)
		ctx = context.WithValue(ctx, "slog", slog)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
