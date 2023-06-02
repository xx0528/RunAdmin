/*
 * @Author: xx
 * @Date: 2023-05-09 10:23:48
 * @LastEditTime: 2023-05-11 19:00:25
 * @Description: 
 */
import service from '@/utils/request'

// @Tags RunOrder
// @Summary 创建RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunOrder true "创建RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runOrder/createRunOrder [post]
export const createRunOrder = (data) => {
  return service({
    url: '/runOrder/createRunOrder',
    method: 'post',
    data
  })
}

// @Tags RunOrder
// @Summary 删除RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunOrder true "删除RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runOrder/deleteRunOrder [delete]
export const deleteRunOrder = (data) => {
  return service({
    url: '/runOrder/deleteRunOrder',
    method: 'delete',
    data
  })
}

// @Tags RunOrder
// @Summary 删除RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /runOrder/deleteRunOrder [delete]
export const deleteRunOrderByIds = (data) => {
  return service({
    url: '/runOrder/deleteRunOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags RunOrder
// @Summary 更新RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RunOrder true "更新RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /runOrder/updateRunOrder [put]
export const updateRunOrder = (data) => {
  return service({
    url: '/runOrder/updateRunOrder',
    method: 'put',
    data
  })
}

// @Tags RunOrder
// @Summary 用id查询RunOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RunOrder true "用id查询RunOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /runOrder/findRunOrder [get]
export const findRunOrder = (params) => {
  return service({
    url: '/runOrder/findRunOrder',
    method: 'get',
    params
  })
}

// @Tags RunOrder
// @Summary 分页获取RunOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RunOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /runOrder/getRunOrderList [get]
export const getRunOrderList = (params) => {
  return service({
    url: '/runOrder/getRunOrderList',
    method: 'get',
    params
  })
}

export const getRunPages = (params) => {
  return service({
    url: '/runOrder/getRunPages',
    method: 'get',
    params
  })
}

export const getOrderNums = (params) => {
  return service({
    url: '/runOrder/getOrderNums',
    method: 'get',
    params
  })
}

export const createRunNums = (data) => {
  console.log('data --- ' + JSON.stringify(data))
  return service({
    url: '/runOrder/createRunNums',
    method: 'post',
    data
  })
}
