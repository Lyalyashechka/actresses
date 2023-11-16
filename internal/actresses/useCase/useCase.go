package usecase

import (
	"github.com/Lyalyashechka/actresses/internal/actresses"
	"github.com/Lyalyashechka/actresses/internal/tools/logger"
)

type UseCase struct {
	logger              logger.Logger
	actressesRepository actresses.Repository
}

func New(logger logger.Logger, actressesRepository actresses.Repository) UseCase {
	return UseCase{
		logger:              logger,
		actressesRepository: actressesRepository,
	}
}
