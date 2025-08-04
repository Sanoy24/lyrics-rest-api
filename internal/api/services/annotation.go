package services

import (
	"github.com/Sanoy24/lyrics-rest-api/internal/api/repositories/interfaces"
	"go.uber.org/zap"
)

type AnnotationService struct {
	annotationRepo interfaces.AnnotationRepository
	logger         *zap.Logger
}

func NewAnnotationService(annotationRepo interfaces.AnnotationRepository, logger *zap.Logger) *AnnotationService {
	return &AnnotationService{
		annotationRepo: annotationRepo,
		logger:         logger,
	}
}
