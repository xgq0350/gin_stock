package response

type StoHolder struct {
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Holder_2021 string `json:"holder_2021"`
	Holder_2022 string `json:"holder_2022"`
	Holder_2023 string `json:"holder_2023"`
	TsCode      string `json:"ts_code"`
}

func (StoHolder) TableName() string {
	return "stock_holder"
}
