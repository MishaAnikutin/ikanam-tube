package di

import (
	"context"

	"github.com/MishaAnikutin/metadata-api/internal/repository"
	"github.com/MishaAnikutin/metadata-api/internal/service"
	"github.com/MishaAnikutin/metadata-api/internal/transport/rest"
	v1 "github.com/MishaAnikutin/metadata-api/internal/transport/rest/v1"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Inject(ctx context.Context, pool *pgxpool.Pool) *gin.Engine {

	VideoRepo := repository.New(pool)

	VideoService := service.New(VideoRepo)

	Handlers := v1.New(VideoService)

	router := rest.SetupRouter(Handlers)

	return router
}
