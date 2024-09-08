package v2

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Upgrade(ctx context.Context, session *pgxpool.Pool) error {
	_, err := session.Exec(ctx, `
		ALTER TABLE video_metadata ADD COLUMN IF NOT EXISTS tag VARCHAR(255);
		ALTER TABLE video_metadata ADD COLUMN IF NOT EXISTS likes int DEFAULT 0;
		ALTER TABLE video_metadata ADD COLUMN IF NOT EXISTS dislikes int DEFAULT 0;
	
	`)
	return err
}

func Downgrade(ctx context.Context, session *pgxpool.Pool) {
	session.Query(ctx, `
		ALTER TABLE video_metadata DROP COLUMN tag
		ALTER TABLE video_metadata DROP COLUMN likes
		ALTER TABLE video_metadata DROP COLUMN dislikes;
	`)
}
