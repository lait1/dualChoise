package storage

import (
	"dualChoose/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QuizRepository struct {
	db *sqlx.DB
}

func NewQuizRepository(db *sqlx.DB) *QuizRepository {
	return &QuizRepository{
		db: db,
	}
}

func (r *QuizRepository) GetQuizzesByCategory(categoryID int) ([]*domain.Quiz, error) {
	var quizzes = make([]*domain.Quiz, 0)

	q := "SELECT q.id, q.name, q.preview " +
		"FROM `quiz` AS q " +
		"WHERE `category_id` = ?"
	err := r.db.Select(&quizzes, q, categoryID)
	if err != nil {
		return nil, fmt.Errorf("get quizzes failed: %w", err)
	}
	return quizzes, nil
}

func (r *QuizRepository) GetPopularQuizzes() ([]*domain.Quiz, error) {
	var quizzes = make([]*domain.Quiz, 0)

	q := "SELECT q.id, q.name, q.preview FROM `quiz` AS q LEFT JOIN `quiz_options` AS qo ON q.id = qo.quiz_id GROUP BY q.id ORDER BY SUM(qo.wins) desc LIMIT 3;"
	err := r.db.Select(&quizzes, q)
	if err != nil {
		return nil, fmt.Errorf("get popular quizzes failed: %w", err)
	}
	return quizzes, nil
}

func (r *QuizRepository) GetQuiz(ID int) (*domain.Quiz, error) {
	var quiz domain.Quiz

	q := "SELECT q.id, q.name, q.preview FROM `quiz` AS q  WHERE q.id = ?"
	err := r.db.Get(&quiz, q, ID)
	if err != nil {
		return nil, fmt.Errorf("get quiz failed: %w", err)
	}
	return &quiz, nil
}

func (r *QuizRepository) GetQuizOptions(ID int) ([]*domain.Option, error) {
	var options = make([]*domain.Option, 0)

	q := "SELECT o.id, o.name, o.preview, o.priority " +
		"FROM `quiz` AS q " +
		"LEFT JOIN `quiz_options` AS qo ON q.id = qo.quiz_id " +
		"LEFT JOIN `options` AS o ON o.id = qo.option_id " +
		"WHERE q.id = ? "
	err := r.db.Select(&options, q, ID)
	if err != nil {
		return nil, fmt.Errorf("get options failed: %w", err)
	}
	return options, nil
}
