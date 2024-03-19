package storage

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type TranslationRepository struct {
	db *sqlx.DB
}
type Translation struct {
	VarName string `db:"var_name"`
	Value   string `db:"value"`
}

func NewTranslationRepository(db *sqlx.DB) *TranslationRepository {
	return &TranslationRepository{
		db: db,
	}
}

func (r *TranslationRepository) GetTranslations(local string) (*map[string]string, error) {
	var translations []Translation

	q := fmt.Sprintf("SELECT `var_name`, `%s` as value FROM `translations`;", local)
	err := r.db.Select(&translations, q)
	if err != nil {
		return nil, fmt.Errorf("get translations failed: %w", err)
	}
	translationMap := make(map[string]string)
	for _, t := range translations {
		translationMap[t.VarName] = t.Value
	}
	return &translationMap, nil
}
