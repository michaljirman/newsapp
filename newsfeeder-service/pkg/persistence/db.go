package persistence

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/michaljirman/newsapp/newsfeeder-service/pkg/configs"
)

func InitPostgresDbWithMigration(cfg *configs.DBConfig, logger *logrus.Logger) (*sql.DB, error) {
	// dbHandle setup
	dbHandle, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open failed")
	}
	dbHandle.SetMaxOpenConns(cfg.MaxOpenConns)
	dbHandle.SetMaxIdleConns(cfg.MaxIdleConns)
	dbHandle.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	dbLogger := logger.WithField("db", "postgresql").Logger

	// Run migrations
	dbLogger.Debug("database migration initiated")

	driver, err := postgres.WithInstance(dbHandle, &postgres.Config{})
	if err != nil {
		dbLogger.Fatalf("could not start sql migration... %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", cfg.MigrationsDirectoryPath), cfg.Name, driver)
	if err != nil {
		dbLogger.Fatalf("migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		dbLogger.Fatalf("An error occurred while syncing the database.. %v", err)
	}

	dbLogger.Debug("database migration completed")

	return dbHandle, nil
}
