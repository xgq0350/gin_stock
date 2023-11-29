package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
)

type StockDaily struct {
	stock.StockHuaBasic
	stock.StockCloseMax
	stock.StockClose
	stock.StockDate
	Name     string `json:"name"`
	Industry string `json:"industry" form:"industry" gorm:"comment:申万"`
	Order
}
