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

	q := "SELECT c.id, c.name, c.preview, c.created, count(q.id) as quizzes FROM `categories` AS c LEFT JOIN quiz AS q ON q.category_id = c.id GROUP BY c.id ORDER by quizzes DESC"
	if limit > 0 {
		q += fmt.Sprintf(" LIMIT %d", limit)
	}
	err := r.db.Select(&categories, q)
	if err != nil {
		return nil, fmt.Errorf("get categories failed: %w", err)
	}

	return categories, nil
}
