package stock

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
	stockRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockBasicApi struct{}

func (e *StockBasicApi) GetStockBasicList(c *gin.Context) {

	var pageInfo stockReq.SearchStoBasic

	err := c.ShouldBindJSON(&pageInfo)
	fmt.Println(pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	stockBasicList, total, err := stockBasicApi.GetStockBasicList(utils.GetUserAuthorityId(c), pageInfo.PageInfo, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     stockBasicList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetIndustry(c *gin.Context) {
	stockBasicList, err := stockBasicApi.GetStockIndustry()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: stockBasicList,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetMaxDate(c *gin.Context) {
	var date stockReq.MaxDay
	fmt.Println("1212212")
	err := c.ShouldBindQuery(&date)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	stockDate, err := stockDateServer.GetStockMaxDate(date.Inter)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(stock.StockDate{
		stockDate,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetGainianStock(c *gin.Context) {
	//var date stockReq.MaxDay
	//err := c.ShouldBindJSON(&date)
	//if err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	list, err := stockIndustryServer.GetGainianStock()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetMingci(c *gin.Context) {
	list, err := stockBasicApi.GetMingci()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetZhiShuCode(c *gin.Context) {
	var pageOrder stockReq.PageOrder
	c.ShouldBindQuery(&pageOrder)
	fmt.Println(pageOrder)
	list, err := stockBasicApi.GetZhiShuCode(pageOrder)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetLatest(c *gin.Context) {
	var basic stock.Basic
	c.ShouldBindQuery(&basic)
	list, err := stockBasicApi.GetLatest(basic)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (e *StockBasicApi) GetDailyLimit(c *gin.Context) {
	var basic stockReq.StockDaily
	c.ShouldBindQuery(&basic)
	list, err := stockDailyServer.GetDailyLimit(basic)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)

}

func (e *StockBasicApi) GetDailyPctChg(c *gin.Context) {
	var basic stockRes.StockMonthPct
	c.ShouldBindQuery(&basic)
	list, err := stockDailyServer.GetDailyPctChg(basic)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)

}
