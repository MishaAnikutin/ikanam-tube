package v1

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Upgrade(ctx context.Context, session *pgxpool.Pool) error {
	_, err := session.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS channels (
			channel_id           SERIAL PRIMARY KEY,
			channel_name         VARCHAR(100) NOT NULL,
			channel_description  TEXT,
			channel_owner        VARCHAR(100) NOT NULL,
			created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		
		CREATE TABLE IF NOT EXISTS video_metadata (
			video_id           SERIAL PRIMARY KEY,
			video_title        VARCHAR(255) NOT NULL,
			video_description  TEXT,
			video_url          VARCHAR(255) NOT NULL,
			channel_id         INT NOT NULL,
			upload_date        TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (channel_id) REFERENCES channels(channel_id)
		);
	`)
	return err
}

func Downgrade(ctx context.Context, session *pgxpool.Pool) {
	session.Query(ctx, `
		DROP TABLE video_metadata;
		DROP TABLE channels;
	`)
}
