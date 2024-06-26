package storage

import (
	"dualChoose/internal/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
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

func (r *QuizRepository) GetPopularQuizzes(limit int) ([]*domain.Quiz, error) {
	var quizzes = make([]*domain.Quiz, 0)

	q := "SELECT q.id, q.name, q.preview FROM `quiz` AS q LEFT JOIN `quiz_options` AS qo ON q.id = qo.quiz_id GROUP BY q.id ORDER BY SUM(qo.wins) desc "
	if limit > 0 {
		q += fmt.Sprintf(" LIMIT %d", limit)
	}
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

func (r *QuizRepository) GetOptionsByQuizId(ID int) ([]*domain.Option, error) {
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

	// Shuffle options
	randSource := rand.New(rand.NewSource(time.Now().UnixNano()))
	randSource.Shuffle(len(options), func(i, j int) {
		options[i], options[j] = options[j], options[i]
	})

	// Return the first 10 options, or all if less than 10
	numOptions := 100
	if len(options) < numOptions {
		numOptions = len(options)
	}

	return options[:numOptions], nil
}

func (r *QuizRepository) GetQuizOptions(quizID int) ([]*domain.QuizOption, error) {
	var options = make([]*domain.QuizOption, 0)

	q := "SELECT quiz_id, option_id, wins " +
		"FROM `quiz_options` " +
		"WHERE quiz_id = ? "
	err := r.db.Select(&options, q, quizID)
	if err != nil {
		return nil, fmt.Errorf("get options failed: %w", err)
	}
	return options, nil
}

func (r *QuizRepository) SaveResult(result domain.QuizResult) error {
	q := "UPDATE `quiz_options` SET `wins`= `wins` + 1 WHERE `quiz_id` = ? and `option_id` = ?"
	_, err := r.db.Exec(q, result.QuizID, result.OptionID)
	if err != nil {
		return fmt.Errorf("save result failed: %w", err)
	}
	return nil
}
