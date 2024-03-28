package main

import (
	"dualChoose/internal/category"
	"dualChoose/internal/common"
	"dualChoose/internal/config"
	"dualChoose/internal/quiz"
	"dualChoose/internal/storage"
	"dualChoose/internal/translation"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func main() {
	cfg, err := config.New()
	logger, err := common.NewLogger(cfg.Level)
	db, err := common.SetDataBaseMySQL(logger, cfg)
	if err != nil {
		logger.Fatal("cannot connect to MySQL", zap.Error(err),
			zap.String("db", cfg.MySQLDatabase))
	}

	defer func() {
		db.Close()
	}()

	categoryRepository := storage.NewCategoryRepository(db)
	quizRepository := storage.NewQuizRepository(db)
	translationRepository := storage.NewTranslationRepository(db)

	quizService := quiz.NewQuizService(quizRepository, categoryRepository)
	categoryService := category.NewCategoryService(categoryRepository)
	i18n := translation.NewTranslationService(translationRepository)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Route("/api", func(r chi.Router) {
		r.Get("/get-popular-quizzes", quizService.GetPopularQuizzes)
		r.Get("/get-categories", categoryService.GetCategories)
		r.Get("/get-quizzes-by-category/{id}", quizService.GetQuizzesByCategory)
		r.Get("/translations/{local}", i18n.GetTranslations)
		r.Get("/start-quiz/{id}", quizService.StartQuiz)
		r.Post("/send-result", quizService.SaveResult)
	})
	http.ListenAndServe(":"+strconv.Itoa(cfg.Port), r)
}
