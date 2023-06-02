/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-17 11:46:07
 * @Description:
 */
/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-09 10:47:01
 * @Description:
 */
package runPkg

import (
	"fmt"

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

type RunPageApi struct {
}

var runPageService = service.ServiceGroupApp.RunPkgServiceGroup.RunPageService
var runOrderNumsService = service.ServiceGroupApp.RunPkgServiceGroup.RunOrderNumsService

func init() {
	go runOrderNumsService.StartTimer()
}

// CreateRunPage 创建RunPage
// @Tags RunPage
// @Summary 创建RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunPage true "创建RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runPage/createRunPage [post]
func (runPageApi *RunPageApi) CreateRunPage(c *gin.Context) {
	var runPage runPkg.RunPage
	err := c.ShouldBindJSON(&runPage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runPage.CreatedBy = utils.GetUserID(c)
	runPage.State = 1
	runPage.UserId = runPage.CreatedBy

	if haveRunPage, err := runPageService.GetRunPageByName(runPage); err == nil {
		if len(haveRunPage.PageId) > 0 {
			response.FailWithMessage(fmt.Sprintf("落地页名【%s】已存在，请重新命名", haveRunPage.PageName), c)
			return
		}
	}

	runPage.PageId, runPage.Url = runPage.GenerateUrl()

	if err := runPageService.CreateRunPage(&runPage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRunPage 删除RunPage
// @Tags RunPage
// @Summary 删除RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunPage true "删除RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runPage/deleteRunPage [delete]
func (runPageApi *RunPageApi) DeleteRunPage(c *gin.Context) {
	var runPage runPkg.RunPage
	err := c.ShouldBindJSON(&runPage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runPage.DeletedBy = utils.GetUserID(c)
	if err := runPageService.DeleteRunPage(runPage); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRunPageByIds 批量删除RunPage
// @Tags RunPage
// @Summary 批量删除RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /runPage/deleteRunPageByIds [delete]
func (runPageApi *RunPageApi) DeleteRunPageByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := runPageService.DeleteRunPageByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRunPage 更新RunPage
// @Tags RunPage
// @Summary 更新RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunPage true "更新RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runPage/updateRunPage [put]
func (runPageApi *RunPageApi) UpdateRunPage(c *gin.Context) {
	var runPage runPkg.RunPage
	err := c.ShouldBindJSON(&runPage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runPage.UpdatedBy = utils.GetUserID(c)
	if err := runPageService.UpdateRunPage(runPage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func UpdateRunPageUsers(pageId string) {
	runNums, err := runNumService.GetAllRunNumsByPageId(pageId)
	if err == nil {
		runPageInfo := runPkg.RunPage{
			PageId:  pageId,
			UserNum: len(runNums),
		}
		runPageService.UpdateRunPageUsers(runPageInfo)
	}
}

// FindRunPage 用id查询RunPage
// @Tags RunPage
// @Summary 用id查询RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkg.RunPage true "用id查询RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runPage/findRunPage [get]
func (runPageApi *RunPageApi) FindRunPage(c *gin.Context) {
	var runPage runPkg.RunPage
	err := c.ShouldBindQuery(&runPage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerunPage, err := runPageService.GetRunPage(runPage.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerunPage": rerunPage}, c)
	}
}

// GetRunPageList 分页获取RunPage列表
// @Tags RunPage
// @Summary 分页获取RunPage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkgReq.RunPageSearch true "分页获取RunPage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runPage/getRunPageList [get]
func (runPageApi *RunPageApi) GetRunPageList(c *gin.Context) {
	var pageInfo runPkgReq.RunPageSearch
	//没传state 就设置默认值是-1 状态不受限制
	state := c.Query("state")
	if state == "" {
		pageInfo.State = -1
	}
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, searchOptions, err := runPageService.GetRunPageInfoList(pageInfo); err != nil {
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
