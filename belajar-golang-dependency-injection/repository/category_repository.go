package repository

import (
	"context"
	"database/sql"
	"dhany007/belajar-golang-restful-api/model/domain"
)

// transactional sqlnya adalah pointer
// klik aja, karena dia struct, harus pass by reference
type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
