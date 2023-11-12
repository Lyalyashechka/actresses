package usecase

import (
	"context"

	"github.com/Lyalyashechka/actresses/internal/models"
)

func (uc UseCase) GetTopActresses(ctx context.Context) ([]models.Actress, error) {
	return []models.Actress{
		models.Actress{
			Uuid:   "kwa1",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa2",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa3",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
	}, nil
}

func (uc UseCase) GetAllActresses(ctx context.Context) ([]models.Actress, error) {
	return []models.Actress{
		models.Actress{
			Uuid:   "kwa1",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa2",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa3",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa4",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa5",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa6",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa7",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa8",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
		models.Actress{
			Uuid:   "kwa9",
			Name:   "Eva",
			Rating: 5,
			Photo:  "url",
		},
	}, nil
}
