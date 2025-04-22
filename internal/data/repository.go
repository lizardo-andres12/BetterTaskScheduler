package data

import (
	"context"

	models "github.com/lizardo-andres12/taskmanager/internal"
)

type Repository interface {
	Create(context.Context, *models.Task) error
	ReadOne(context.Context, uint64) (*models.Task, error)
	ReadMany(context.Context, []uint64) ([]models.Task, error)
	Update(context.Context, *models.Task) error
	Delete(context.Context, uint64) error
}
