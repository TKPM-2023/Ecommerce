package uploadstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
)

func (s *sqlStore) CreateImage(ctx context.Context, data *common.Image) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
