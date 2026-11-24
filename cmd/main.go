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

var assets embed.FS

func main() {
	// Инициализация БД
	db, err := repository.InitDB()
	if err != nil {
		log.Fatal("Ошибка БД:", err)
	}
	defer db.Close()

	// Слой бизнес-логики: Создание сервиса (внедряем БД)
	appService := service.NewAppService(db)

	// Слой представления (Wails): Внедряем сервис
	wailsApp := ui.NewWailsApp(appService)

	// 4. Запуск приложения
	err = wails.Run(&options.App{
		Title:  "PADA",
		Width:  1920,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
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
