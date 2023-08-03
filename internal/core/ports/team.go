package ports

import (
	"context"

	"github.com/daxpat2000/cricserver/internal/core/domain"
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, team *domain.Team) error

	GetTeam(ctx context.Context, name string, teamType string) (*domain.Team, error)

	ListTeams(ctx context.Context) ([]domain.Team, error)
}

type TeamService interface {
	CreateTeam(ctx context.Context, team *domain.Team) error

	GetTeam(ctx context.Context, name string, teamType string) (*domain.Team, error)

	ListTeams(ctx context.Context) ([]domain.Team, error)
}
