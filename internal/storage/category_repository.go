package storage

import (
	"dualChoose/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (r *CategoryRepository) GetCategories(limit int) ([]*domain.Category, error) {
	var categories = make([]*domain.Category, 0)

	q := "SELECT c.id, c.name, c.preview, c.created FROM categories AS c"
	if limit > 0 {
		q += fmt.Sprintf(" LIMIT %d", limit)
	}
	err := r.db.Select(&categories, q)
	if err != nil {
		return nil, fmt.Errorf("get categories failed: %w", err)
	}

	return categories, nil
}
