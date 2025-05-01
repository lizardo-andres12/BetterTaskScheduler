package main

import (
	"context"
	"fmt"
	"time"

	//"github.com/lizardo-andres12/taskmanager/internal"
	"github.com/lizardo-andres12/taskmanager/internal/data"

	"go.uber.org/fx"
)

func TryDB(ur *data.UserRepository) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	err := ur.DeleteOneByID(ctx, 1)
	if err != nil {
		return fmt.Errorf("No id")
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
				data.NewUserRepository,
			),
		),

		fx.Invoke(TryDB),
	).Run()
}
