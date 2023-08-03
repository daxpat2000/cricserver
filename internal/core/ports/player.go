package ports

import (
	"context"

	"github.com/daxpat2000/cricserver/internal/core/domain"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, player *domain.Player) error

	GetPlayer(ctx context.Context, id uint) (*domain.Player, error)

	ListPlayers(ctx context.Context) ([]domain.Player, error)
}

type PlayerService interface {
	CreatePlayer(ctx context.Context, player *domain.Player) error

	GetPlayer(ctx context.Context, id uint) (*domain.Player, error)

	ListPlayers(ctx context.Context) ([]domain.Player, error)
}
