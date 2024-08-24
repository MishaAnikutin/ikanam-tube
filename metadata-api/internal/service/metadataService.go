package service

import (
	"context"

	"github.com/MishaAnikutin/metadata-api/internal/model"
	"github.com/MishaAnikutin/metadata-api/internal/repository"
)

type metadataService struct {
	repo *repository.MetadataRepo
}

func (s *metadataService) GetByID(ctx context.Context, ID string) (model.Metadata, error) {
	return s.repo.GetByID(ctx, ID)
}

func (s *metadataService) GetByTag(ctx context.Context, Tag string) ([]model.Metadata, error) {
	return s.repo.GetByTag(ctx, Tag)
}

func (s *metadataService) GetAll(ctx context.Context) ([]model.Metadata, error) {
	return s.repo.GetAll(ctx)
}
