package sqlite

import (
	"context"

	"github.com/daxpat2000/cricserver/internal/core/domain"
	"gorm.io/gorm"
)

func (sqlite *SQLiteRepository) CreatePlayer(ctx context.Context, player *domain.Player) error {
	req := sqlite.db.WithContext(ctx).Create(player)
	if req.RowsAffected == 0 {
		return req.Error
	}

	return nil
}

func (sqlite *SQLiteRepository) GetPlayer(ctx context.Context, id uint) (*domain.Player, error) {
	player := &domain.Player{ID: id}

	req := sqlite.db.WithContext(ctx).First(&player)
	if req.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return player, nil
}

func (sqlite *SQLiteRepository) ListPlayers(ctx context.Context) ([]domain.Player, error) {
	var players []domain.Player

	req := sqlite.db.WithContext(ctx).Find(&players)
	if req.Error != nil {
		return nil, req.Error
	}

	return players, nil
}
