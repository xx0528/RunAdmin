package runPkg

import (
	"fmt"
	"net/http"

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

type RunTplApi struct {
}

var runTplService = service.ServiceGroupApp.RunPkgServiceGroup.RunTplService

// CreateRunTpl 创建RunTpl
// @Tags RunTpl
// @Summary 创建RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunTpl true "创建RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runTpl/createRunTpl [post]
func (runTplApi *RunTplApi) CreateRunTpl(c *gin.Context) {
	var runTpl runPkg.RunTpl
	err := c.ShouldBindJSON(&runTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runTpl.CreatedBy = utils.GetUserID(c)
	runTpl.UserId = runTpl.CreatedBy
	verify := utils.Rules{
		"TplName": {utils.NotEmpty()},
	}
	if err := utils.Verify(runTpl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	runTpl.TplId, runTpl.PageUrl = runTpl.GenerateUrl()

	if err := runTplService.CreateRunTpl(&runTpl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (runTplApi *RunTplApi) CopyRunTpl(c *gin.Context) {
	var runTpl runPkg.RunTpl
	err := c.ShouldBindJSON(&runTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rerunTpl, err := runTplService.GetRunTpl(runTpl.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	newRunTpl := runPkg.RunTpl{}
	newRunTpl.UserId = rerunTpl.UserId
	newRunTpl.TplName = rerunTpl.TplName + "[复制]"
	newRunTpl.TplType = rerunTpl.TplType
	newRunTpl.ClickUrl = rerunTpl.ClickUrl
	newRunTpl.ClickDesc = rerunTpl.ClickDesc
	newRunTpl.DialogDesc = rerunTpl.DialogDesc
	newRunTpl.Text1 = rerunTpl.Text1
	newRunTpl.Text2 = rerunTpl.Text2
	newRunTpl.Text3 = rerunTpl.Text3
	newRunTpl.Text4 = rerunTpl.Text4
	newRunTpl.Text5 = rerunTpl.Text5
	newRunTpl.Text6 = rerunTpl.Text6
	newRunTpl.Text7 = rerunTpl.Text7
	newRunTpl.Text8 = rerunTpl.Text8
	newRunTpl.Text9 = rerunTpl.Text9
	newRunTpl.Text10 = rerunTpl.Text10
	newRunTpl.Text11 = rerunTpl.Text11
	newRunTpl.Text12 = rerunTpl.Text12
	newRunTpl.Text13 = rerunTpl.Text13
	newRunTpl.Text14 = rerunTpl.Text14
	newRunTpl.Text15 = rerunTpl.Text15
	newRunTpl.Text16 = rerunTpl.Text16
	newRunTpl.Text17 = rerunTpl.Text17
	newRunTpl.Text18 = rerunTpl.Text18
	newRunTpl.Text19 = rerunTpl.Text19
	newRunTpl.Text20 = rerunTpl.Text20
	newRunTpl.PicName1 = rerunTpl.PicName1
	newRunTpl.PicName2 = rerunTpl.PicName2
	newRunTpl.PicName3 = rerunTpl.PicName3
	newRunTpl.PicName4 = rerunTpl.PicName4
	newRunTpl.PicName5 = rerunTpl.PicName5
	newRunTpl.PicName6 = rerunTpl.PicName6
	newRunTpl.PicName7 = rerunTpl.PicName7
	newRunTpl.PicName8 = rerunTpl.PicName8
	newRunTpl.PicName9 = rerunTpl.PicName9
	newRunTpl.PicName10 = rerunTpl.PicName10
	newRunTpl.CreatedBy = rerunTpl.CreatedBy
	newRunTpl.UpdatedBy = rerunTpl.UpdatedBy
	newRunTpl.DeletedBy = rerunTpl.DeletedBy

	newRunTpl.TplId, newRunTpl.PageUrl = newRunTpl.GenerateUrl()

	if err := runTplService.CreateRunTpl(&newRunTpl); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("复制成功", c)
	}
}

// DeleteRunTpl 删除RunTpl
// @Tags RunTpl
// @Summary 删除RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunTpl true "删除RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runTpl/deleteRunTpl [delete]
func (runTplApi *RunTplApi) DeleteRunTpl(c *gin.Context) {
	var runTpl runPkg.RunTpl
	err := c.ShouldBindJSON(&runTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runTpl.DeletedBy = utils.GetUserID(c)
	if err := runTplService.DeleteRunTpl(runTpl); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRunTplByIds 批量删除RunTpl
// @Tags RunTpl
// @Summary 批量删除RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /runTpl/deleteRunTplByIds [delete]
func (runTplApi *RunTplApi) DeleteRunTplByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := runTplService.DeleteRunTplByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRunTpl 更新RunTpl
// @Tags RunTpl
// @Summary 更新RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body runPkg.RunTpl true "更新RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runTpl/updateRunTpl [put]
func (runTplApi *RunTplApi) UpdateRunTpl(c *gin.Context) {
	var runTpl runPkg.RunTpl
	err := c.ShouldBindJSON(&runTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	runTpl.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"TplName": {utils.NotEmpty()},
	}
	if err := utils.Verify(runTpl, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := runTplService.UpdateRunTpl(runTpl); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRunTpl 用id查询RunTpl
// @Tags RunTpl
// @Summary 用id查询RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkg.RunTpl true "用id查询RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runTpl/findRunTpl [get]
func (runTplApi *RunTplApi) FindRunTpl(c *gin.Context) {
	var runTpl runPkg.RunTpl
	err := c.ShouldBindQuery(&runTpl)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerunTpl, err := runTplService.GetRunTpl(runTpl.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerunTpl": rerunTpl}, c)
	}
}

// GetRunTplList 分页获取RunTpl列表
// @Tags RunTpl
// @Summary 分页获取RunTpl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query runPkgReq.RunTplSearch true "分页获取RunTpl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runTpl/getRunTplList [get]
func (runTplApi *RunTplApi) GetRunTplList(c *gin.Context) {
	var pageInfo runPkgReq.RunTplSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pageInfo.UserId = utils.GetUserID(c)
	if list, total, err := runTplService.GetRunTplInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 获取到模板网页
func (runTplApi *RunTplApi) GetTplPage(c *gin.Context) {
	tplId := c.Query("tplId")
	rerunTpl, err := runTplService.GetRunTplByTplId(tplId)
	tplType := rerunTpl.TplType
	prefix := "/api/"
	if tplType == runPkg.TplType_Normal {
		c.HTML(http.StatusOK, "normal1", gin.H{
			"clickDesc":  rerunTpl.ClickDesc,
			"dialogDesc": rerunTpl.DialogDesc,
			"clickUrl":   rerunTpl.ClickUrl,
			"text1":      rerunTpl.Text1,
			"text2":      rerunTpl.Text2,
			"text3":      rerunTpl.Text3,
			"text4":      rerunTpl.Text4,
			"text5":      rerunTpl.Text5,
			"text6":      rerunTpl.Text6,
			"text7":      rerunTpl.Text7,
			"text8":      rerunTpl.Text8,
			"text9":      rerunTpl.Text9,
			"text10":     rerunTpl.Text10,
			"text11":     rerunTpl.Text11,
			"text12":     rerunTpl.Text12,
			"text13":     rerunTpl.Text13,
			"text14":     rerunTpl.Text14,
			"text15":     rerunTpl.Text15,
			"img1":       prefix + rerunTpl.PicName1,
			"img2":       prefix + rerunTpl.PicName2,
			"img3":       prefix + rerunTpl.PicName3,
			"img4":       prefix + rerunTpl.PicName4,
			"img5":       prefix + rerunTpl.PicName5,
			"img6":       prefix + rerunTpl.PicName6,
			"img7":       prefix + rerunTpl.PicName7,
		})
	} else if tplType == runPkg.TplType_Stock_Thai {
		c.HTML(http.StatusOK, "stockThai", gin.H{"clickUrl": rerunTpl.ClickUrl})
	} else if tplType == runPkg.TplType_Stock_Malaysia {
		c.HTML(http.StatusOK, "stockMalaysia", gin.H{"clickUrl": rerunTpl.ClickUrl})
	} else {
		global.GVA_LOG.Error(fmt.Sprintf("未找到模板类型 %s!", tplId), zap.Error(err))
		c.HTML(http.StatusOK, "error", gin.H{
			"title": "error ~",
			"code":  404,
			"msg":   "url not fund",
		})
		return
	}

}
