package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type VoteService struct {
	voteRepo       interfaces.VoteRepository
	songRepo       interfaces.SongRepository
	commentRepo    interfaces.CommentRepository
	annotationRepo interfaces.AnnotationRepository
	logger         *zap.Logger
}

func NewVoteService(voteRepo interfaces.VoteRepository, logger *zap.Logger) *VoteService {
	return &VoteService{
		voteRepo: voteRepo,
		logger:   logger,
	}
}

func (v *VoteService) CastVote(ctx context.Context, req *models.CreateVoteRequest, id int) error {

	// Convert vote type to numerical value
	var voteValue int8
	switch req.VoteType {
	case "upvote":
		voteValue = 1
	case "downvote":
		voteValue = -1
	case "unvote":
		voteValue = 0
	default:
		return fmt.Errorf("invalid vote type: %s", req.VoteType)
	}

	vote := &models.Vote{
		UserID:     req.UserID,
		EntityType: req.EntityType,
		EntityID:   req.EntityID,
		VoteType:   req.VoteType,
		Value:      voteValue,
	}

	// Set the appropriate foreign key based on entity type
	switch req.EntityType {
	case "song":
		songID := uint(id)
		vote.SongID = &songID
	case "annotation": // Fixed typo
		annotationID := uint(id)
		vote.AnnotationID = &annotationID
	case "comment":
		commentID := uint(id)
		vote.CommentID = &commentID
	default:
		return fmt.Errorf("invalid entity type: %s", req.EntityType)
	}

	// Check if user has already voted on this entity
	existingVote, err := v.voteRepo.GetUserVote(ctx, int(req.UserID), req.EntityType, int(req.EntityID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check existing vote: %w", err)
	}

	var scoreDelta int8 = 0

	if existingVote != nil {
		// User has already voted, calculate the score change
		scoreDelta = voteValue - existingVote.Value

		if req.VoteType == "unvote" {
			// Remove the vote
			if err := v.voteRepo.DeleteVote(ctx, existingVote.ID); err != nil {
				return fmt.Errorf("failed to delete vote: %w", err)
			}
		} else {
			// Update existing vote
			existingVote.VoteType = req.VoteType
			existingVote.Value = voteValue
			if err := v.voteRepo.UpdateVote(ctx, existingVote); err != nil {
				return fmt.Errorf("failed to update vote: %w", err)
			}
		}
	} else {
		// New vote
		if req.VoteType != "unvote" {
			scoreDelta = voteValue
			if err := v.voteRepo.CreateVote(ctx, vote); err != nil {
				return fmt.Errorf("failed to create vote: %w", err)
			}
		}
		// If it's unvote and no existing vote, do nothing
	}

	// Update the vote score on the target entity if there's a change
	if scoreDelta != 0 {
		if err := v.updateEntityVoteScore(ctx, req.EntityType, req.EntityID, int(scoreDelta)); err != nil {
			return fmt.Errorf("failed to update entity vote score: %w", err)
		}
	}

	return nil

}

func (v *VoteService) updateEntityVoteScore(ctx context.Context, entityType string, entityID uint, scoreDelta int) error {
	switch entityType {
	case "annotation":
		return v.annotationRepo.UpdateVoteScore(ctx, entityID, scoreDelta)
	case "song":
		return v.songRepo.UpdateVoteScore(ctx, entityID, scoreDelta)
	case "comment":
		return v.commentRepo.UpdateVoteScore(ctx, entityID, scoreDelta)
	default:
		return fmt.Errorf("unsupported entity type for vote score update: %s", entityType)
	}
}
