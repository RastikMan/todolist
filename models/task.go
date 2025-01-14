package models

type Task struct {
	ID       int    `json:"id"`       // Унікальний ідентифікатор
	Title    string `json:"title"`    // Заголовок задачі
	Complete bool   `json:"complete"` // Статус виконання
}
