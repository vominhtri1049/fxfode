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

func ListRestaurant(appContext component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		paging.Fulfill()

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
