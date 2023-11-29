package stock

type StockDingyi struct {
	Mingci   string `json:"mingci"`
	Mcdingyi string `json:"mcdingyi"`
}

func (StockDingyi) TableName() string {
	return "dingyi"
}
