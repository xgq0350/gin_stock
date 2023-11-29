package stock

type BuySaleDate struct {
	TsCode    string  `json:"ts_code" gorm:"comment:股票名称"`     // 客户手机号
	Symbol    string  `json:"symbol" gorm:"comment:股票名称"`      // 客户手机号
	Name      string  `json:"name"  gorm:"comment:股票行业"`       // 股票
	Money     float64 `json:"money"  gorm:"comment:股票行业"`      // 股票
	Num       string  `json:"num"  gorm:"comment:股票行业"`        // 股票
	TradeDate string  `json:"trade_date"  gorm:"comment:股票行业"` // 股票
	Type      string  `json:"type"  gorm:"comment:股票行业"`       // 股票
	High      string  `json:"high"  gorm:"comment:最高价"`        // 股票
	Low       string  `json:"low"  gorm:"comment:最低价"`         // 股票
	Zhang     string  `json:"zhang" form:"comment:涨跌"`
	Vol       float64 `json:"vol" form:"comment:发生"`
	Leiji     string  `json:"leiji" form:"comment:累积金额"`
	Pingjun   string  `json:"pingjun" form:"comment:"`
}

type BuySaleSummary struct {
	Name        string `json:"name" form:"name" gorm:"comment:股票名称"`         // 客户手机号
	Industry    string `json:"industry" form:"industry" gorm:"comment:股票行业"` // 股票行业
	TsCode      string `json:"ts_code" form:"ts_code" gorm:"comment:股票标识"`
	IsHolder    bool   `json:"is_holder" form:"is_holder" gorm:"comment:是否持有"`
	BuyNum      int    `json:"buy_num" `
	BuyHisRate  string `json:"buy_his_rate" `
	SaleHisRate string `json:"sale_his_rate" `
	Open        string `json:"open" form:"open" gorm:"comment:最后持有年份"`
	High        string `json:"high" form:"high" gorm:"comment:最后持有年份"`
	Low         string `json:"low" form:"low" gorm:"comment:最后持有年份"`
	Close       string `json:"close" form:"close" gorm:"comment:最后持有年份"`
	PctChg      string `json:"pcg_chg" form:"pct_chg" gorm:"comment:涨幅"`
}

type BuySale struct {
	Name     string `json:"name" form:"name" gorm:"comment:股票名称"` // 客户手机号
	TsCode   string `json:"ts_code" form:"ts_code" gorm:"comment:股票标识"`
	IsHolder bool   `json:"is_holder" form:"is_holder" gorm:"comment:是否持有"`
	Shouyi   string `json:"shouyi" `
	BuyDate  string `json:"buy_date" form:"buy_date" gorm:"comment:买入日期"`
	SaleDate string `json:"sale_date" form:"sale_date" gorm:"comment:卖出日期"`
	BuyRate  string `json:"buy_rate" form:"buy_rate" gorm:"comment:买入价格"`
	SaleRate string `json:"sale_rate" form:"sale_rate" gorm:"comment:卖出价格"`
}

type StatisticMonth struct {
	LatestClose float64 `json:"latest_close"  gorm:"comment:最新价"`
	MonthFudu   string  `json:"month_fudu" gorm:"comment:近一月涨幅"`
	YearFudu    string  `json:"year_fudu" gorm:"comment:一年涨幅"`
}

type BuyIntention struct {
	Name      string `json:"name" form:"name" gorm:"comment:股票名称"`         // 客户手机号
	Industry  string `json:"industry" form:"industry" gorm:"comment:股票行业"` // 股票行业
	TsCode    string `json:"ts_code" form:"ts_code" gorm:"comment:股票标识"`
	TradeDate string `json:"trade_date" form:"trade_date" gorm:"comment:是否持有"`
	BuyRate   string `json:"buy_rate" form:"buy_rate" gorm:"comment:是否持有"`
	Record    bool   `json:"record" form:"record" gorm:"comment:是否持有"`
	Symbol    string `json:"symbol" form:"symbol" gorm:"comment:是否持有"`
}

func (BuySaleDate) TableName() string {
	return "buy_sale_date"
}
func (BuySale) TableName() string {
	return "buy_sale_date"
}
