package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RunUserRecordRouter struct {
}

// InitRunUserRecordRouter 初始化 RunUserRecord 路由信息
func (s *RunUserRecordRouter) InitRunUserRecordRouter(Router *gin.RouterGroup) {
	runUserRecordRouter := Router.Group("runUserRecord").Use(middleware.OperationRecord())
	runUserRecordRouterWithoutRecord := Router.Group("runUserRecord")
	var runUserRecordApi = v1.ApiGroupApp.RunPkgApiGroup.RunUserRecordApi
	{
		runUserRecordRouter.POST("createRunUserRecord", runUserRecordApi.CreateRunUserRecord)   // 新建RunUserRecord
		runUserRecordRouter.DELETE("deleteRunUserRecord", runUserRecordApi.DeleteRunUserRecord) // 删除RunUserRecord
		runUserRecordRouter.DELETE("deleteRunUserRecordByIds", runUserRecordApi.DeleteRunUserRecordByIds) // 批量删除RunUserRecord
		runUserRecordRouter.PUT("updateRunUserRecord", runUserRecordApi.UpdateRunUserRecord)    // 更新RunUserRecord
	}
	{
		runUserRecordRouterWithoutRecord.GET("findRunUserRecord", runUserRecordApi.FindRunUserRecord)        // 根据ID获取RunUserRecord
		runUserRecordRouterWithoutRecord.GET("getRunUserRecordList", runUserRecordApi.GetRunUserRecordList)  // 获取RunUserRecord列表
	}
}
