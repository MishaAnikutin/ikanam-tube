package migrations

import (
	"context"

	v1 "github.com/MishaAnikutin/metadata-api/internal/migrations/v1"
	"github.com/jackc/pgx/v5/pgxpool"
)

func UpgradeHead(ctx context.Context, session *pgxpool.Pool) error {
	return v1.Upgrade(ctx, session)
}
