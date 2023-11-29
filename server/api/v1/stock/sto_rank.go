package stock

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StockRankApi struct{}

func (e *StockRankApi) GetStockRank(c *gin.Context) {
	var rank stock.StockRank
	err := c.ShouldBindJSON(&rank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := stockIndustryServer.GetStockRank(rank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

// 行业涨跌排名
func (e *StockRankApi) GetIndustryRank(c *gin.Context) {
	list, err := stockIndustryServer.GetIndustryRank()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

// 行业涨跌排名
func (e *StockRankApi) GetIndustryLatest(c *gin.Context) {
	list, err := stockIndustryServer.GetIndustryLatest()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

// 行业涨跌明细
func (e *StockRankApi) GetIndustryDetail(c *gin.Context) {
	var industry stock.IndustryRank
	err := c.ShouldBindQuery(&industry)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := stockIndustryServer.GetIndustryDetail(industry.Industry)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

// 行业涨跌明细
func (e *StockRankApi) GetZlTotal(c *gin.Context) {
	var ZlTotal stockReq.StoZlTotal
	err := c.ShouldBindQuery(&ZlTotal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := stockZlServer.GetZlLowByDate(ZlTotal.StockZlList, ZlTotal.StockDate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// 行业涨跌明细
func (e *StockRankApi) GetWeekAvgPro(c *gin.Context) {
	var ZlTotal stockReq.StoZlTotal
	err := c.ShouldBindQuery(&ZlTotal)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := stockZlServer.GetWeekAvgPro(ZlTotal.StockZlList)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

// 行业顶尖
func (e *StockRankApi) GetIndustryTop(c *gin.Context) {
	var page request.PageInfo
	c.ShouldBindQuery(&page)
	list, total, err := stockIndustryServer.GetIndustryTop(page)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List:  list,
		Total: total,
	}, c)
}

// 月度明细
func (e *StockRankApi) GetStockRankMonth(c *gin.Context) {
	var rank stockReq.StockMonthRes
	err := c.ShouldBindQuery(&rank)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(rank)
	list, err := stockIndustryServer.GetStockRankMonth(rank.Industry, rank.YearC, rank.PageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

func (e *StockRankApi) GetYearClose(c *gin.Context) {

	list, err := stockIndustryServer.GetYearClose()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

func (e *StockRankApi) GetPctSum(c *gin.Context) {

	list, err := stockIndustryServer.GetPctSum()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}

/*
*
获取分段数据
*/
func (e *StockRankApi) GetSection(c *gin.Context) {
	var section stock.Section
	err := c.ShouldBindQuery(&section)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("参数获取失败"+err.Error(), c)
		return
	}
	list, err := stockController.GetSection(section)
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
获取最近一个月可买布林带
*/
func (e *StockRankApi) BbAnds(c *gin.Context) {
	var bands stock.Bands
	err := c.ShouldBindQuery(&bands)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("参数获取失败"+err.Error(), c)
		return
	}
	list, err := stockController.GetBands(bands)
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
获取最近一个月可买布林带
*/
func (e *StockRankApi) GetBandsList(c *gin.Context) {
	var bands stock.Bands
	err := c.ShouldBindQuery(&bands)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("参数获取失败"+err.Error(), c)
		return
	}
	list, err := stockController.GetBandsList(bands)
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
获取最近一个月可买布林带
*/
func (e *StockRankApi) BbAndDetails(c *gin.Context) {
	var bands stock.Bands
	err := c.ShouldBindQuery(&bands)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("参数获取失败"+err.Error(), c)
		return
	}
	list, err := stockController.GetBabAndDetails(bands)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}
