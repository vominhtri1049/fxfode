package ginrestaurant

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vominhtri1049/foody/common"
	"github.com/vominhtri1049/foody/component"
	"github.com/vominhtri1049/foody/modules/restaurant/restaurantbiz"
	"github.com/vominhtri1049/foody/modules/restaurant/restaurantmodel"
	"github.com/vominhtri1049/foody/modules/restaurant/restaurantstorage"
)

func CreateRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
