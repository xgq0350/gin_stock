package stock

type StockDate struct {
	TradeDate string `json:"trade_date" form:"trade_date" gorm:"comment:股票日期"` //日期
}

func (StockDate) TableName() string {
	return "stock_date"
}
