package transport

import (
	"apiGo/internal/onlineSub/appOnlineSub"
	"apiGo/internal/onlineSub/config/databaseConfig"
	"apiGo/internal/onlineSub/database/postgreSQL"
	"apiGo/internal/onlineSub/model/interfaces"
	"apiGo/internal/onlineSub/service"

	"context"
	"log/slog"
	"net/http"
)

type InventoryService struct {
	interfaces.HandlersOnlineSub
	*databaseConfig.PostgreSQL
}

func AddSwaggerRoutes(mux *http.ServeMux) {
	// Страница с Swagger UI (используем CDN)
	mux.HandleFunc("/docs/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8" />
<title>Swagger UI</title>
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist/swagger-ui-bundle.js"></script>
<script>
  const ui = SwaggerUIBundle({
    url: "/docs/swagger.json",
    dom_id: '#swagger-ui',
  });
</script>
</body>
</html>`))
	})

	// Отдача файла swagger.json
	mux.HandleFunc("/docs/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})
}

func AllHandles(ctx context.Context) *http.ServeMux {
	db, err := databaseConfig.ConstructorDB(ctx)
	if err != nil {
		slog.Error(err.Error())
		return nil
	}

	repo := postgreSQL.New(db)
	svc := service.New(repo)
	handlers := appOnlineSub.New(svc)

	mux := http.NewServeMux()

	AddSwaggerRoutes(mux)

	mux.HandleFunc("/add/", handlers.AddingARecord)
	mux.HandleFunc("/sum/", handlers.ConclusionARecord)
	mux.HandleFunc("/all/", handlers.AllSubscriptions)
	mux.HandleFunc("/update/", handlers.UpdateSubscriptionRecord)
	mux.HandleFunc("/api/health/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
