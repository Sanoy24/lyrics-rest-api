package interfaces

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

type VoteRepository interface {
	SaveVote(ctx context.Context, vote *models.Vote) error
}
