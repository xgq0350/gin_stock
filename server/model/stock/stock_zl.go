package stock

type StockZlList struct {
	TsCode string `json:"ts_code" form:"ts_code"`
	ZlList
	StockDate
}

type ZlList struct {
	ProfitPercent float64 `json:"profit_percent" form:"profit_percent"`
	AvgPrice      float64 `json:"avg_price" form:"avg_price"`
	LatestPrice   float64 `json:"latest_price" form:"latest_price"`
	LowerestPrice float64 `json:"lowerest_price"`
	LatLow        float64 `json:"lat_low"`
}

type StockZlTotal struct {
	Basic
	Total
	StockCloseMax
	StockClose
}

type Total struct {
	Total string `json:"total"`
}
