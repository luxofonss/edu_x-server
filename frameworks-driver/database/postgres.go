package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connection implements DatabaseConnection for PostgresSQL.
type PostgresDatabase struct {
	DB *gorm.DB
}

// Connect connects to a PostgreSQL database and returns a GORM DB instance.
func (PostgresDatabase) Connect(config DBConfig) (DBConnection, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dialector := postgres.Open(fmt.Sprintf(dsn, config.Host, config.User, config.Password, config.DB, config.Port, config.SSLMode, config.TimeZone))
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, err
	}

	fmt.Println("You successfully connected to Postgres!")

	return db, nil
}
