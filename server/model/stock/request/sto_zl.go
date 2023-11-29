package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/stock"

type StoZlTotal struct {
	stock.StockZlList
	stock.StockDate
}
