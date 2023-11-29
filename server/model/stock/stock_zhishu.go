package stock

type StockZhishu struct {
	CodeName string `json:"code_name"`
	Type     string `json:"type"`
}

func (StockZhishu) TableName() string {
	return "zhishu"
}
