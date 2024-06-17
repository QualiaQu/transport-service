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

	// Проверяем и создаем таблицы, если их нет
	tables := map[string]string{
		"cities":          "cities.sql",
		"transport_types": "transport_types.sql",
		"routes":          "routes.sql",
		"bookings":        "bookings.sql",
	}

	for tableName, file := range tables {
		exists, err := tableExists(ctx, db, tableName)
		if err != nil {
			return err
		}
		if !exists {
			filePath := filepath.Join(sqlDumpsPath, file)
			if err := executeSQLFile(ctx, db, filePath); err != nil {
				return err
			}
		}
	}

	// Заполняем тестовыми данными только если таблицы были созданы
	for tableName, file := range tables {
		exists, err := tableExists(ctx, db, tableName)
		if err != nil {
			return err
		}
		if !exists {
			testDataFile := strings.Replace(file, ".sql", "_data.sql", 1)
			filePath := filepath.Join(testDataPath, testDataFile)
			if _, err := os.Stat(filePath); err == nil {
				if err := executeSQLFile(ctx, db, filePath); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func tableExists(ctx context.Context, db *pg.Conn, tableName string) (bool, error) {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_name = '%s'
	);`, tableName)
	err := db.GetContext(ctx, &exists, query)
	if err != nil {
		return false, fmt.Errorf("ошибка проверки существования таблицы %s: %w", tableName, err)
	}
	return exists, nil
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
