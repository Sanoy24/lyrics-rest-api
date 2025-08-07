package vote

import (
	"context"
	"time"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VoteRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewVoteRepo(db *gorm.DB, logger *zap.Logger) interfaces.VoteRepository {
	return &VoteRepo{
		db:     db,
		logger: logger,
	}
}

// Create or update a vote (upvote/downvote/unvote)
func (v *VoteRepo) CastVote(ctx context.Context, vote *models.Vote) error {
	vote.CreatedAt = time.Now()
	vote.UpdatedAt = time.Now()

	if err := v.db.WithContext(ctx).Create(&vote).Error; err != nil {
		v.logger.Error("failed to cast vote", zap.Error(err))
		return err
	}
	if vote.Value == 1 {
		v.logger.Info("vote create successfully with an upvote")
	}
	if vote.Value == -1 {
		v.logger.Info("vote create successfully with a downvote")
	}
	return nil
}

// Get a user's vote on a specific entity
func (v *VoteRepo) GetUserVote(ctx context.Context, userID int, entityType string, entityID int) (*models.Vote, error) {
	return nil, nil
}

// Get entity vote score (sum of vote scores)
func (v *VoteRepo) GetEntityVoteScore(ctx context.Context, entityType string, entityID int) (int, error) {
	return 0, nil
}

// Count upvotes/downvotes separately (optional)
func (v *VoteRepo) CountVotesByType(ctx context.Context, entityType string, entityID uint, value int8) (int64, error) {
	return 0, nil
}

func (r *VoteRepo) UpdateVote(ctx context.Context, vote *models.Vote) error {
	return r.db.WithContext(ctx).Save(vote).Error
}

func (r *VoteRepo) DeleteVote(ctx context.Context, voteID uint) error {
	return r.db.WithContext(ctx).Delete(&models.Vote{}, voteID).Error
}
