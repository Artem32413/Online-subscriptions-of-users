package convert

import (
	"log/slog"
	"time"
)

func ConvertTime(str string) time.Time {
	date, err := time.Parse("01-2006", str)
	if err != nil {
		slog.Error("Ошибка в форматировании даты: %v", err)
		return time.Time{}
	}

	return date
}