package store

import (
	"context"

	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/pairswap-be/src/model"
	"gorm.io/gorm"
)

type HistoryStore struct {
	db *gorm.DB
}

func NewHistoryStore(db *gorm.DB) *HistoryStore {
	return &HistoryStore{db: db}
}

func (s *HistoryStore) Create(ctx context.Context, history *model.History) error {
	if history == nil {
		log.Warn("history is nil")
		return nil
	}

	if err := s.db.WithContext(ctx).Create(history).Error; err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (s *HistoryStore) GetAllByAddress(ctx context.Context, address string) ([]*model.History, error) {
	histories := []*model.History{}

	err := s.db.WithContext(ctx).
		Model(model.History{}).
		Where("address = ?", address).
		Order("created_at DESC").Find(&histories).Error
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return histories, nil
}
