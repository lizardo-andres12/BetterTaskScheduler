package main

import (
	"context"
	"time"

	"github.com/lizardo-andres12/taskmanager/internal"
	"github.com/lizardo-andres12/taskmanager/internal/data"

	"go.uber.org/fx"
)

func TryDB(tr *data.TaskRepository) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	err := tr.Create(ctx, &models.Task{
		ID: 1,
		TeamID: 1,
		Importance: 10,
		Author: "Meat",
		Description: "Simple task",
		Type: "Physical",
		Completed: false,
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fx.New(
		fx.Module(
			"database",

			// Private constructors.
			fx.Provide(
				fx.Private,

				data.GenerateConfig,
				data.NewDatabase,
			),

			// Public constructors.
			fx.Provide(
				data.NewTaskRepository,
			),
		),

		fx.Invoke(TryDB),
	).Run()
}
