package response

type BuySaleSummaryRes struct {
	Name        string `json:"name" form:"name" gorm:"comment:股票名称"`         // 客户手机号
	Industry    string `json:"industry" form:"industry" gorm:"comment:股票行业"` // 股票行业
	TsCode      string `json:"ts_code" form:"ts_code" gorm:"comment:股票标识"`
	Symbol      string `json:"symbol" form:"symbol" gorm:"comment:股票标识"`
	IsHolder    bool   `json:"is_holder" form:"is_holder" gorm:"comment:是否持有"`
	LastSale    string `json:"last_sale" form:"last_sale" gorm:"comment:最近卖出时间"`
	BuyNum      int    `json:"buy_num" `
	BuyHisRate  string `json:"buy_his_rate" `
	SaleHisRate string `json:"sale_his_rate" `
	BuyDate     string `json:"buy_date" form:"buy_date"`
	SaleDate    string `json:"sale_date" form:"sale_date"`
	Open        string `json:"open" gorm:"comment:最后持有年份"`
	High        string `json:"high" gorm:"comment:最后持有年份"`
	Low         string `json:"low" gorm:"comment:最后持有年份"`
	Close       string `json:"close" gorm:"comment:最后持有年份"`
	PctChg      string `json:"pct_chg" gorm:"comment:涨幅"`
}

func (BuySaleSummaryRes) TableName() string {
	return "buy_sale_summary"
}
