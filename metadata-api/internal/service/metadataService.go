package service

import (
	"context"

	"github.com/MishaAnikutin/metadata-api/internal/model"
	"github.com/MishaAnikutin/metadata-api/internal/repository"
)

type MetadataService struct {
	repo *repository.VideoRepo
}

func New(repo *repository.VideoRepo) *MetadataService {
	return &MetadataService{repo: repo}
}

func (s *MetadataService) GetByID(ctx context.Context, ID string) (*model.Video, error) {
	return s.repo.GetByID(ctx, ID)
}

func (s *MetadataService) GetByTag(ctx context.Context, Tag string) (*[]model.Video, error) {
	return s.repo.GetByTag(ctx, Tag)
}

func (s *MetadataService) GetAll(ctx context.Context) (*[]model.Video, error) {
	return s.repo.GetAll(ctx)
}
