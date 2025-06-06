package db

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect(models ...interface{}) error {
	var err error
	Conn, err = gorm.Open(sqlite.Open("./db.sqlite"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("unable to connect to database, err=%w", err)
	}

	slog.Info("connected to database")

	err = Conn.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("failed to migrate models err=%v", err)
	}
	return nil
}
