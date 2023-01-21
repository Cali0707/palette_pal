package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Cali0707/palette_pal/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	dpgx "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	DB *sqlx.DB
}

func RunMigrationsClean(c config.Config) (err error) {
	databaseString := getDatabaseString(c)

	db, err := sql.Open("pgx", databaseString)
	if err != nil {
		fmt.Printf("Failed to open database: %s", err.Error())
		return
	}
	driver, err := dpgx.WithInstance(db, &dpgx.Config{})
	if err != nil {
		fmt.Printf("Failed to create driver: %s", err.Error())
		return
	}
	m, err := migrate.NewWithDatabaseInstance("file:///go/go/src/palette_pal/db/migrations", c.PostgresDB, driver)
	if err != nil {
		fmt.Printf("Failed to setup migration: %s", err.Error())
		return
	}
	version, dirty, err := m.Version()
	if version == c.MigrationVersion {
		return nil
	}
	if dirty {
		fmt.Printf("Migration is dirty, please fix this")
		return errors.New("dirty migration")
	}
	err = m.Migrate(c.MigrationVersion)
	if err != nil {
		fmt.Printf("Failed to migrate: %s", err.Error())
		return
	}
	return nil
}

func Connect(c config.Config) (db DB, err error) {
	_db, err := sqlx.Connect("pgx", getDatabaseString(c))
	//conn, err := pgx.Connect(ctx, getDatabaseString(c))
	if err != nil {
		fmt.Printf("Unable to connect to database: %s", err.Error())
		return
	}

	db.DB = _db

	return
}

func getDatabaseString(c config.Config) string {
	databaseString := fmt.Sprintf("postgres://%s:%s@database:5432/%s", c.PostgresUser, c.PostgresPassword, c.PostgresDB)
	return databaseString
}
