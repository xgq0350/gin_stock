package stock

type StockCloseMax struct {
	TsCode        string `json:"ts_code" gorm:"comment:股票名称"` // 客户手机号
	CloseMax_2023 string `json:"close_max_2023" `
	CloseMax_2022 string `json:"close_max_2022" `
	CloseMax_2021 string `json:"close_max_2021" `
	CloseMin_2023 string `json:"close_min_2023" `
	CloseMin_2022 string `json:"close_min_2022" `
	CloseMin_2021 string `json:"close_min_2021" `
}

type StockClose struct {
	Close    string `json:"close" gorm:"comment:今日收盘"`
	PctChg   string `json:"pct_chg" gorm:"comment:涨跌"`
	PreClose string `json:"pre_close" gorm:"comment:昨日收盘"`
}

func (StockCloseMax) TableName() string {
	return "close_max_min"
}
