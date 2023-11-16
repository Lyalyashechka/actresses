package postgre

import (
	"github.com/Lyalyashechka/actresses/internal/tools/logger"
	"gorm.io/gorm"
)

type Postgre struct {
	logger logger.Logger
	db     *gorm.DB
}

func New(logger logger.Logger, db *gorm.DB) Postgre {
	return Postgre{logger: logger, db: db}
}
