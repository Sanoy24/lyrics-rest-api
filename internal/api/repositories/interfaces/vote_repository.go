package interfaces

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

type VoteRepository interface {
	// Create or update a vote (upvote/downvote/unvote)
	CastVote(ctx context.Context, vote *models.Vote) error

	// Get a user's vote on a specific entity
	GetUserVote(ctx context.Context, userID int, entityType string, entityID int) (*models.Vote, error)

	// Get entity vote score (sum of vote scores)
	GetEntityVoteScore(ctx context.Context, entityType string, entityID int) (int, error)

	// Count upvotes/downvotes separately (optional)
	CountVotesByType(ctx context.Context, entityType string, entityID uint, value int8) (int64, error)
}
