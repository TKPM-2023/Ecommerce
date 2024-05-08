package gincart

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartbiz"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartmodel"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartstorage"
	"net/http"
)

func RemoveProducts(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		pubsub := ctx.GetPubSub()
		var data cartmodel.RemoveCartProducts
		requester := context.MustGet(common.CurrentUser).(common.Requester)

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewRemoveProductsBusiness(store, pubsub)
		if err := business.RemoveProducts(context.Request.Context(), requester.GetCartId(), data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
