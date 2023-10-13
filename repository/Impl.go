package repository

import (
	"api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type RepoImpl struct {
}

// repository layer
func NewRepository() Repository {
	return &RepoImpl{}
}

func (repository *RepoImpl) Save(ctx context.Context, tx *sql.Tx, datalist domain.Datasiswa) domain.Datasiswa {
	sql := "insert into category(name, kelas) values (?, ?)"
	result, err := tx.ExecContext(ctx, sql, datalist.Name, datalist.Kelas)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	datalist.Id = int(id)
	return datalist
}
func (repository *RepoImpl) Update(ctx context.Context, tx *sql.Tx, datalist domain.Datasiswa) domain.Datasiswa {
	sql := "update category set name = ?,kelas = ? kelas where id = ?"
	_, err := tx.ExecContext(ctx, sql, datalist.Name, datalist.Kelas, datalist.Id)
	if err != nil {
		panic(err)
	}
	return datalist
}
func (repository *RepoImpl) Delete(ctx context.Context, tx *sql.Tx, datalist domain.Datasiswa) {
	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, sql, datalist.Id)
	if err != nil {
		panic(err)
	}

}
func (repository *RepoImpl) FindByid(ctx context.Context, tx *sql.Tx, datalistId int) (domain.Datasiswa, error) {
	sql := "select id, name, kelas from category where id = ?"
	rows, err := tx.QueryContext(ctx, sql, datalistId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	datalist := domain.Datasiswa{}
	if rows.Next() {
		err := rows.Scan(&datalist.Id, &datalist.Name, &datalist.Kelas)
		if err != nil {
			panic(err)
		}
		return datalist, nil
	} else {
		return datalist, errors.New("not data found")
	}
}
func (repository *RepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Datasiswa {
	sql := "select id, name, kelas from category"
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var categories []domain.Datasiswa
	for rows.Next() {
		datalist := domain.Datasiswa{}
		err := rows.Scan(&datalist.Id, &datalist.Name, &datalist.Kelas)
		if err != nil {
			panic(err)
		}
		categories = append(categories, datalist)
	}
	return categories
}
