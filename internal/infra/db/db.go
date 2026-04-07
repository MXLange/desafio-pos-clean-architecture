package db

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func NewConnection(ctx context.Context, driver, dsn string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(path string, dbFile string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", absPath),
		fmt.Sprintf("sqlite3://%s", dbFile),
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
