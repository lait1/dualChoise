package translation

import (
	"dualChoose/internal/domain"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type repository interface {
	GetTranslations(local string) (*map[string]string, error)
}
type Service struct {
	repository repository
}

func NewTranslationService(r repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetTranslations(w http.ResponseWriter, r *http.Request) {
	local := chi.URLParam(r, "local")

	validLocals := map[domain.Lang]bool{domain.RU: true, domain.EN: true, domain.ES: true}
	if !validLocals[domain.Lang(local)] {
		http.Error(w, "Invalid local value", http.StatusBadRequest)
		return
	}

	translations, err := s.repository.GetTranslations(local)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resultJson, _ := json.Marshal(translations)

	w.Header().Set("Content-Type", "application/json")
	w.Write(resultJson)
}
