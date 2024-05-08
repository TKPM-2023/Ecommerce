package orderstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindOrderWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*ordermodel.Order, error) {
	var data ordermodel.Order
	db := s.db

	db = db.Table(ordermodel.Order{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i], "status = ?", 1)
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
