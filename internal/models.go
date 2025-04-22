package models

type Task struct {
	ID, TeamID uint64
	Importance uint8
	Author, Description, Type string
	Completed bool
}
