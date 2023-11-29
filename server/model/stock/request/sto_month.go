package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
)

type StockMonthRes struct {
	YearCs []int `json:"year_c_s"`
	stock.StockRankMonth
	request.PageInfo
}
