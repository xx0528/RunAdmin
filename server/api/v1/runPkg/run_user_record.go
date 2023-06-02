package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
	runPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/runPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RunUserRecordApi struct {
}

var runUserRecordService = service.ServiceGroupApp.RunPkgServiceGroup.RunUserRecordService

// CreateRunUserRecord 创建RunUserRecord
// @Tags RunUserRecord
// @Summary 创建RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunUserRecord true "创建RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runUserRecord/createRunUserRecord [post]
func (runUserRecordApi *RunUserRecordApi) CreateRunUserRecord(c *gin.Context) {
	var runUserRecord runPkg.RunUserRecord
	err := c.ShouldBindJSON(&runUserRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runUserRecordService.CreateRunUserRecord(&runUserRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRunUserRecord 删除RunUserRecord
// @Tags RunUserRecord
// @Summary 删除RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunUserRecord true "删除RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runUserRecord/deleteRunUserRecord [delete]
func (runUserRecordApi *RunUserRecordApi) DeleteRunUserRecord(c *gin.Context) {
	var runUserRecord runPkg.RunUserRecord
	err := c.ShouldBindJSON(&runUserRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runUserRecordService.DeleteRunUserRecord(runUserRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRunUserRecordByIds 批量删除RunUserRecord
// @Tags RunUserRecord
// @Summary 批量删除RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /runUserRecord/deleteRunUserRecordByIds [delete]
func (runUserRecordApi *RunUserRecordApi) DeleteRunUserRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runUserRecordService.DeleteRunUserRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRunUserRecord 更新RunUserRecord
// @Tags RunUserRecord
// @Summary 更新RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunUserRecord true "更新RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runUserRecord/updateRunUserRecord [put]
func (runUserRecordApi *RunUserRecordApi) UpdateRunUserRecord(c *gin.Context) {
	var runUserRecord runPkg.RunUserRecord
	err := c.ShouldBindJSON(&runUserRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runUserRecordService.UpdateRunUserRecord(runUserRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRunUserRecord 用id查询RunUserRecord
// @Tags RunUserRecord
// @Summary 用id查询RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkg.RunUserRecord true "用id查询RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runUserRecord/findRunUserRecord [get]
func (runUserRecordApi *RunUserRecordApi) FindRunUserRecord(c *gin.Context) {
	var runUserRecord runPkg.RunUserRecord
	err := c.ShouldBindQuery(&runUserRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerunUserRecord, err := runUserRecordService.GetRunUserRecord(runUserRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerunUserRecord": rerunUserRecord}, c)
	}
}

// GetRunUserRecordList 分页获取RunUserRecord列表
// @Tags RunUserRecord
// @Summary 分页获取RunUserRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkgReq.RunUserRecordSearch true "分页获取RunUserRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runUserRecord/getRunUserRecordList [get]
func (runUserRecordApi *RunUserRecordApi) GetRunUserRecordList(c *gin.Context) {
	var pageInfo runPkgReq.RunUserRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, searchOptions, err := runUserRecordService.GetRunUserRecordInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:          list,
			Total:         total,
			SearchOptions: searchOptions,
			Page:          pageInfo.Page,
			PageSize:      pageInfo.PageSize,
		}, "获取成功", c)
	}
}
