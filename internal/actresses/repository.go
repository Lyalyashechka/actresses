package actresses

import (
	"context"

	"github.com/Lyalyashechka/actresses/internal/models"
)

type Repository interface {
	SelectAllActresses(ctx context.Context) ([]models.Actress, error)
	SelectTop3Actresses(ctx context.Context) ([]models.Actress, error)
	AddVotes(ctx context.Context, vote models.Vote) error
}
