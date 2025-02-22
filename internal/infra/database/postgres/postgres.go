package postgres

import (
	"database/sql"
	"fmt"
	"go-boilerplate/config"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connect() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		config.Config.DBUser,
		config.Config.DBPassword,
		config.Config.DBName,
		config.Config.DBHost,
		config.Config.DBPort,
		"disable",
	)

	fmt.Println(connectionString)

	db, err := sql.Open(config.Config.DBDriver, connectionString)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(0)
	fmt.Println(db.Stats().MaxIdleClosed)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m := &migrate.Migrate{}
	m, err = migrate.NewWithDatabaseInstance("file://../internal/infra/database/postgres/migrations", config.Config.DBDriver, driver)
	if err != nil {
		m, err = migrate.NewWithDatabaseInstance("file:///internal/infra/database/postgres/migrations", config.Config.DBDriver, driver)
		if err != nil {
			panic(err)
		}
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Printf("DATABASE MIGRATIONS: %s", err.Error())
		panic(err)
	}
}
