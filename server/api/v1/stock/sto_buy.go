package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stoRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockBuyApi struct{}

/*
*
获取买卖汇总
*/
func (e *StockBuyApi) GetStockBuySummary(c *gin.Context) {
	list, err := stockBuyServer.GetBuySummary()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

/*
*
获取买卖明细
*/
func (e *StockBuyApi) GetStockBuyList(c *gin.Context) {
	var basic stock.StockBasic
	err := c.ShouldBindQuery(&basic)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("参数获取失败"+err.Error(), c)
		return
	}
	list, err := stockBuyServer.GetBuyList(basic.Symbol)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

/*
*
获取买卖明细
*/
func (e *StockBuyApi) GetStockBuyDetail(c *gin.Context) {
	var summary stoRes.BuySaleSummaryRes
	err := c.ShouldBindQuery(&summary)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("参数获取失败"+err.Error(), c)
		return
	}
	list, err := stockController.GetBuyDetail(summary)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

/*
*
获取买卖明细
*/
func (e *StockBuyApi) GetIntention(c *gin.Context) {
	//var summary stoRes.BuySaleSummaryRes
	//err := c.ShouldBindQuery(&summary)
	//if err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("参数获取失败"+err.Error(), c)
	//	return
	//}
	list, err := stockBuyServer.GetIntention()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}
