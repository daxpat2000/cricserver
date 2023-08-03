package services

import (
	"context"

	"github.com/daxpat2000/cricserver/internal/core/domain"
	"github.com/daxpat2000/cricserver/internal/core/ports"
)

type PlayerService struct {
	repository ports.PlayerRepository
}

func NewPlayerService(repository ports.PlayerRepository) *PlayerService {
	return &PlayerService{
		repository: repository,
	}
}

func (ps *PlayerService) CreatePlayer(ctx context.Context, player *domain.Player) error {
	return ps.repository.CreatePlayer(ctx, player)
}

func (ps *PlayerService) GetPlayer(ctx context.Context, id uint) (*domain.Player, error) {
	return ps.repository.GetPlayer(ctx, id)
}

func (ps *PlayerService) ListPlayers(ctx context.Context) ([]domain.Player, error) {
	return ps.repository.ListPlayers(ctx)
}
