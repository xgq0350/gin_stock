package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"

type StockDetail struct {
	res   response.PageResult
	Money string `json:"money"`
}

type StockMonthPct struct {
	Year        int    `json:"year" form:"year"`
	Month       int    `json:"month"`
	TsCode      string `json:"ts_code" form:"ts_code"`
	PctChg      string `json:"pct_chg"`
	PctPopular4 int    `json:"pct_popular4"`
	PctPos4     int    `json:"pct_pos4"`
	PctPopular2 int    `json:"pct_popular2"`
	PctPos2     int    `json:"pct_pos2"`
}
