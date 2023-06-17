package user

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorytransport/gincategory"
	"TKPM-Go/module/product/producttransport/ginproduct"
	"TKPM-Go/module/upload/uploadtransport/ginupload"
	"TKPM-Go/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func UserRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.Login(appContext))

	//upload service
	v1.POST("/upload", ginupload.Upload(appContext))

	//Category
	category := v1.Group("/categories")
	category.GET("/:id", gincategory.GetCategory(appContext))
	category.GET("/", gincategory.ListCategory(appContext))

	//product
	product := v1.Group("/products")
	product.GET("/:id", ginproduct.GetProduct(appContext))
	product.GET("/", ginproduct.ListProduct(appContext))
}
