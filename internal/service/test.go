package service

import (
	"database/sql"
	"pada/internal/model"
	"pada/internal/repository"
)

type AppService struct {
	db *sql.DB
}

func NewAppService(db *sql.DB) *AppService {
	return &AppService{db: db}
}

// GetPersonalMetrics делегирует запрос в репозиторий
func (s *AppService) GetPersonalMetrics() ([]model.PersonalMetrics, error) {
	return repository.GetPersonalMetrics(s.db)
}

//

// ... ваш метод GetPersonalMetrics ...

// GetGroupReportList делегирует запрос в репозиторий
func (s *AppService) GetGroupReportList() ([]model.GroupReport, error) {
	return repository.GetGroupReportList(s.db)
}
