package repository

import (
	"api/model/domain"
	scurity "api/security"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (repository *RepoImpl) Register(ctx context.Context, tx *sql.Tx, RegCategory domain.Register) {
	sql := "insert into register(nama, passwords) values (?, ?)"
	encrypt, err := scurity.PassEncrypt(RegCategory)
	if err != nil {
		panic("error from hashing")
	}
	result, err := tx.ExecContext(ctx, sql, RegCategory.Name, encrypt)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	RegCategory.Id = int(id)
}

func (repository *RepoImpl) Login(ctx context.Context, tx *sql.Tx, login domain.Login) (string, error) {
	sql := "SELECT nama, passwords FROM register WHERE nama = ?"
	rows, err := tx.QueryContext(ctx, sql, login.Name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	loglist := domain.Login{}
	if rows.Next() {
		err := rows.Scan(&loglist.Name, &loglist.Password)
		if err != nil {
			panic(err)
		}
		token, eror := scurity.ClaimsJwt(loglist, login.Password)
		if eror != nil {
			panic("error jwt")
		}
		fmt.Println(token)
		return token, nil
	} else {
		return "", errors.New("data not found")
	}
}
