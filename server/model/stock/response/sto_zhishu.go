package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/stock"

type StoZhishu struct {
	TsCode     string `json:"ts_code"`
	CodeName   string `json:"code_name"`
	Industry   string `json:"industry"`
	ZsName     string `json:"zs_name"`
	Caiwufenxi string `json:"caiwufenxi"`
	Name       string `json:"name"`
}

type ZhishuMaxClose struct {
	StoZhishu
	stock.StockCloseMax
	stock.StockClose
}
