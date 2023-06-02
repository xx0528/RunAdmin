/*
 * @Author: xx
 * @Date: 2023-05-17 15:44:34
 * @LastEditTime: 2023-05-18 16:06:32
 * @Description:
 */
package runPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RunTplRouter struct {
}

// InitRunTplRouter 初始化 RunTpl 路由信息
func (s *RunTplRouter) InitRunTplRouter(Router *gin.RouterGroup) {
	runTplRouter := Router.Group("runTpl").Use(middleware.OperationRecord())
	runTplRouterWithoutRecord := Router.Group("runTpl")
	var runTplApi = v1.ApiGroupApp.RunPkgApiGroup.RunTplApi
	{
		runTplRouter.POST("createRunTpl", runTplApi.CreateRunTpl)             // 新建RunTpl
		runTplRouter.POST("copyRunTpl", runTplApi.CopyRunTpl)                 // 复制RunTpl
		runTplRouter.DELETE("deleteRunTpl", runTplApi.DeleteRunTpl)           // 删除RunTpl
		runTplRouter.DELETE("deleteRunTplByIds", runTplApi.DeleteRunTplByIds) // 批量删除RunTpl
		runTplRouter.PUT("updateRunTpl", runTplApi.UpdateRunTpl)              // 更新RunTpl
	}
	{
		runTplRouterWithoutRecord.GET("findRunTpl", runTplApi.FindRunTpl)       // 根据ID获取RunTpl
		runTplRouterWithoutRecord.GET("getRunTplList", runTplApi.GetRunTplList) // 获取RunTpl列表
	}
}

func (s *RunTplRouter) InitTplPageRouter(Router *gin.RouterGroup) {
	tplRouter := Router.Group("tpl")
	var runTplApi = v1.ApiGroupApp.RunPkgApiGroup.RunTplApi
	{
		tplRouter.GET("get", runTplApi.GetTplPage)
	}
}
