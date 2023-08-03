package services

import (
	"context"

	"github.com/daxpat2000/cricserver/internal/core/domain"
	"github.com/daxpat2000/cricserver/internal/core/ports"
)

type TeamService struct {
	repository ports.TeamRepository
}

func NewTeamService(repository ports.TeamRepository) *TeamService {
	return &TeamService{
		repository: repository,
	}
}

func (ts *TeamService) CreateTeam(ctx context.Context, team *domain.Team) error {
	return ts.repository.CreateTeam(ctx, team)
}

func (ts *TeamService) GetTeam(ctx context.Context, name, teamType string) (*domain.Team, error) {
	return ts.repository.GetTeam(ctx, name, teamType)
}

func (ts *TeamService) ListTeams(ctx context.Context) ([]domain.Team, error) {
	return ts.repository.ListTeams(ctx)
}
