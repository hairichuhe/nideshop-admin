package router

import (
	"nideshop-admin/controller/api"
	"nideshop-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("jsonInBlacklist", api.JsonInBlacklist)   //jwt加入黑名单
	}
}