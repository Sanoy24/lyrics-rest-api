package services

import (
	"context"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
)

type AnnotationService struct {
	annotationRepo interfaces.AnnotationRepository
	songRepo       interfaces.SongRepository
	logger         *zap.Logger
}

func NewAnnotationService(annotationRepo interfaces.AnnotationRepository, logger *zap.Logger) *AnnotationService {
	return &AnnotationService{
		annotationRepo: annotationRepo,
		logger:         logger,
	}
}

func (a *AnnotationService) CreateAnnotation(ctx context.Context, songId, userId int, req *models.CreateAnnotationRequest) (*models.Annotation, error) {
	song, err := a.songRepo.GetSongByID(ctx, songId)
	if err != nil {
		a.logger.Error("failed to get song", zap.Error(err))
		return nil, err
	}
	lyrics := song.Lyrics

	if req.StartIndex < 0 || req.EndIndex > len(lyrics) || req.EndIndex < req.StartIndex {
		a.logger.Error("invalid start or end index", zap.Error(err))
		return nil, err
	}

	fragment := lyrics[req.StartIndex:req.EndIndex]

	annotation := &models.Annotation{
		StartIndex: req.StartIndex,
		EndIndex:   req.EndIndex,
		Content:    req.Content,
		SongID:     uint(songId),
		UserID:     uint(userId),
		Fragment:   fragment,
	}
	if err := a.annotationRepo.CreateAnnotation(ctx, annotation); err != nil {
		a.logger.Error("failed to create annotation", zap.Error(err))
		return nil, err
	}
	a.logger.Info("annotation created successfully", zap.Any("annotation", annotation))
	return annotation, nil
}
