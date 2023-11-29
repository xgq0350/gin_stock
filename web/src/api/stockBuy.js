import service from '@/utils/request'
// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
export const getBuySummary = (params) => {
  return service({
    url: '/stock/buy_summary',
    method: 'get',
    params
  })
}

export const getBuyList = (params) => {
  return service({
    url: '/stock/buy_list',
    method: 'get',
    params
  })
}

export const getBuyDetail = (params) => {
  return service({
    url: '/stock/buy_detail',
    method: 'get',
    params
  })
}

export const getBuyIntention = (params) => {
  return service({
    url: '/stock/buy_intention',
    method: 'get',
    params
  })
}
export const getDailyLimit = (params) => {
  return service({
    url: '/stock/daily_limit',
    method: 'get',
    params
  })
}

export const getDetailPctChg = (params) => {
  return service({
    url: '/stock/daily_pct_chg',
    method: 'get',
    params
  })
}

export const getBandSummary = (params) => {
  return service({
    url: '/stock/talib_bands_summary',
    method: 'get',
    params
  })
}

export const getBandDetail = (params) => {
  return service({
    url: '/stock/talib_bands_detail',
    method: 'get',
    params
  })
}

export const getBandList = (params) => {
  return service({
    url: '/stock/talib_bands_list',
    method: 'get',
    params
  })
}
