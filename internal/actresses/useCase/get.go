package usecase

import (
	"context"

	"github.com/Lyalyashechka/actresses/internal/models"
)

func (uc UseCase) GetTopActresses(ctx context.Context) ([]models.Actress, error) {
	return uc.actressesRepository.SelectTop3Actresses(ctx)
}

func (uc UseCase) GetAllActresses(ctx context.Context) ([]models.Actress, error) {
	return uc.actressesRepository.SelectAllActresses(ctx)
}
