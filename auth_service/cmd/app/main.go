package main

import (
	"context"
	"os"
	"time"

	//"github.com/lizardo-andres12/taskmanager/internal"
	//"github.com/lizardo-andres12/taskmanager/internal"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lizardo-andres12/taskmanager/internal/data"
	"github.com/lizardo-andres12/taskmanager/internal/service"

	"go.uber.org/fx"
)

func TryDB(us *service.UserService) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancel()

	tokenString, err := us.LogIn(ctx, "lah@uga.edu", "abcdefgh")
	if err != nil {
		return err
	}
	_, err = jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SecretKey")), nil
	})
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
				data.NewUserRepository,
			),
		),

		fx.Module(
			"service",

			fx.Provide(
				service.NewUserService,
			),
		),
		fx.Invoke(TryDB),
	).Run()
}
