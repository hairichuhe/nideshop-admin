package router

import (
	"nideshop-admin/controller/api"
	"nideshop-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWorkflowRouter(Router *gin.RouterGroup) {
	WorkflowRouter := Router.Group("workflow").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		WorkflowRouter.POST("createWorkFlow", api.CreateWorkFlow) // 创建工作流
	}
}
