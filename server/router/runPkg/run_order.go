/*
 * @Author: xx
 * @Date: 2023-05-09 10:23:48
 * @LastEditTime: 2023-05-11 19:01:41
 * @Description:
 */
package runPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RunOrderRouter struct {
}

// InitRunOrderRouter 初始化 RunOrder 路由信息
func (s *RunOrderRouter) InitRunOrderRouter(Router *gin.RouterGroup) {
	runOrderRouter := Router.Group("runOrder").Use(middleware.OperationRecord())
	runOrderRouterWithoutRecord := Router.Group("runOrder")
	var runOrderApi = v1.ApiGroupApp.RunPkgApiGroup.RunOrderApi
	{
		runOrderRouter.POST("createRunOrder", runOrderApi.CreateRunOrder)             // 新建RunOrder
		runOrderRouter.POST("createRunNums", runOrderApi.CreateRunNums)               // 创建号码
		runOrderRouter.DELETE("deleteRunOrder", runOrderApi.DeleteRunOrder)           // 删除RunOrder
		runOrderRouter.DELETE("deleteRunOrderByIds", runOrderApi.DeleteRunOrderByIds) // 批量删除RunOrder
		runOrderRouter.PUT("updateRunOrder", runOrderApi.UpdateRunOrder)              // 更新RunOrder
	}
	{
		runOrderRouterWithoutRecord.GET("findRunOrder", runOrderApi.FindRunOrder)       // 根据ID获取RunOrder
		runOrderRouterWithoutRecord.GET("getRunOrderList", runOrderApi.GetRunOrderList) // 获取RunOrder列表
		runOrderRouterWithoutRecord.GET("getOrderNums", runOrderApi.GetOrderNums)       // 获取工单号码
		runOrderRouterWithoutRecord.GET("getRunPages", runOrderApi.GetRunPages)       // 获取自己所有落地页名
	}
}
