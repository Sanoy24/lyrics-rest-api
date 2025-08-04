package interfaces

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/models"
)

// POST		/songs/:songID/annotations			Create a new annotation for a song
// GET		/songs/:songID/annotations			Get all annotations for a specific song
// GET		/annotations/:id					Get a single annotation by its ID
// PUT		/annotations/:id					Update an annotation (by owner or admin)
// DELETE	/annotations/:id					Delete an annotation (by owner or admin)
// POST		/annotations/:id/votes				Upvote or downvote an annotation
// GET		/annotations/:id/comments			Get comments for an annotation
// POST		/annotations/:id/comments			Add comment to an annotation
// GET		/users/:userID/annotations			Get all annotations created by a user
// GET		/songs/:songID/annotations/summary	Get summary info (e.g., count, average votes)
// POST		/annotations/:id/verify				Admin marks annotation as verified
// GET		/annotations/pending				Admin lists unverified annotations

type AnnotationRepository interface {
	CreateAnnotation(ctx context.Context, annotation *models.Annotation) error
	GetAnnotationsBySongID(ctx context.Context, songID uint) ([]models.Annotation, error)
	GetAnnotationByID(ctx context.Context, annotationID uint) (*models.Annotation, error)
	UpdateAnnotation(ctx context.Context, annotationID uint, updates *models.UpdateAnnotationRequest) error
	DeleteAnnotation(ctx context.Context, annotationID uint) error

	AddAnnotationVote(ctx context.Context, vote *models.Vote) error
	AddAnnotationComment(ctx context.Context, comment *models.Comment) error
	GetCommentsForAnnotation(ctx context.Context, annotationID uint) ([]models.Comment, error)

	GetAnnotationsByUser(ctx context.Context, userID uint) ([]models.Annotation, error)
	GetAnnotationSummary(ctx context.Context, songID uint) (*models.AnnotationSummary, error)

	VerifyAnnotation(ctx context.Context, annotationID uint) error
	GetPendingAnnotations(ctx context.Context) ([]models.Annotation, error)
}
