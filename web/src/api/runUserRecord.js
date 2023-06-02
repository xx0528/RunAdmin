import service from '@/utils/request'

// @Tags RunUserRecord
// @Summary 创建RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunUserRecord true "创建RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runUserRecord/createRunUserRecord [post]
export const createRunUserRecord = (data) => {
  return service({
    url: '/runUserRecord/createRunUserRecord',
    method: 'post',
    data
  })
}

// @Tags RunUserRecord
// @Summary 删除RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunUserRecord true "删除RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runUserRecord/deleteRunUserRecord [delete]
export const deleteRunUserRecord = (data) => {
  return service({
    url: '/runUserRecord/deleteRunUserRecord',
    method: 'delete',
    data
  })
}

// @Tags RunUserRecord
// @Summary 删除RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runUserRecord/deleteRunUserRecord [delete]
export const deleteRunUserRecordByIds = (data) => {
  return service({
    url: '/runUserRecord/deleteRunUserRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags RunUserRecord
// @Summary 更新RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunUserRecord true "更新RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runUserRecord/updateRunUserRecord [put]
export const updateRunUserRecord = (data) => {
  return service({
    url: '/runUserRecord/updateRunUserRecord',
    method: 'put',
    data
  })
}

// @Tags RunUserRecord
// @Summary 用id查询RunUserRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RunUserRecord true "用id查询RunUserRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runUserRecord/findRunUserRecord [get]
export const findRunUserRecord = (params) => {
  return service({
    url: '/runUserRecord/findRunUserRecord',
    method: 'get',
    params
  })
}

// @Tags RunUserRecord
// @Summary 分页获取RunUserRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RunUserRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runUserRecord/getRunUserRecordList [get]
export const getRunUserRecordList = (params) => {
  return service({
    url: '/runUserRecord/getRunUserRecordList',
    method: 'get',
    params
  })
}
