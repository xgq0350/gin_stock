package stock

type StockRankMonth struct {
	Industry string `json:"industry" form:"industry"`
	YearC    int    `json:"year_c" form:"year_c"`
	MonthC   int    `json:"month_c"`
	TsCode   string `json:"ts_code"`
	Name     string `json:"name"`
	Rank1    string `json:"rank1"`
	RankMonth
}

type RankMonth struct {
	Month_1  string `json:"month1"`
	Month_2  string `json:"month2"`
	Month_3  string `json:"month3" `
	Month_4  string `json:"month4" `
	Month_5  string `json:"month5" `
	Month_6  string `json:"month6" `
	Month_7  string `json:"month7" `
	Month_8  string `json:"month8" `
	Month_9  string `json:"month9" `
	Month_10 string `json:"month10" `
	Month_11 string `json:"month11" `
	Month_12 string `json:"month12" `
}

func (StockRankMonth) TableName() string {
	return "stock_rank_month"
}
