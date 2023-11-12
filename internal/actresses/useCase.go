package actresses

import (
	"context"

	"github.com/Lyalyashechka/actresses/internal/models"
)

type UseCase interface {
	GetTopActresses(ctx context.Context) ([]models.Actress, error)
	GetAllActresses(ctx context.Context) ([]models.Actress, error)
	VoteActress(ctx context.Context, vote models.Vote) error
}
