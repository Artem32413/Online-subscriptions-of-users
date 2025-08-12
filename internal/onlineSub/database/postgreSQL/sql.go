package postgreSQL

import (
	"apiGo/internal/onlineSub/config/databaseConfig"
	"apiGo/internal/onlineSub/database/postgreSQL/convert"
	"apiGo/internal/onlineSub/model/structs"

	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var (
	addition   = `INSERT INTO Subscription (service_name, price, user_id, start_date) VALUES ($1, $2, $3, $4)`
	filtration = `SELECT COALESCE(SUM(price), 0) FROM Subscription WHERE start_date >= $1 AND start_date <= $2 AND user_id = $3 AND service_name = $4`
	getAll     = `SELECT * FROM Subscription`
)

type DBService struct {
	db *pgx.Conn
}

func New(db *databaseConfig.PostgreSQL) *DBService {
	return &DBService{db: db.Db}
}

func (s *DBService) AddARecordSQL(ctx context.Context, str structs.Subscription) error {
	date := convert.ConvertTime(str.Start_date)

	if _, err := s.db.Exec(ctx, addition, str.Service_name, str.Price, str.User_id, date); err != nil {
		return fmt.Errorf("Ошибка в добавлении подписки: %v", err)
	}

	return nil
}

func (s *DBService) ConclusionARecordSQL(ctx context.Context, str structs.Subscription) (*int, error) {
	start_date := convert.ConvertTime(str.Start_date)

	end_date := convert.ConvertTime(str.End_date)

	r, err := s.db.Query(ctx, filtration, start_date, end_date, str.User_id, str.Service_name)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в запросе с фильтрацией: %v", err)
	}

	defer r.Close()

	var allAmounts *int

	if r.Next() {
		var a structs.Subscription

		if err = r.Scan(&a.Price); err != nil {
			return nil, fmt.Errorf("Ошибка в сканировании: %v", err)
		}
		allAmounts = &a.Price
	}

	return allAmounts, nil
}

func (s *DBService) AllSubscriptionsSQL(ctx context.Context) ([]structs.Subscription, error) {
	var a []structs.Subscription

	r, err := s.db.Query(ctx, addition)
	if err != nil {
		return nil, fmt.Errorf("Ошибка в добавлении подписки: %v", err)
	}

	for r.Next() {
		var all structs.Subscription
		if err := r.Scan(&all); err != nil {
			return nil, fmt.Errorf("Ошибка в сканировании: %v", err)
		}
		a = append(a, all)
	}

	return a, nil
}
