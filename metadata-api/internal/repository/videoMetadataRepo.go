package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"         // Драйвер PostgreSQL
	"github.com/jackc/pgx/v5/pgxpool" // Пул соединений для PostgreSQL

	"github.com/MishaAnikutin/metadata-api/internal/model"
)

type VideoRepo struct {
	session *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *VideoRepo {
	return &VideoRepo{
		session: pool,
	}
}

func (repo *VideoRepo) GetByID(ctx context.Context, ID string) (*model.Video, error) {

	query := `SELECT video_id, video_url, video_title, video_description, tag, channel_id, likes, dislikes
			  FROM video_metadata
			  WHERE video_id = @VideoID
			  LIMIT 1`

	row := repo.session.QueryRow(ctx, query, pgx.NamedArgs{"VideoID": ID})

	video := model.Video{}

	err := row.Scan(
		&video.ID,
		&video.VideoPath,
		&video.Title,
		&video.Description,
		&video.Tag,
		&video.ChannelID,
		&video.Likes,
		&video.Dislikes,
	)

	if err != nil {
		return nil, fmt.Errorf("Невозможно сканировать данные: %w", err)
	}

	return &video, nil
}

func (repo *VideoRepo) GetAll(ctx context.Context) (*[]model.Video, error) {
	query := `SELECT video_id, video_url, video_title, video_description, tag, channel_id, likes, dislikes FROM video_metadata`

	rows, err := repo.session.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("Невозможно исполнить запрос: %w", err)
	}
	defer rows.Close()

	var metadataList []model.Video

	for rows.Next() {
		var video model.Video
		err := rows.Scan(
			&video.ID,
			&video.VideoPath,
			&video.Title,
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

	return &metadataList, nil
}

func (repo *VideoRepo) GetByTag(ctx context.Context, Tag string) (*[]model.Video, error) {
	query := `SELECT video_id, video_url, video_title, video_description, tag, channel_id, likes, dislikes
			  FROM video_metadata
			  WHERE tag = @Tag`

	rows, err := repo.session.Query(ctx, query, pgx.NamedArgs{"VideoID": Tag})

	if err != nil {
		return nil, fmt.Errorf("Невозможно исполнить запрос: %w", err)
	}
	defer rows.Close()

	var metadataList []model.Video

	for rows.Next() {
		var video model.Video
		err := rows.Scan(
			&video.ID,
			&video.VideoPath,
			&video.Title,
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

	return &metadataList, nil
}
