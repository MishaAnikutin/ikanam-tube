package v3

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Upgrade(ctx context.Context, session *pgxpool.Pool) error {
	_, err := session.Exec(ctx, `ALTER TABLE video_metadata ADD COLUMN IF NOT EXISTS picture_url VARCHAR(255);`)
	return err
}

func Downgrade(ctx context.Context, session *pgxpool.Pool) {
	session.Query(ctx, `ALTER TABLE video_metadata DROP COLUMN picture_url`)
}
