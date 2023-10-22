package repositories

import (
	"context"

	"github.com/1boombacks1/botInsurance/internal/models"
	"github.com/1boombacks1/botInsurance/internal/repositories/pgdb"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	CreateClient(ctx context.Context, client models.Client) (int, error)
	GetClientById(ctx context.Context, id int) (models.Client, error)
}

type Insurer interface {
	CreateInsurer(ctx context.Context, insurer models.Insurer) (int, error)
	GetInsurerById(ctx context.Context, id int) (models.Insurer, error)
}

type Repositories struct {
	Client Client
	// Insurer Insurer
}

func NewRepositories(pgx *pgxpool.Pool) *Repositories {
	return &Repositories{
		Client: pgdb.NewClientRepo(pgx),
	}
}
