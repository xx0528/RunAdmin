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

type RunNumApi struct {
}

var runNumService = service.ServiceGroupApp.RunPkgServiceGroup.RunNumService

// CreateRunNum 创建RunNum
// @Tags RunNum
// @Summary 创建RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunNum true "创建RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runNum/createRunNum [post]
func (runNumApi *RunNumApi) CreateRunNum(c *gin.Context) {
	var runCreateInfo runPkgReq.RunCreateInfo
	err := c.ShouldBindJSON(&runCreateInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	verify := utils.Rules{
		"State":        {utils.NotEmpty()},
		"OrderName":    {utils.NotEmpty()},
		"EachEnterNum": {utils.NotEmpty()},
		"NumType":      {utils.NotEmpty()},
	}
	if err := utils.Verify(runCreateInfo, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var reErr error
	for _, num := range runCreateInfo.Nums {
		var runNum = runPkg.RunNum{}
		runNum.CreatedBy = utils.GetUserID(c)
		runNum.Num = num
		runNum.State = runCreateInfo.State
		runNum.NumType = runCreateInfo.NumType
		runNum.UserId = runNum.CreatedBy
		runNum.UserNum = runCreateInfo.UserNum
		runNum.OrderName = runCreateInfo.OrderName
		runNum.EachEnterNum = runCreateInfo.EachEnterNum
		runNum.SayHi = runCreateInfo.SayHi
		if runOrder, err := runOrderService.GetRunOrderByName(runNum.OrderName, runNum.UserId); err == nil {
			runNum.PageName = runOrder.PageName
			runNum.PageId = runOrder.PageId
			runCreateInfo.PageId = runOrder.PageId
		}

		if err := runNumService.CreateRunNum(&runNum); err != nil {
			reErr = err
		}

		runOrderNumsService.AddRunNum(runNum)
	}

	if reErr != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(reErr))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}

	UpdateRunPageUsers(runCreateInfo.PageId)
	UpdateRunOrderUsers(runCreateInfo.OrderName, runCreateInfo.UserId)
}

// DeleteRunNum 删除RunNum
// @Tags RunNum
// @Summary 删除RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunNum true "删除RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runNum/deleteRunNum [delete]
func (runNumApi *RunNumApi) DeleteRunNum(c *gin.Context) {
	var runNum runPkg.RunNum
	err := c.ShouldBindJSON(&runNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runNum.DeletedBy = utils.GetUserID(c)
	delRunNum, delErr := runNumService.GetRunNum(runNum.ID)
	if err := runNumService.DeleteRunNum(runNum); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

	if delErr == nil {
		UpdateRunPageUsers(delRunNum.PageId)
		UpdateRunOrderUsers(delRunNum.OrderName, delRunNum.UserId)
		runOrderNumsService.DeleteRunNum(delRunNum)
	}
}

// DeleteRunNumByIds 批量删除RunNum
// @Tags RunNum
// @Summary 批量删除RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /runNum/deleteRunNumByIds [delete]
func (runNumApi *RunNumApi) DeleteRunNumByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	delRunNums, delErr := runNumService.GetRunNumByIds(IDS)

	if err := runNumService.DeleteRunNumByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
	//这里有bug 如果号码来自多个落地页或工单 这里要找到要删除的号码都来自哪个落地页和工单
	if delErr == nil {
		for _, runNum := range delRunNums {
			UpdateRunPageUsers(runNum.PageId)
			UpdateRunOrderUsers(runNum.OrderName, runNum.UserId)
			runOrderNumsService.DeleteRunNum(runNum)
		}
	}
}

// UpdateRunNum 更新RunNum
// @Tags RunNum
// @Summary 更新RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunNum true "更新RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runNum/updateRunNum [put]
func (runNumApi *RunNumApi) UpdateRunNum(c *gin.Context) {
	var runNum runPkg.RunNum
	err := c.ShouldBindJSON(&runNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runNum.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Num":          {utils.NotEmpty()},
		"NumType":      {utils.NotEmpty()},
		"UserId":       {utils.NotEmpty()},
		"EachEnterNum": {utils.NotEmpty()},
	}
	if err := utils.Verify(runNum, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//因为可能要改Num的值 在监控数据里是以Num值来确定的哪一个，所以这里要取到改之前的Num值，只能先查下
	prevNum := ""
	if rerunNum, err := runNumService.GetRunNum(runNum.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		prevNum = rerunNum.Num
	}

	if err := runNumService.UpdateRunNum(runNum); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

	runOrderNumsService.UpdateRunNum(prevNum, runNum)
}

// FindRunNum 用id查询RunNum
// @Tags RunNum
// @Summary 用id查询RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkg.RunNum true "用id查询RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runNum/findRunNum [get]
func (runNumApi *RunNumApi) FindRunNum(c *gin.Context) {
	var runNum runPkg.RunNum
	err := c.ShouldBindQuery(&runNum)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerunNum, err := runNumService.GetRunNum(runNum.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerunNum": rerunNum}, c)
	}
}

// GetRunNumList 分页获取RunNum列表
// @Tags RunNum
// @Summary 分页获取RunNum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkgReq.RunNumSearch true "分页获取RunNum列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runNum/getRunNumList [get]
func (runNumApi *RunNumApi) GetRunNumList(c *gin.Context) {
	var pageInfo runPkgReq.RunNumSearch
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
	if list, total, searchOptions, err := runNumService.GetRunNumInfoList(pageInfo); err != nil {
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

// @Router /runNum/GetRunOrders [get]
func (runNumApi *RunNumApi) GetRunOrders(c *gin.Context) {
	userId := utils.GetUserID(c)
	if orderNames, err := runOrderService.GetRunOrdersByUserId(userId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(orderNames, c)
	}
}

func (runNumApi *RunNumApi) UpdateRunNumByIds(c *gin.Context) {
	var updateInfos runPkgReq.RunNumUpdateByIds
	err := c.ShouldBindJSON(&updateInfos)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	for _, id := range updateInfos.Ids {
		runNum, err := runNumService.GetRunNum(uint(id))
		if err != nil {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
		}
		runNum.UpdatedBy = utils.GetUserID(c)
		runNum.State = updateInfos.State
		runNum.EachEnterNum = updateInfos.EachEnterNum
		runNum.SayHi = updateInfos.SayHi

		if err := runNumService.UpdateRunNum(runNum); err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			response.FailWithMessage("更新失败", c)
			return
		}

		runOrderNumsService.UpdateRunNum(runNum.Num, runNum)
	}

	response.OkWithMessage("更新成功", c)
}
