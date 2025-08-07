package annotation

import (
	"context"
	"fmt"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	customerror "github.com/Sanoy24/lyrics-rest-api/pkg/custom_error"
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
func (a *annotationRepo) GetAnnotationsBySongID(ctx context.Context, songID int) ([]models.Annotation, error) {
	var annotations []models.Annotation
	if err := a.db.WithContext(ctx).
		Where("song_id = ?", songID).
		Preload("Song", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, title")
		}).
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username")
		}).Find(&annotations).Error; err != nil {
		return nil, err
	}
	a.logger.Info("annotations fetched successfully", zap.Any("== Annotation Data ==", annotations))
	return annotations, nil
}

// GetAnnotationByID retrieves a single annotation from the database by its ID.
//
// Parameters:
//
//	ctx: The context for the request, enabling cancellation and timeouts.
//	annotationID: The unique identifier of the annotation to retrieve.
//
// Returns:
//
//	*models.Annotation: A pointer to the retrieved Annotation model if found.
//	error: An error if the annotation is not found or if a database error occurs.
//	       Returns gorm.ErrRecordNotFound if no record matches the ID.
func (a *annotationRepo) GetAnnotationByID(ctx context.Context, annotationID uint) (*models.Annotation, error) {
	var annotation models.Annotation
	// Use WithContext to pass the context to the database operation.
	// Where specifies the condition for retrieval (matching ID).
	// First retrieves the first record that matches the condition.
	if err := a.db.WithContext(ctx).Where("id = ?", annotationID).First(&annotation).Error; err != nil {
		return nil, err
	}
	return &annotation, nil
}

// UpdateAnnotation updates an existing annotation in the database.
//
// Parameters:
//
//	ctx: The context for the request, enabling cancellation and timeouts.
//	annotationID: The unique identifier of the annotation to update.
//	updates: A pointer to a models.UpdateAnnotationRequest struct containing
//	         the fields and their new values to be updated. GORM will only
//	         update non-zero/non-empty fields by default.
//
// Returns:
//
//	error: An error if the update operation fails (e.g., database error,
//	       or if the record with the given ID does not exist).
func (a *annotationRepo) UpdateAnnotation(ctx context.Context, annotationID int, updates *models.UpdateAnnotationRequest) error {

	result := a.db.WithContext(ctx).Model(&models.Annotation{}).Where("id = ?", annotationID).Updates(updates)

	if result.Error != nil {
		a.logger.Error("failed to update annotation", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		err := fmt.Errorf("%w: annotation with id %d", customerror.ErrNotFound, annotationID)
		a.logger.Info("Annotation not found with the given id")
		return err
	}

	a.logger.Info("annotation updated successfully", zap.Int("annotation_id", annotationID))
	return nil
}

// DeleteAnnotation deletes an annotation from the database by its ID.
//
// Parameters:
//
//	ctx: The context for the request, enabling cancellation and timeouts.
//	annotationID: The unique identifier of the annotation to delete.
//
// Returns:
//
//	error: An error if the deletion operation fails (e.g., database error,
//	       or if the record with the given ID does not exist).
//	       GORM performs a soft delete by default if the model has a `gorm.DeletedAt` field.
func (a *annotationRepo) DeleteAnnotation(ctx context.Context, annotationID int) error {
	annotation := &models.Annotation{ID: uint(annotationID)}
	if err := a.db.WithContext(ctx).Delete(annotation).Error; err != nil {
		return err
	}
	return nil
}

func (r *annotationRepo) UpdateVoteScore(ctx context.Context, annotationID uint, scoreDelta int) error {
	return r.db.WithContext(ctx).Model(&models.Annotation{}).
		Where("id = ?", annotationID).
		UpdateColumn("vote_score", gorm.Expr("vote_score + ?", scoreDelta)).Error
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
