package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type StockBasic struct {
	global.GVA_MO
	Basic
	StockHuaBasic
	ZlList
}

type Basic struct {
	Name     string `json:"name" form:"name" gorm:"comment:股票名称"`         // 客户手机号
	Industry string `json:"industry" form:"industry" gorm:"comment:股票行业"` // 股票行业
	Fullname string `json:"fullname" form:"fullname" gorm:"comment:股票全名"` // 股票全名
	Symbol   string `json:"symbol" form:"symbol" gorm:"comment:股票标识"`
	TsCode   string `json:"ts_code" form:"ts_code" gorm:"comment:股票标识" db:"code"`
	Close    string `json:"close" form:"close" gorm:"comment:股票标识"`
	Area     string `json:"area" form:"area" gorm:"comment:股票地方"`
}

type BasicInfo struct {
	global.GVA_MO
	Basic
	StockHuaBasic
	StockCloseMax
	YearCloseBasic
}

func (StockBasic) TableName() string {
	return "stock_basic"
}
