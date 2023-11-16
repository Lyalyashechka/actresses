package postgre

import (
	"context"

	"github.com/Lyalyashechka/actresses/internal/models"
	"gorm.io/gorm"
)

func (r Postgre) AddVotes(ctx context.Context, vote models.Vote) error {
	res := r.db.Table(models.Actress{}.TableName()).Where("uuid = ?", vote.Uuid).Update("rating", gorm.Expr("actresses.rating + ?", vote.Rating))
	if err := res.Error; err != nil {
		return err
	}

	return nil
}
