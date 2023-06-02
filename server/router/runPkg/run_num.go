/*
 * @Author: xx
 * @Date: 2023-04-24 18:50:51
 * @LastEditTime: 2023-05-17 13:05:08
 * @Description:
 */
package runPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RunNumRouter struct {
}

// InitRunNumRouter 初始化 RunNum 路由信息
func (s *RunNumRouter) InitRunNumRouter(Router *gin.RouterGroup) {
	runNumRouter := Router.Group("runNum").Use(middleware.OperationRecord())
	runNumRouterWithoutRecord := Router.Group("runNum")
	var runNumApi = v1.ApiGroupApp.RunPkgApiGroup.RunNumApi
	{
		runNumRouter.POST("createRunNum", runNumApi.CreateRunNum)             // 新建RunNum
		runNumRouter.DELETE("deleteRunNum", runNumApi.DeleteRunNum)           // 删除RunNum
		runNumRouter.DELETE("deleteRunNumByIds", runNumApi.DeleteRunNumByIds) // 批量删除RunNum
		runNumRouter.PUT("updateRunNum", runNumApi.UpdateRunNum)              // 更新RunNum
		runNumRouter.PUT("updateRunNumByIds", runNumApi.UpdateRunNumByIds)    // 批量更新RunNum
	}
	{
		runNumRouterWithoutRecord.GET("findRunNum", runNumApi.FindRunNum)       // 根据ID获取RunNum
		runNumRouterWithoutRecord.GET("getRunNumList", runNumApi.GetRunNumList) // 获取RunNum列表
		runNumRouterWithoutRecord.GET("getRunOrders", runNumApi.GetRunOrders)   // 获取RunOrder列表
	}
}
