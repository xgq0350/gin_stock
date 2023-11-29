package stock

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type StockBasicRouter struct{}

func (e *StockBasicRouter) InitStockBasicRouter(Router *gin.RouterGroup) {
	stockBasicRouter := Router.Group("stock")
	stockBasicApi := v1.ApiGroupApp.StockBasicGroup.StockBasicApi
	stockRankApi := v1.ApiGroupApp.StockBasicGroup.StockRankApi
	stockMessageApi := v1.ApiGroupApp.StockBasicGroup.StockMessageApi
	stockBuyApi := v1.ApiGroupApp.StockBasicGroup.StockBuyApi
	{
		stockBasicRouter.POST("basicList", stockBasicApi.GetStockBasicList)
		stockBasicRouter.GET("industry", stockBasicApi.GetIndustry)
		stockBasicRouter.GET("trade_date", stockBasicApi.GetMaxDate) //获取最近的日期
		stockBasicRouter.POST("gainian_stock", stockBasicApi.GetGainianStock)
		stockBasicRouter.GET("mingci", stockBasicApi.GetMingci)     //获取概念名词列表
		stockBasicRouter.GET("zhishu", stockBasicApi.GetZhiShuCode) //获取指数关联股票
		stockBasicRouter.GET("basic_info", stockBasicApi.GetLatest) //获取股票明細
	}

	{
		stockBasicRouter.GET("industry_rank", stockRankApi.GetIndustryRank)
		stockBasicRouter.GET("industry_latest", stockRankApi.GetIndustryLatest)
		stockBasicRouter.GET("industry_rank_detail", stockRankApi.GetIndustryDetail)
		stockBasicRouter.GET("zl_low_total", stockRankApi.GetZlTotal)
		stockBasicRouter.GET("zl_week_avg", stockRankApi.GetWeekAvgPro)
		stockBasicRouter.GET("industry_top", stockRankApi.GetIndustryTop)
		stockBasicRouter.GET("daily_limit", stockBasicApi.GetDailyLimit)
		stockBasicRouter.GET("daily_pct_chg", stockBasicApi.GetDailyPctChg)

	}

	{
		stockBasicRouter.GET("holders", stockMessageApi.GetHolderByIndustry)
	}

	{
		stockBasicRouter.GET("buy_summary", stockBuyApi.GetStockBuySummary)
		stockBasicRouter.GET("buy_list", stockBuyApi.GetStockBuyList)
		stockBasicRouter.GET("buy_detail", stockBuyApi.GetStockBuyDetail)
		stockBasicRouter.GET("buy_intention", stockBuyApi.GetIntention)
	}

	{
		stockBasicRouter.GET("rank_month", stockRankApi.GetStockRankMonth)
		stockBasicRouter.POST("stock_rank", stockRankApi.GetStockRank)
		stockBasicRouter.GET("year_close", stockRankApi.GetYearClose)
		stockBasicRouter.GET("pct_sum", stockRankApi.GetPctSum)
		stockBasicRouter.GET("section", stockRankApi.GetSection)
		stockBasicRouter.GET("talib_bands_summary", stockRankApi.BbAnds)
		stockBasicRouter.GET("talib_bands_list", stockRankApi.GetBandsList)
		stockBasicRouter.GET("talib_bands_detail", stockRankApi.BbAndDetails)

	}
}
