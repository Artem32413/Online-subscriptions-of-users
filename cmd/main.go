//	@title			API Go Service
//	@version		1.0
//	@description	Тестовое задание Junior Golang Developer Effective Mobile

//	@contact.name	API Support
//	@contact.url	http://www.example.com/support
//	@contact.email	support@example.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1
//	@schemes	http

package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"apiGo/cmd/app"
	_ "apiGo/docs"
)

// @Summary		Запуск приложения
// @Description	Основная точка входа для API сервиса
func main() {
	if err := realMain(); err != nil {
		slog.Error(err.Error())
		return
	}
}

// realMain содержит основную логику приложения
//
//	@Summary		Основная логика приложения
//	@Description	Инициализирует контекст и запускает API сервер
func realMain() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	return app.StartMain(ctx)
}
