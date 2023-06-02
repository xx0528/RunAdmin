import service from '@/utils/request'

// @Tags RunPage
// @Summary 创建RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunPage true "创建RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runPage/createRunPage [post]
export const createRunPage = (data) => {
  return service({
    url: '/runPage/createRunPage',
    method: 'post',
    data
  })
}

// @Tags RunPage
// @Summary 删除RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunPage true "删除RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runPage/deleteRunPage [delete]
export const deleteRunPage = (data) => {
  return service({
    url: '/runPage/deleteRunPage',
    method: 'delete',
    data
  })
}

// @Tags RunPage
// @Summary 删除RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runPage/deleteRunPage [delete]
export const deleteRunPageByIds = (data) => {
  return service({
    url: '/runPage/deleteRunPageByIds',
    method: 'delete',
    data
  })
}

// @Tags RunPage
// @Summary 更新RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunPage true "更新RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runPage/updateRunPage [put]
export const updateRunPage = (data) => {
  return service({
    url: '/runPage/updateRunPage',
    method: 'put',
    data
  })
}

// @Tags RunPage
// @Summary 用id查询RunPage
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RunPage true "用id查询RunPage"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runPage/findRunPage [get]
export const findRunPage = (params) => {
  return service({
    url: '/runPage/findRunPage',
    method: 'get',
    params
  })
}

// @Tags RunPage
// @Summary 分页获取RunPage列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RunPage列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runPage/getRunPageList [get]
export const getRunPageList = (params) => {
  console.log('params -- ' + JSON.stringify(params))
  return service({
    url: '/runPage/getRunPageList',
    method: 'get',
    params
  })
}
