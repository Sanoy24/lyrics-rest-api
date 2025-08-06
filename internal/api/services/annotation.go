package services

import (
	"context"
	"errors"

	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"github.com/Sanoy24/lyrics-rest-api/internal/models"
	"go.uber.org/zap"
)

type AnnotationService struct {
	annotationRepo interfaces.AnnotationRepository
	songRepo       interfaces.SongRepository
	logger         *zap.Logger
}

func NewAnnotationService(annotationRepo interfaces.AnnotationRepository, songRepo interfaces.SongRepository, logger *zap.Logger) *AnnotationService {
	return &AnnotationService{
		annotationRepo: annotationRepo,
		songRepo:       songRepo,
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

func (a *AnnotationService) GetAnnotationsBySongID(ctx context.Context, songID int) ([]models.AnnotationResponse, error) {
	annotations, err := a.annotationRepo.GetAnnotationsBySongID(ctx, songID)
	if err != nil {
		return nil, err
	}
	var response []models.AnnotationResponse
	for _, ann := range annotations {
		res := models.AnnotationResponse{
			ID:         ann.ID,
			StartIndex: ann.StartIndex,
			EndIndex:   ann.EndIndex,
			Fragment:   ann.Fragment,
			Content:    ann.Content,
		}
		if ann.User.ID != 0 {
			res.User = &models.SimpleUser{
				ID:       ann.User.ID,
				Username: ann.User.Username,
			}
		}
		if ann.Song.ID != 0 {
			res.Song = &models.SimpleSong{
				ID:    ann.Song.ID,
				Title: ann.Song.Title,
			}
		}
		response = append(response, res)
	}

	return response, nil
}

func (a *AnnotationService) UpdateAnnotation(ctx context.Context, annotationID, songID int, req *models.UpdateAnnotationRequest) error {

	song, err := a.songRepo.GetSongByID(ctx, songID)
	if err != nil {
		return err
	}

	lyrics := song.Lyrics
	a.logger.Info("lyrics", zap.String("lyrics", lyrics))

	if req.StartIndex < 0 || req.EndIndex > len(lyrics) || req.EndIndex < req.StartIndex {
		return errors.New("invalid start or end index")
	}

	fragment := lyrics[req.StartIndex:req.EndIndex]
	req.Fragment = fragment

	err = a.annotationRepo.UpdateAnnotation(ctx, annotationID, req)
	if err != nil {
		a.logger.Error("failed to update annotation", zap.Error(err))
		return err
	}
	return nil

}

func (a *AnnotationService) DeleteAnnotation(ctx context.Context, annotationID int) error {
	err := a.annotationRepo.DeleteAnnotation(ctx, annotationID)
	if err != nil {
		a.logger.Error("failed to delete annotation", zap.Error(err))
		return err
	}
	return nil
}
