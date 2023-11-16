package postgre

import (
	"context"

	"github.com/Lyalyashechka/actresses/internal/models"
)

func (r Postgre) SelectAllActresses(ctx context.Context) ([]models.Actress, error) {
	var actresses []models.Actress
	res := r.db.Table(models.Actress{}.TableName()).Find(&actresses)
	if err := res.Error; err != nil {
		return []models.Actress{}, err
	}

	return actresses, nil
}

func (r Postgre) SelectTop3Actresses(ctx context.Context) ([]models.Actress, error) {
	var actresses []models.Actress
	res := r.db.Table(models.Actress{}.TableName()).Order("rating desc").Limit(3).Find(&actresses)
	if err := res.Error; err != nil {
		return []models.Actress{}, err
	}

	return actresses, nil
}
