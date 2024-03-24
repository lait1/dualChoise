package quiz

import (
	"dualChoose/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type quizRepository interface {
	GetQuizzesByCategory(categoryID int) ([]*domain.Quiz, error)
	GetPopularQuizzes() ([]*domain.Quiz, error)
	GetQuiz(ID int) (*domain.Quiz, error)
	GetOptionsByQuizId(ID int) ([]*domain.Option, error)
	SaveResult(result domain.QuizResult) error
	GetQuizOptions(quizID int) ([]*domain.QuizOption, error)
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

	options, err := s.quizRepository.GetOptionsByQuizId(quizID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	quiz := QuizGame{
		ID:      quizInfo.ID,
		Name:    quizInfo.Name,
		Options: options,
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

func (s *Service) SaveResult(w http.ResponseWriter, r *http.Request) {
	var result domain.QuizResult
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = s.quizRepository.SaveResult(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	quizInfo, err := s.quizRepository.GetQuizOptions(result.QuizID)

	quizGameResult := domain.QuizGameResult{
		QuizID:         result.QuizID,
		PercentageWins: calculatePercentageWins(result.OptionID, quizInfo),
	}
	resultJson, err := json.Marshal(quizGameResult)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}

func calculatePercentageWins(resultOptionId int, quizInfo []*domain.QuizOption) int {
	var totalWins int
	var resultOptionWins int
	for _, option := range quizInfo {
		if option.OptionId == resultOptionId {
			resultOptionWins = option.Wins
		}
		totalWins += option.Wins
	}
	fmt.Println(resultOptionWins, totalWins)

	if resultOptionWins == 0 || totalWins == 0 {
		return 0
	}

	return int(float32(resultOptionWins) / float32(totalWins) * 100)
}
