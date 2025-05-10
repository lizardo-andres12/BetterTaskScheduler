package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lizardo-andres12/taskmanager/internal"
)

const userValues = "id, teamId, email, password"

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(ctx context.Context, user *models.User) error {
	stmt, err := ur.db.PrepareContext(
		ctx,
		fmt.Sprintf("insert into users (%s) values (?, ?, ?, ?)", userValues),
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		user.User.ID,
		user.User.TeamID,
		user.User.Email,
		user.Password,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) ReadOneByID(ctx context.Context, id uint64) (*models.User, error) {
	stmt, err := ur.db.PrepareContext(
		ctx,
		fmt.Sprintf("select %s from users where id=?", userValues),
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	row := stmt.QueryRowContext(
		ctx,
		id,
	)
	if err = row.Scan(
		&user.User.ID,
		&user.User.TeamID,
		&user.User.Email,
		&user.Password,
	); err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) FindUserIDByEmail(ctx context.Context, email string) (*uint64, error) {
	stmt, err := ur.db.PrepareContext(ctx, "select id from users where email=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var id uint64
	row := stmt.QueryRowContext(ctx, email)
	if err = row.Scan(&id); err != nil {
		return nil, err
	}
	return &id, nil
}

func (ur *UserRepository) DeleteOneByID(ctx context.Context, id uint64) error {
	stmt, err := ur.db.PrepareContext(
		ctx,
		"delete from users where id=?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
