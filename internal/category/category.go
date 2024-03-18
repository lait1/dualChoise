package category

import (
	"dualChoose/internal/domain"
	"encoding/json"
	"net/http"
)

type repository interface {
	GetCategories() ([]*domain.Category, error)
}
type Service struct {
	repository repository
}

func NewCategoryService(r repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := s.repository.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resultJson, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}
