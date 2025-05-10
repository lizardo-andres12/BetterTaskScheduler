package data

import (
	"context"
	"database/sql"
	"fmt"

	"go.tskmgr.com/internal/models"
)

const values = "(id, teamId, importance, author, description, type, completed)"

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) Create(ctx context.Context, task *models.Task) error {
	stmt, err := tr.db.PrepareContext(
		ctx,
		fmt.Sprintf("insert into test %s values (?,?,?,?,?,?,?)", values),
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		task.ID,
		task.TeamID,
		task.Importance,
		task.Author,
		task.Description,
		task.Type,
		task.Completed,
	)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TaskRepository) ReadOne(ctx context.Context, task *models.Task) (*models.Task, error) {
	//stmt, err :=
	return nil, nil
}
