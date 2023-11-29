package stock

type Section struct {
	Tscode  string  `json:"ts_code" form:"ts_code"`
	Money   float64 `json:"money"`
	Type    string  `json:"type" form:"type"`
	Num     int64   `json:"num"`
	OrderBy int64   `json:"order_by"`
}

type Bands struct {
	Name      string  `json:"name"`
	Industry  string  `json:"industry"`
	TsCode    string  `json:"ts_code" form:"ts_code"`
	Close     string  `json:"close" `
	Lower     string  `json:"lower" `
	TradeDate string  `json:"trade_date" `
	PctChg    string  `json:"pct_chg" `
	SzzsC     string  `json:"szzs_c" `
	Span      string  `json:"span" form:"span"`
	Scope     float64 `json:"scope" form:"scope"`
	Type      float64 `json:"type" form:"type"`
	Num       int64   `json:"num"`
	Value     float64 `json:"value"`
}
