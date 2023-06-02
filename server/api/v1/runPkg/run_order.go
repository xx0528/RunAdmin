package runPkg

import (
	"fmt"
	"strings"
	"time"

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

type RunOrderApi struct {
}

var runOrderService = service.ServiceGroupApp.RunPkgServiceGroup.RunOrderService

// CreateRunOrder 创建RunOrder
// @Tags RunOrder
// @Summary 创建RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunOrder true "创建RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runOrder/createRunOrder [post]
func (runOrderApi *RunOrderApi) CreateRunOrder(c *gin.Context) {
	var runOrder runPkg.RunOrder
	err := c.ShouldBindJSON(&runOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runOrder.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"OrderName":    {utils.NotEmpty()},
		"PageId":       {utils.NotEmpty()},
		"MaxEnterNum":  {utils.NotEmpty()},
		"EachEnterNum": {utils.NotEmpty()},
	}
	if err := utils.Verify(runOrder, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runOrderService.CreateRunOrder(&runOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRunOrder 删除RunOrder
// @Tags RunOrder
// @Summary 删除RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunOrder true "删除RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runOrder/deleteRunOrder [delete]
func (runOrderApi *RunOrderApi) DeleteRunOrder(c *gin.Context) {
	var runOrder runPkg.RunOrder
	err := c.ShouldBindJSON(&runOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runOrder.DeletedBy = utils.GetUserID(c)
	runOrder.UserId = runOrder.DeletedBy
	if err := runOrderService.DeleteRunOrder(runOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

	if err := runNumService.DeleteRunNumByOrder(runOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}

	runOrderNumsService.DeleteRunOrder(runOrder)
	UpdateRunPageUsers(runOrder.PageId)
}

// DeleteRunOrderByIds 批量删除RunOrder
// @Tags RunOrder
// @Summary 批量删除RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /runOrder/deleteRunOrderByIds [delete]
func (runOrderApi *RunOrderApi) DeleteRunOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)

	results, _ := runOrderService.GetRunOrderByIds(IDS)
	for _, runOrder := range results {
		runOrder.DeletedBy = deletedBy
		runNumService.DeleteRunNumByOrder(runOrder)
		runOrderNumsService.DeleteRunOrder(runOrder)
		UpdateRunPageUsers(runOrder.PageId)
	}

	if err := runOrderService.DeleteRunOrderByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRunOrder 更新RunOrder
// @Tags RunOrder
// @Summary 更新RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunOrder true "更新RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runOrder/updateRunOrder [put]
func (runOrderApi *RunOrderApi) UpdateRunOrder(c *gin.Context) {
	var runOrder runPkg.RunOrder
	err := c.ShouldBindJSON(&runOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runOrder.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"OrderName":    {utils.NotEmpty()},
		"PageId":       {utils.NotEmpty()},
		"MaxEnterNum":  {utils.NotEmpty()},
		"EachEnterNum": {utils.NotEmpty()},
	}
	if err := utils.Verify(runOrder, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runOrderService.UpdateRunOrder(runOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

	runOrderNumsService.UpdateRunOrder(runOrder)
}

// FindRunOrder 用id查询RunOrder
// @Tags RunOrder
// @Summary 用id查询RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkg.RunOrder true "用id查询RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runOrder/findRunOrder [get]
func (runOrderApi *RunOrderApi) FindRunOrder(c *gin.Context) {
	var runOrder runPkg.RunOrder
	err := c.ShouldBindQuery(&runOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerunOrder, err := runOrderService.GetRunOrder(runOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerunOrder": rerunOrder}, c)
	}
}

// GetRunOrderList 分页获取RunOrder列表
// @Tags RunOrder
// @Summary 分页获取RunOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkgReq.RunOrderSearch true "分页获取RunOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runOrder/getRunOrderList [get]
func (runOrderApi *RunOrderApi) GetRunOrderList(c *gin.Context) {
	var pageInfo runPkgReq.RunOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.UserId = utils.GetUserID(c)
	if list, total, searchOptions, err := runOrderService.GetRunOrderInfoList(pageInfo); err != nil {
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

// 获取自己的所有落地页
func (runOrderApi *RunOrderApi) GetRunPages(c *gin.Context) {
	userId := utils.GetUserID(c)
	if pageNames, err := runPageService.GetRunPageNamesByUserId(userId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(pageNames, c)
	}
}

// 获取工单号码
func (runOrderApi *RunOrderApi) GetOrderNums(c *gin.Context) {
	var orderInfo runPkgReq.RunOrderUrlInfo
	var orderData map[string]interface{}
	err := c.ShouldBindQuery(&orderInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	orderType := ""
	//云控工单
	if strings.Contains(orderInfo.Url, "share/share") {
		orderType = runPkg.OrderType_Share
		//007工单
	} else if strings.Contains(orderInfo.Url, "kf.007.tools") {
		orderType = runPkg.OrderType_KF007
		//goo.su工单
	} else if strings.Contains(orderInfo.Url, "//goo.su") {
		orderType = runPkg.OrderType_GooSu
		//ok0.xyz工单
	} else if strings.Contains(orderInfo.Url, ".xyz") {
		orderType = runPkg.OrderType_XYZ
		//url工单
	} else if strings.Contains(orderInfo.Url, "://url") {
		orderType = runPkg.OrderType_URL
	}
	if len(orderType) <= 0 {
		response.FailWithMessage(fmt.Sprint("没有找到工单链接类型 url--", orderInfo.Url), c)
		return
	}

	resultChan := make(chan map[string]interface{}, 1) // 定义函数内部缓冲通道

	// 将耗时操作放到goroutine中执行
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// 异常处理逻辑
				fmt.Println(r)
				global.GVA_LOG.Error("获取工单号异常------------------------")
				resultChan <- map[string]interface{}{
					"error": "获取工单号码失败，请稍后再试",
				}
			}
		}()

		orderData, err = runOrderNumsService.GetOrderData(orderInfo.Url, orderInfo.Psw, orderType, "firstStep")
		if err != nil {
			// response.FailWithMessage(err.Error(), c)
			// return
			panic(err.Error())
		}

		// 将返回结果发送到指定的channel中
		resultChan <- orderData
	}()

	select {
	case <-time.After(time.Second * 8): // 等待5秒超时
		response.FailWithMessage("获取工单号码超时，请稍后再试", c)
	case result := <-resultChan: // 成功获取工单号码
		if errMsg, ok := result["error"]; ok {
			response.FailWithMessage(errMsg.(string), c)
		} else {
			response.OkWithData(result, c)
		}
		// response.OkWithData(result, c)
	}
}

// 创建号码
func (runOrderApi *RunOrderApi) CreateRunNums(c *gin.Context) {
	var createNumsInfo runPkgReq.RunCreateNumsInfo

	err := c.ShouldBindJSON(&createNumsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	createNumsInfo.UserId = utils.GetUserID(c)

	if haveRunOrder, err := runOrderService.GetRunOrderByName(createNumsInfo.OrderName, createNumsInfo.UserId); err == nil {
		if len(haveRunOrder.OrderName) > 0 {
			response.FailWithMessage(fmt.Sprintf("工单名【%s】已存在，请重新命名", haveRunOrder.OrderName), c)
			return
		}
	}

	runPage, err := runPageService.GetRunPageByPageId(createNumsInfo.PageId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//保存需要监控的号码
	createNumsInfo.PageId = runPage.PageId
	if err := runOrderNumsService.SaveOrderNums(createNumsInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//创建工单号码
	if err := runNumService.CreateRunOrderNums(createNumsInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//创建工单
	runOrder := runPkg.RunOrder{
		CreatedBy:    utils.GetUserID(c),
		PageId:       createNumsInfo.PageId,
		UserId:       createNumsInfo.UserId,
		PageName:     createNumsInfo.PageName,
		OrderName:    createNumsInfo.OrderName,
		MaxEnterNum:  createNumsInfo.MaxEnterNum,
		EachEnterNum: createNumsInfo.EachEnterNum,
		UserNum:      len(createNumsInfo.Nums),
	}
	verify := utils.Rules{
		"OrderName":    {utils.NotEmpty()},
		"PageId":       {utils.NotEmpty()},
		"MaxEnterNum":  {utils.NotEmpty()},
		"EachEnterNum": {utils.NotEmpty()},
	}
	if err := utils.Verify(runOrder, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runOrderService.CreateRunOrder(&runOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	UpdateRunPageUsers(runPage.PageId)

	response.OkWithMessage("添加成功", c)
}

func UpdateRunOrderUsers(orderName string, userId uint) {
	runNums, err := runNumService.GetAllRunNumsByOrderName(orderName, userId)
	if err == nil {
		runOrderInfo := runPkg.RunOrder{
			OrderName: orderName,
			UserNum:   len(runNums),
			UserId:    userId,
		}
		runOrderService.UpdateRunOrderUsers(runOrderInfo)
	}
}
