package main

import (
	"embed"
	"log"
	"pada/internal/repository"
	"pada/internal/service"
	"pada/internal/ui"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 1. Слой данных: Инициализация БД
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal("Ошибка БД:", err)
	}
	defer db.Close()

	// 2. Слой бизнес-логики: Создание сервиса (внедряем БД)
	appService := service.NewAppService(db)

	// 3. Слой представления (Wails): Внедряем сервис
	wailsApp := ui.NewWailsApp(appService)

	// 4. Запуск приложения
	err = wails.Run(&options.App{
		Title:  "PADA",
		Width:  1920,
		Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets, // Обязательно остается в main.go
		},
		OnStartup: wailsApp.Startup,
		Bind: []interface{}{
			wailsApp, // Привязываем UI к JS
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
