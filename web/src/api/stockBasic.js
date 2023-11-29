import service from '@/utils/request'
// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const getStockBasic = (data) => {
  return service({
    url: '/stock/basicList',
    method: 'post',
    data
  })
}

export const getIndustry = (params) => {
  return service({
    url: '/stock/industry',
    method: 'get',
    params
  })
}
export const getBasicInfo = (params) => {
  return service({
    url: '/stock/basic_info',
    method: 'get',
    params
  })
}

export const getBolders = (params) => {
  return service({
    url: '/stock/holders',
    method: 'get',
    params
  })
}

export const getMingci = (params) => {
  return service({
    url: '/stock/mingci',
    method: 'get',
    params
  })
}
export const getZhishu = (params) => {
  return service({
    url: '/stock/zhishu',
    method: 'get',
    params
  })
}


