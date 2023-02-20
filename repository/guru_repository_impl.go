package repository

import (
	"belajar-go-restapi-gin-latihan2/helper"
	"belajar-go-restapi-gin-latihan2/model/domain"
	"database/sql"
	"errors"

	"golang.org/x/net/context"
)

type GuruRepositoryImpl struct {
}

func NewGuruRepositoryImpl() GuruRepository {
	return &GuruRepositoryImpl{}
}

func (guru_repository *GuruRepositoryImpl) Create(ctx context.Context, DB *sql.DB, guru domain.Guru) domain.Guru {
	sql := "Insert Into guru(id_guru,name,status) Values(?,?,?)"
	_, err := DB.ExecContext(ctx, sql, guru.Id_guru, guru.Name, guru.Status)
	helper.PanicIfError(err)
	return guru
}

func (guru_repository *GuruRepositoryImpl) Update(ctx context.Context, DB *sql.DB, guru domain.Guru) domain.Guru {
	sql := "Update guru set name=?,status=? Where id_guru=?"
	_, err := DB.ExecContext(ctx, sql, guru.Name, guru.Status, guru.Id_guru)
	helper.PanicIfError(err)
	return guru
}

func (guru_repository *GuruRepositoryImpl) FindById(ctx context.Context, DB *sql.DB, IdGuru string) (domain.Guru, error) {
	sql := "Select id_guru,name,status from guru Where id_guru=?"
	rows, err := DB.QueryContext(ctx, sql, IdGuru)
	helper.PanicIfError(err)
	defer rows.Close()
	guru := domain.Guru{}
	if rows.Next() {
		err := rows.Scan(&guru.Id_guru, &guru.Name, &guru.Status)
		helper.PanicIfError(err)
		return guru, nil
	} else {
		return guru, errors.New("NOT FOUND")
	}
}

func (guru_repository *GuruRepositoryImpl) FindAll(ctx context.Context, DB *sql.DB) []domain.Guru {
	sql := "Select id_guru,name,status from guru"
	rows, err := DB.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	Allguru := []domain.Guru{}
	guru := domain.Guru{}
	for rows.Next() {
		err := rows.Scan(&guru.Id_guru, &guru.Name, &guru.Status)
		helper.PanicIfError(err)
		Allguru = append(Allguru, guru)
	}
	return Allguru
}

func (guru_repository *GuruRepositoryImpl) Delete(ctx context.Context, DB *sql.DB, IdGuru string) {
	sql := "Delete from guru Where id_guru=?"
	_, err := DB.ExecContext(ctx, sql, IdGuru)
	helper.PanicIfError(err)
}
