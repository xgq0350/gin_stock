package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/stock"

type StoIntentionRes struct {
	stock.BuyIntention
	stock.StockCloseMax
	stock.StockClose
}
