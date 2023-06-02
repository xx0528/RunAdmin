import service from '@/utils/request'

// @Tags RunTpl
// @Summary 创建RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunTpl true "创建RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runTpl/createRunTpl [post]
export const createRunTpl = (data) => {
  return service({
    url: '/runTpl/createRunTpl',
    method: 'post',
    data
  })
}

// @Tags RunTpl
// @Summary 删除RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunTpl true "删除RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runTpl/deleteRunTpl [delete]
export const deleteRunTpl = (data) => {
  return service({
    url: '/runTpl/deleteRunTpl',
    method: 'delete',
    data
  })
}

// @Tags RunTpl
// @Summary 删除RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runTpl/deleteRunTpl [delete]
export const deleteRunTplByIds = (data) => {
  return service({
    url: '/runTpl/deleteRunTplByIds',
    method: 'delete',
    data
  })
}

// @Tags RunTpl
// @Summary 更新RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunTpl true "更新RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runTpl/updateRunTpl [put]
export const updateRunTpl = (data) => {
  return service({
    url: '/runTpl/updateRunTpl',
    method: 'put',
    data
  })
}

// @Tags RunTpl
// @Summary 用id查询RunTpl
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RunTpl true "用id查询RunTpl"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runTpl/findRunTpl [get]
export const findRunTpl = (params) => {
  return service({
    url: '/runTpl/findRunTpl',
    method: 'get',
    params
  })
}

// @Tags RunTpl
// @Summary 分页获取RunTpl列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RunTpl列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runTpl/getRunTplList [get]
export const getRunTplList = (params) => {
  return service({
    url: '/runTpl/getRunTplList',
    method: 'get',
    params
  })
}

// 复制一行
export const copyRunTpl = (data) => {
  return service({
    url: '/runTpl/copyRunTpl',
    method: 'post',
    data
  })
}
