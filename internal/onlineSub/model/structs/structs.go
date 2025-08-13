package structs

type Subscription struct {
	Service_name string `json:"service_name" example:"Yandex Plus"`                     // Название сервиса, предоставляющего подписку
	Price        int    `json:"price" example:"400"`                                    // Стоимость месячной подписки в рублях
	User_id      string `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"` // ID пользователя
	Start_date   string `json:"start_date,omitempty" example:"07-2025"`                 // Дата начала подписки (месяц и год)
	End_date     string `json:"end_date,omitempty" example:"03-2026"`                   // Опционально дата окончания подписки
}

type Subscription2 struct {
	Service_name string `json:"service_name" example:"Yandex Plus"`                     // Название сервиса, предоставляющего подписку
	Price        int    `json:"price" example:"400"`                                    // Стоимость месячной подписки в рублях
	User_id      string `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"` // ID пользователя
	Start_date   string `json:"start_date,omitempty" example:"07-2025"`                 // Дата начала подписки (месяц и год)
}