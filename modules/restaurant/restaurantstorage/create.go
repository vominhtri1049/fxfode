package restaurantstorage

import (
	"context"

	"github.com/vominhtri1049/foody/common"
	"github.com/vominhtri1049/foody/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
