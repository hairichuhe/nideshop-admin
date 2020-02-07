package router

import (
	"nideshop-admin/controller/api"
	"nideshop-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("casbin").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		BaseRouter.POST("casbinPUpdata", api.CasbinPUpdata)
		BaseRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)

	}
}
