package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type YearClose struct {
	global.GVA_MO
	Basic
	YearCloseBasic
}

type PctSum struct {
	global.GVA_MO
	Basic
	PctClose
}

type PctClose struct {
	Chg_60 string `json:"chg_60"`
	Chg_30 string `json:"chg_30"`
}

type YearCloseBasic struct {
	Max_180d   string `json:"max_180d" form:"name" gorm:"comment:股票名称"`     // 客户手机号
	Min_180d   string `json:"min_180d" form:"industry" gorm:"comment:股票行业"` // 股票行业
	Max_90d    string `json:"max_90d" form:"fullname" gorm:"comment:股票全名"`  // 股票全名
	Min_90d    string `json:"min_90d" form:"symbol" gorm:"comment:股票标识"`
	Max_60d    string `json:"max_60d" form:"ts_code" gorm:"comment:股票标识"`
	Min_60d    string `json:"min_60d" form:"ts_code" gorm:"comment:股票标识"`
	YearClose  string `json:"year_close" form:"ts_code" gorm:"comment:股票标识"`
	MonthClose string `json:"month_close" form:"ts_code" gorm:"comment:股票标识"`
	YearChg    string `json:"year_chg" form:"ts_code" gorm:"comment:股票标识"`
	MonthChg   string `json:"month_chg" form:"ts_code" gorm:"comment:股票标识"`
}
