package category

import (
	"dualChoose/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"
)

type repository interface {
	GetCategories(limit int) ([]*domain.Category, error)
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
	limitStr := r.URL.Query().Get("limit")
	limit := 0

	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	categories, err := s.repository.GetCategories(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resultJson, _ := json.Marshal(categories)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}
