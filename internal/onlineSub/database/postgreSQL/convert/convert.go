package convert

import (
	"fmt"
	"time"
)

func ConvertTime(str string) (time.Time, error) {
	date, err := time.Parse("01-2006", str)
	if err != nil {
		return time.Time{}, fmt.Errorf("Ошибка в форматировании даты: %v", err)
	}

	return date, nil
}

func ConvertString(t time.Time) string {
	return t.Format("01-2006")
}