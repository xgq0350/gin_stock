package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
)

type StockDateServer struct{}

func (e *StockDateServer) GetStockMaxDate(day int) (date string, err error) {
	var stockDate stock.StockDate
	key := ""
	date, errKR := redisGet(getMethodName() + key)
	if errKR == nil {
		return date, nil
	}
	db := global.MustGetGlobalDBByDBName("stock").Model(&stock.StockDate{})

	db.Select("trade_date").Order("trade_date desc").Offset(day).Limit(1).Find(&stockDate)

	redisSetStr(getMethodName()+key, stockDate.TradeDate)
	return stockDate.TradeDate, nil
}
