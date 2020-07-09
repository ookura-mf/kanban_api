package models

type Task struct {
	ID int
	KanbanID int
	Kanban   Kanban
	Title    string
	Content  string
}
