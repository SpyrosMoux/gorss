package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Conn *gorm.DB

// Init initializes the database connection and runs migrations.
func Init(dsn, dbSchema string, models ...interface{}) error {
	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbSchema + ".",
			SingularTable: false,
		},
	})
	if err != nil {
		return fmt.Errorf("unable to connect to database, %w", err)
	}

	err = createSchemaIfNotExists(dbSchema)
	if err != nil {
		return err
	}

	// Run migrations
	err = Conn.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}

func createSchemaIfNotExists(dbSchema string) error {
	var count int
	Conn.Raw("SELECT COUNT(*) FROM pg_catalog.pg_namespace WHERE nspname = ?", dbSchema).Scan(&count)

	if count > 0 {
		return nil
	}

	err := Conn.Exec(fmt.Sprintf("CREATE SCHEMA %s", dbSchema)).Error
	if err != nil {
		return fmt.Errorf("failed to create schema=%s, err=%v", dbSchema, err)
	}
	return nil
}
