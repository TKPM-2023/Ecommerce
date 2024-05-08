package productstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindProductWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*productmodel.Product, error) {
	var data productmodel.Product
	db := s.db.Table(productmodel.Product{}.TableName())

	var length int64
	if err := db.Count(&length).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if length == 0 {
		return nil, nil
	}

	for i := range moreKeys {
		db.Preload(moreKeys[i], "status = ?", 1)
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
