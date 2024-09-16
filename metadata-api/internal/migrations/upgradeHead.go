package migrations

import (
	"context"

	v1 "github.com/MishaAnikutin/metadata-api/internal/migrations/v1"
	v2 "github.com/MishaAnikutin/metadata-api/internal/migrations/v2"
	v3 "github.com/MishaAnikutin/metadata-api/internal/migrations/v3"

	"github.com/jackc/pgx/v5/pgxpool"
)

func UpgradeHead(ctx context.Context, session *pgxpool.Pool) error {
	if err := v1.Upgrade(ctx, session); err != nil {
		return err
	}

	if err := v2.Upgrade(ctx, session); err != nil {
		return err
	}

	if err := v3.Upgrade(ctx, session); err != nil {
		return err
	}

	return nil
}
