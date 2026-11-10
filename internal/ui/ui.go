package ui

import (
	"context"
	"pada/internal/model"
	"pada/internal/service"
)

type WailsApp struct {
	ctx        context.Context
	appService *service.AppService
}

func NewWailsApp(srv *service.AppService) *WailsApp {
	return &WailsApp{appService: srv}
}

func (a *WailsApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

// GetPersonalMetrics — этот метод будет вызван из JS.
// Wails автоматически превратит []model.PersonalMetrics в массив JS-объектов.
func (a *WailsApp) GetPersonalMetrics() ([]model.PersonalMetrics, error) {
	return a.appService.GetPersonalMetrics()
}

// GetGroupReportList — этот метод будет доступен в JS
func (a *WailsApp) GetGroupReportList() ([]model.GroupReport, error) {
	return a.appService.GetGroupReportList()
}
