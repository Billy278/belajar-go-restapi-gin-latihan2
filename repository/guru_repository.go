package repository

import (
	"database/sql"

	"belajar-go-restapi-gin-latihan2/model/domain"

	"golang.org/x/net/context"
)

type GuruRepository interface {
	Create(ctx context.Context, DB *sql.DB, guru domain.Guru) domain.Guru
	Update(ctx context.Context, DB *sql.DB, guru domain.Guru) domain.Guru
	FindById(ctx context.Context, DB *sql.DB, IdGuru string) (domain.Guru, error)
	FindAll(ctx context.Context, DB *sql.DB) []domain.Guru
	Delete(ctx context.Context, DB *sql.DB, IdGuru string)
}
