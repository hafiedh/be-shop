package postgres

import (
	"be-shop/internal/app/models"
	"be-shop/internal/app/repo/postgres/queries"
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"go.uber.org/dig"
)

type (
	UserRepo interface {
		CreateUser(ctx context.Context, req models.User) (id int, err error)
		GetUserByEmail(ctx context.Context, email string) (user models.User, err error)
		GetUserByID(ctx context.Context, id int64) (user models.User, err error)
	}

	UserRepoImpl struct {
		dig.In

		*sql.DB
	}
)

func NewUserRepo(impl UserRepoImpl) UserRepo {
	return &impl
}

func (u *UserRepoImpl) CreateUser(ctx context.Context, req models.User) (id int, err error) {
	_, err = u.QueryContext(ctx, queries.QueryCreateUser, req.Email, req.Username, req.Password)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[UserRepoImpl.CreateUser] error while CreateUser err: %v", err.Error()))
		return id, err
	}

	return id, nil
}

func (u *UserRepoImpl) GetUserByEmail(ctx context.Context, email string) (user models.User, err error) {
	row := u.QueryRowContext(ctx, queries.QueryGetUserByEmail, email)
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[UserRepoImpl.GetUserByEmail] error while GetUserByEmail err: %v", err.Error()))
		return
	}
	return
}

func (u *UserRepoImpl) GetUserByID(ctx context.Context, id int64) (user models.User, err error) {
	row := u.QueryRowContext(ctx, queries.QueryGetUserByID, id)
	err = row.Scan(&user.Email, &user.Username, &user.Password)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[UserRepoImpl.GetUserByID] error while GetUserByID err: %v", err.Error()))
		return user, err
	}

	return user, nil
}
