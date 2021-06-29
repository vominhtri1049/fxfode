package restaurantstorage

import (
	"context"

	"github.com/vominhtri1049/foody/common"
	"github.com/vominhtri1049/foody/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	var result *restaurantmodel.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).Where(conditions)

	if err := db.
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return result, nil
}
