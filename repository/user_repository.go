package repository

import (
	"context"
	"database/sql"

	"github.com/vincen320/user-service/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User) bool
	FindById(ctx context.Context, tx *sql.Tx, UserId int) (domain.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
