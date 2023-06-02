import service from '@/utils/request'

// @Tags RunNum
// @Summary 创建RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunNum true "创建RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runNum/createRunNum [post]
export const createRunNum = (data) => {
  // console.log('data --- ' + JSON.stringify(data))
  return service({
    url: '/runNum/createRunNum',
    method: 'post',
    data
  })
}

// @Tags RunNum
// @Summary 删除RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunNum true "删除RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runNum/deleteRunNum [delete]
export const deleteRunNum = (data) => {
  return service({
    url: '/runNum/deleteRunNum',
    method: 'delete',
    data
  })
}

// @Tags RunNum
// @Summary 删除RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runNum/deleteRunNum [delete]
export const deleteRunNumByIds = (data) => {
  return service({
    url: '/runNum/deleteRunNumByIds',
    method: 'delete',
    data
  })
}

// @Tags RunNum
// @Summary 更新RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunNum true "更新RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runNum/updateRunNum [put]
export const updateRunNum = (data) => {
  return service({
    url: '/runNum/updateRunNum',
    method: 'put',
    data
  })
}

// @Tags RunNum
// @Summary 用id查询RunNum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RunNum true "用id查询RunNum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runNum/findRunNum [get]
export const findRunNum = (params) => {
  return service({
    url: '/runNum/findRunNum',
    method: 'get',
    params
  })
}

// @Tags RunNum
// @Summary 分页获取RunNum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RunNum列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runNum/getRunNumList [get]
export const getRunNumList = (params) => {
  console.log('params -- ' + JSON.stringify(params))
  return service({
    url: '/runNum/getRunNumList',
    method: 'get',
    params
  })
}

export const getRunOrders = (params) => {
  return service({
    url: '/runNum/getRunOrders',
    method: 'get',
    params
  })
}


export const updateRunNumByIds = (data) => {
  return service({
    url: '/runNum/updateRunNumByIds',
    method: 'put',
    data
  })
}

