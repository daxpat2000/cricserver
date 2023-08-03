package sqlite

import (
	"context"

	"github.com/daxpat2000/cricserver/internal/core/domain"
)

func (sqlite *SQLiteRepository) CreateTeam(ctx context.Context, team *domain.Team) error {
	req := sqlite.db.WithContext(ctx).Create(team)
	if req.RowsAffected == 0 {
		return req.Error
	}

	return nil
}

func (sqlite *SQLiteRepository) GetTeam(ctx context.Context, name string, teamType string) (*domain.Team, error) {
	team := domain.NewTeam(name, teamType)

	req := sqlite.db.WithContext(ctx).First(team)
	if req.RowsAffected == 0 {
		return nil, req.Error
	}

	return team, nil
}

func (sqlite *SQLiteRepository) ListTeams(ctx context.Context) ([]domain.Team, error) {
	var teams []domain.Team

	req := sqlite.db.WithContext(ctx).Find(&teams)
	if req.Error != nil {
		return nil, req.Error
	}

	return teams, nil
}
