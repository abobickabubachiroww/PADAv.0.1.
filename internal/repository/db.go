package repository

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // или ваш драйвер
)

// InitDB объединяет открытие и миграцию
func InitDB() (*sql.DB, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}

	if err := Migrate(db); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("БД готова:", DbPath())
	return db, nil
}

// dbPath возвращает путь к файлу БД рядом с .exe или в %APPDATA%.
func DbPath() string {
	// Рядом с приложением (портативный режим)
	exe, err := os.Executable()
	if err == nil {
		return filepath.Join(filepath.Dir(exe), "app.db")
	}
	return "app.db"
}

func OpenDB() (*sql.DB, error) {
	path := DbPath() // путь к бд

	// Параметры через DSN:
	// _pragma=busy_timeout(5000) — ждать до 5с, если БД занята
	// _pragma=journal_mode(WAL)   — режим WAL: читатели не блокируют писателя
	// _pragma=foreign_keys(1)     — включить внешние ключи (в SQLite они выключены по умолчанию!)
	dsn := fmt.Sprintf("file:%s?_pragma=busy_timeout(5000)&_pragma=journal_mode(WAL)&_pragma=foreign_keys(1)", path)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	// SQLite — один писатель за раз. Ограничиваем пул, чтобы избежать "database is locked".
	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
