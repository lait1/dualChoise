package quiz

import (
	"dualChoose/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type quizRepository interface {
	GetQuizzesByCategory(categoryID int) ([]*domain.Quiz, error)
	GetPopularQuizzes() ([]*domain.Quiz, error)
	GetQuiz(ID int) (*domain.Quiz, error)
	GetQuizOptions(ID int) ([]*domain.Option, error)
}
type categoryRepository interface {
	GetCategory(categoryID int) (*domain.Category, error)
}
type Service struct {
	quizRepository     quizRepository
	categoryRepository categoryRepository
}

func NewQuizService(qr quizRepository, cr categoryRepository) *Service {
	return &Service{qr, cr}
}

type CategoryInfo struct {
	ID      int            `json:"categoryId"`
	Name    string         `json:"categoryName"`
	Quizzes []*domain.Quiz `json:"quizzes"`
}

func (s *Service) GetQuizzesByCategory(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	categoryID, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	quizzes, err := s.quizRepository.GetQuizzesByCategory(categoryID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	category, err := s.categoryRepository.GetCategory(categoryID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	categoryInfo := CategoryInfo{
		ID:      category.ID,
		Name:    category.Name,
		Quizzes: quizzes,
	}
	resultJson, err := json.Marshal(categoryInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

func (s *Service) GetPopularQuizzes(w http.ResponseWriter, r *http.Request) {
	quizzes, err := s.quizRepository.GetPopularQuizzes()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resultJson, err := json.Marshal(quizzes)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}

type QuizGame struct {
	ID      int              `json:"id"`
	Name    string           `json:"name"`
	Options []*domain.Option `json:"options"`
}

func (s *Service) StartQuiz(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	quizID, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	quizInfo, err := s.quizRepository.GetQuiz(quizID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	quizOptions, err := s.quizRepository.GetQuizOptions(quizID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	quiz := QuizGame{
		ID:      quizInfo.ID,
		Name:    quizInfo.Name,
		Options: quizOptions,
	}
	resultJson, err := json.Marshal(quiz)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}
