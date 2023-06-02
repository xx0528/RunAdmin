/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-11 11:43:34
 * @Description:
 */
/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-04 15:39:06
 * @Description:
 */
package runPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RunPageRouter struct {
}

// InitRunPageRouter 初始化 RunPage 路由信息
func (s *RunPageRouter) InitRunPageRouter(Router *gin.RouterGroup) {
	runPageRouter := Router.Group("runPage").Use(middleware.OperationRecord())
	runPageRouterWithoutRecord := Router.Group("runPage")
	var runPageApi = v1.ApiGroupApp.RunPkgApiGroup.RunPageApi
	{
		runPageRouter.POST("createRunPage", runPageApi.CreateRunPage)             // 新建RunPage
		runPageRouter.DELETE("deleteRunPage", runPageApi.DeleteRunPage)           // 删除RunPage
		runPageRouter.DELETE("deleteRunPageByIds", runPageApi.DeleteRunPageByIds) // 批量删除RunPage
		runPageRouter.PUT("updateRunPage", runPageApi.UpdateRunPage)              // 更新RunPage
	}
	{
		runPageRouterWithoutRecord.GET("findRunPage", runPageApi.FindRunPage)       // 根据ID获取RunPage
		runPageRouterWithoutRecord.GET("getRunPageList", runPageApi.GetRunPageList) // 获取RunPage列表
	}
}

func (s *RunPageRouter) InitLandingPageRouter(Router *gin.RouterGroup) {
	landingRouter := Router.Group("land")
	var runOrderNumsApi = v1.ApiGroupApp.RunPkgApiGroup.RunOrderNumsApi
	{
		landingRouter.GET("get", runOrderNumsApi.RedirectURL)
	}
	landingRouter2 := Router.Group("burl")
	{
		landingRouter2.GET("get", runOrderNumsApi.RedirectURLOld)
	}
}
