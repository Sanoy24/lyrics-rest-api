package annotation

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type annotationRepo struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewAnnotationRepo(db *gorm.DB, logger *zap.Logger) interfaces.AnnotationRepository {
	return &annotationRepo{
		db:     db,
		logger: logger,
	}
}

// CreateAnnotation creates a new annotation record in the database.
// It takes a context for cancellation/timeout and a pointer to an Annotation model.
// The annotation object will be populated with any database-generated fields (e.g., ID) upon successful creation.
//
// Parameters:
//
//	ctx: The context.Context for controlling the operations lifecycle.
//	annotation: A pointer to the models.Annotation struct containing the data to be created.
//
// Returns:
//
//	An error if the creation fails, otherwise nil.
func (a *annotationRepo) CreateAnnotation(ctx context.Context, annotation *models.Annotation) error {
	if err := a.db.WithContext(ctx).Create(annotation).Error; err != nil {
		return err
	}
	return nil
}

// GetAnnotationsBySongID retrieves all annotations associated with a specific song ID from the database.
// It takes a context for cancellation/timeout and the unique identifier of the song.
//
// Parameters:
//
//	ctx: The context.Context for controlling the operations lifecycle.
//	songID: The unsigned integer ID of the song for which to retrieve annotations.
//
// Returns:
//
//	A slice of models.Annotation containing all found annotations, or an empty slice if none are found.
//	An error if the database query fails, otherwise nil.
func (a *annotationRepo) GetAnnotationsBySongID(ctx context.Context, songID uint) ([]models.Annotation, error) {
	var annotations []models.Annotation
	if err := a.db.WithContext(ctx).Where("song_id = ?", songID).Find(&annotations).Error; err != nil {
		return nil, err
	}
	return annotations, nil
}

func (a *annotationRepo) GetAnnotationByID(ctx context.Context, annotationID uint) (*models.Annotation, error) {
	return nil, nil
}
func (a *annotationRepo) UpdateAnnotation(ctx context.Context, annotationID uint, updates *models.UpdateAnnotationRequest) error {
	return nil
}
func (a *annotationRepo) DeleteAnnotation(ctx context.Context, annotationID uint) error {
	return nil
}

func (a *annotationRepo) AddAnnotationVote(ctx context.Context, vote *models.Vote) error {
	return nil
}
func (a *annotationRepo) AddAnnotationComment(ctx context.Context, comment *models.Comment) error {
	return nil
}
func (a *annotationRepo) GetCommentsForAnnotation(ctx context.Context, annotationID uint) ([]models.Comment, error) {
	return nil, nil
}

func (a *annotationRepo) GetAnnotationsByUser(ctx context.Context, userID uint) ([]models.Annotation, error) {
	return nil, nil
}
func (a *annotationRepo) GetAnnotationSummary(ctx context.Context, songID uint) (*models.AnnotationSummary, error) {
	return nil, nil
}

func (a *annotationRepo) VerifyAnnotation(ctx context.Context, annotationID uint) error {
	return nil
}
func (a *annotationRepo) GetPendingAnnotations(ctx context.Context) ([]models.Annotation, error) {
	return nil, nil
}
