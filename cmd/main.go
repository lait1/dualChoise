package main

import (
	"dualChoose/internal/category"
	"dualChoose/internal/common"
	"dualChoose/internal/config"
	"dualChoose/internal/quiz"
	"dualChoose/internal/storage"
	"github.com/go-chi/chi/v5"
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
			zap.String("connect", common.SafeMysqlCreds(cfg.MySQLConnect)))
	}

	defer func() {
		db.Close()
	}()

	categoryRepository := storage.NewCategoryRepository(db)
	quizRepository := storage.NewQuizRepository(db)

	quizService := quiz.NewQuizService(quizRepository)
	categoryService := category.NewCategoryService(categoryRepository)

	r := chi.NewRouter()
	r.Get("/get-popular-quizzes", quizService.GetPopularQuizzes)
	r.Get("/get-categories", categoryService.GetCategories)
	r.Get("/get-quizzes-by-category/{id}", quizService.GetQuizzesByCategory)
	r.Get("/start-quiz/{id}", quizService.StartQuiz)
	//r.Post("/send-result", quizService.SaveResult)

	http.ListenAndServe(":"+strconv.Itoa(cfg.Port), r)
}
