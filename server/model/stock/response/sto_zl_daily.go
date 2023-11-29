package response

type StoZlDaily struct {
	Year          string `json:"year"`
	Week          string `json:"week"`
	ProfitPercent string `json:"profit_percent"`
	AvgClose      string `json:"avg_close"`
	Buy1          string `json:"buy1"`
	Buy2          string `json:"buy2"`
	Sale1         string `json:"sale1"`
	Sale2         string `json:"sale2"`
	ConcatClose   string `json:"concat_close"`
}
