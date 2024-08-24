package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"         // Драйвер PostgreSQL
	"github.com/jackc/pgx/v5/pgxpool" // Пул соединений для PostgreSQL

	"github.com/MishaAnikutin/metadata-api/internal/model"
)

type MetadataRepo struct {
	session *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *MetadataRepo {
	return &MetadataRepo{
		session: pool,
	}
}

func (repo *MetadataRepo) GetByID(ctx context.Context, ID string) (model.Metadata, error) {

	query := `SELECT id, video_path, title, subtitle, description, tag, channel_id, likes, dislikes
			  FROM video_metadata
			  WHERE id = @VideoID
			  LIMIT 1`

	row := repo.session.QueryRow(ctx, query, pgx.NamedArgs{"VideoID": ID})

	video := model.Metadata{}

	err := row.Scan(
		&video.ID,
		&video.VideoPath,
		&video.Title,
		&video.Subtitle,
		&video.Description,
		&video.Tag,
		&video.ChannelID,
		&video.Likes,
		&video.Dislikes,
	)

	if err != nil {
		return model.Metadata{}, fmt.Errorf("Невозможно сканировать данные: %w", err)
	}

	return video, nil
}

func (repo *MetadataRepo) GetAll(ctx context.Context) ([]model.Metadata, error) {
	query := `SELECT id, video_path, title, subtitle, description, tag, channel_id, likes, dislikes FROM video_metadata`

	rows, err := repo.session.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("Невозможно исполнить запрос: %w", err)
	}
	defer rows.Close()

	var metadataList []model.Metadata

	for rows.Next() {
		var video model.Metadata
		err := rows.Scan(
			&video.ID,
			&video.VideoPath,
			&video.Title,
			&video.Subtitle,
			&video.Description,
			&video.Tag,
			&video.ChannelID,
			&video.Likes,
			&video.Dislikes,
		)
		if err != nil {
			return nil, fmt.Errorf("Невозможно сканировать данные: %w", err)
		}
		metadataList = append(metadataList, video)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return metadataList, nil
}

func (repo *MetadataRepo) GetByTag(ctx context.Context, Tag string) ([]model.Metadata, error) {
	query := `SELECT id, video_path, title, subtitle, description, tag, channel_id, likes, dislikes
			  FROM video_metadata
			  WHERE tag = @Tag`

	rows, err := repo.session.Query(ctx, query, pgx.NamedArgs{"VideoID": Tag})

	if err != nil {
		return nil, fmt.Errorf("Невозможно исполнить запрос: %w", err)
	}
	defer rows.Close()

	var metadataList []model.Metadata

	for rows.Next() {
		var video model.Metadata
		err := rows.Scan(
			&video.ID,
			&video.VideoPath,
			&video.Title,
			&video.Subtitle,
			&video.Description,
			&video.Tag,
			&video.ChannelID,
			&video.Likes,
			&video.Dislikes,
		)
		if err != nil {
			return nil, fmt.Errorf("Невозможно сканировать данные: %w", err)
		}
		metadataList = append(metadataList, video)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return metadataList, nil
}
