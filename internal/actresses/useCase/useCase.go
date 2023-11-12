package usecase

import "github.com/Lyalyashechka/actresses/internal/logger"

type UseCase struct {
	logger logger.Logger
}

func New(logger logger.Logger) UseCase {
	return UseCase{
		logger: logger,
	}
}
