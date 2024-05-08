package orderstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *ordermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]ordermodel.Order, error) {
	var result []ordermodel.Order
	db := s.db

	db = db.Table(ordermodel.Order{}.TableName())

	if f := filter; f != nil {
		if f.Status > 0 {
			db = db.Where("status = ?", f.Status)
		}
		if f.UserId != nil {
			db = db.Where("user_id = ?", f.UserId.GetLocalID())
		}
		if f.OrderStatus >= 0 {
			db = db.Where("order_status = ?", f.OrderStatus)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i], "status = ?", 1)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())

	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))
	}

	if err := db.
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
