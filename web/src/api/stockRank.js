import service from '@/utils/request'
// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]

export const getZlDaily = (params) => {
  return service({
    url: '/stock/zl_week_avg',
    method: 'get',
    params
  })
}

export const getZlLow = (params) => {
  return service({
    url: '/stock/zl_low_total',
    method: 'get',
    params
  })
}

export const getIndustryTop = (params) => {
  return service({
    url: '/stock/industry_top',
    method: 'get',
    params
  })
}

export const getStockRank = (data) => {
  return service({
    url: '/stock/stock_rank',
    method: 'post',
    data
  })
}

export const getStockRankMonth = (params) => {
  return service({
    url: '/stock/rank_month',
    method: 'get',
    params
  })
}
export const getStockYearClose = (params) => {
  return service({
    url: '/stock/year_close',
    method: 'get',
    params
  })
}

export const getPctSum = (params) => {
  return service({
    url: '/stock/pct_sum',
    method: 'get',
    params
  })
}

export const getSection = (params) => {
  return service({
    url: '/stock/section',
    method: 'get',
    params
  })
}

