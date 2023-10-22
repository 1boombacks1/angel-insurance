package pgdb

import (
	"context"
	"errors"
	"fmt"

	"github.com/1boombacks1/botInsurance/internal/models"
	"github.com/1boombacks1/botInsurance/internal/repositories/repo_errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ClientRepo struct {
	*pgxpool.Pool
}

func NewClientRepo(pgpool *pgxpool.Pool) *ClientRepo {
	return &ClientRepo{pgpool}
}

func (r *ClientRepo) CreateClient(ctx context.Context, client models.Client) (int, error) {
	query := `insert into clients(
			last_name,first_name,patronymic, phone, link_to_chat,login,password
		) values (
			@lastname,@firstname,@patronymic,@phone,@linkToChat,@login,@password
		) returning id`
	args := pgx.NamedArgs{
		"lastname":   client.LastName,
		"firstname":  client.FirstName,
		"patronymic": client.Patronymic,
		"phone":      client.Phone,
		"linkToChat": client.LinkToChat,
		"login":      client.Login,
		"password":   client.Password,
	}

	var id int
	err := r.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repo_errors.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("ClientRepo.CreateClient - r.QueryRow: %v", err)
	}
	return id, nil
}

func (r *ClientRepo) GetClientById(ctx context.Context, id int) (models.Client, error) {
	query := "select last_name,first_name,patronymic,phone,link_to_chat,login,password from clients where id = $1"

	var client models.Client
	err := r.QueryRow(ctx, query, id).Scan(
		&client.LastName,
		&client.FirstName,
		&client.Patronymic,
		&client.Phone,
		&client.LinkToChat,
		&client.Login,
		&client.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Client{}, repo_errors.ErrNotFound
		}
		return models.Client{}, fmt.Errorf("ClientRepo.GetClientByID - r.QueryRow: %v", err)
	}

	return client, nil
}
