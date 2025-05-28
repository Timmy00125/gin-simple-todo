package models

// Todo represents a todo item
type Todo struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
