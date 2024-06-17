package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"transport-service/internal/db/pg"

	_ "github.com/lib/pq"
)

func InitDB(ctx context.Context, db *pg.Conn) error {
	basePath, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("ошибка получения текущего каталога: %w", err)
	}
	sqlDumpsPath := filepath.Join(basePath, "internal/db/sql_dumps")
	testDataPath := filepath.Join(basePath, "internal/db/test_data")

	sqlFiles := []string{
		"cities.sql",
		"transport_types.sql",
		"routes.sql",
		"bookings.sql",
	}

	for _, file := range sqlFiles {
		filePath := filepath.Join(sqlDumpsPath, file)
		if err := executeSQLFile(ctx, db, filePath); err != nil {
			return err
		}
	}

	testDataFiles := []string{
		"cities_data.sql",
		"transport_types_data.sql",
		"routes_data.sql",
	}

	for _, file := range testDataFiles {
		filePath := filepath.Join(testDataPath, file)
		if err := executeSQLFile(ctx, db, filePath); err != nil {
			return err
		}
	}

	return nil
}

func executeSQLFile(ctx context.Context, db *pg.Conn, filepath string) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("ошибка чтения файла %s: %w", filepath, err)
	}

	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		if _, err := db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("ошибка выполнения запроса из файла %s: %w", filepath, err)
		}
	}

	return nil
}
