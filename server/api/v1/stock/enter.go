package stock

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	StockBasicApi
	StockRankApi
	StockMessageApi
	StockBuyApi
}

var (
	stockBasicApi       = service.ServiceGroupApp.StockServiceGroup.StockBasicServer
	stockDateServer     = service.ServiceGroupApp.StockServiceGroup.StockDateServer
	stockIndustryServer = service.ServiceGroupApp.StockServiceGroup.StockIndustryServer
	stockZlServer       = service.ServiceGroupApp.StockServiceGroup.StockZlServer
	stockHolderServer   = service.ServiceGroupApp.StockServiceGroup.StockHolderServer
	stockBuyServer      = service.ServiceGroupApp.StockServiceGroup.StockBuyServer
	stockController     = service.ServiceGroupApp.StockServiceController.StockController
	stockDailyServer    = service.ServiceGroupApp.StockServiceGroup.StockDailyServer
)
