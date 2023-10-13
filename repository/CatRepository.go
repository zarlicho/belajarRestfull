package repository

import (
	"api/model/domain"
	"context"
	"database/sql"
)

type Repository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Datasiswa) domain.Datasiswa
	Update(ctx context.Context, tx *sql.Tx, category domain.Datasiswa) domain.Datasiswa
	Delete(ctx context.Context, tx *sql.Tx, category domain.Datasiswa)
	FindByid(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Datasiswa, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Datasiswa
}
