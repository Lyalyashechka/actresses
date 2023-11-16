package postgre

import (
	"fmt"

	"github.com/Lyalyashechka/actresses/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config config.Postgres) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
