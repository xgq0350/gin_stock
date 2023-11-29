package stock

type IndustryDetail struct {
	Industry  string  `json:"industry"`
	Rise      float64 `json:"rise"`
	Num       float64 `json:"num"`
	Money     float64 `json:"money"`
	AddMoney  float64 `json:"add_money"`
	UpShop    int     `json:"up_shop"`
	downShop  int     `json:"down_shop"`
	TradeDate int     `json:"trade_date"`
	RankTop   string  `json:"rank_top"`
}

type StockRank struct {
	StockDate
	IndustryRank
	Name string `json:"name"`
}
type IndustryRank1 struct {
	Industry string `json:"industry"`
	Rank1    string `json:"rank1"`
}

type IndustryRank struct {
	Industry string `json:"industry"`
	Rank1    string `json:"rank1"`
}

type IndustryDaily struct {
	TradeDate string `json:"trade_date"`
	Top1      string `json:"top1"`
	Top2      string `json:"top2"`
	Top3      string `json:"top3"`
	Top4      string `json:"top4"`
	Top5      string `json:"top5"`
	Top6      string `json:"top6"`
	Top7      string `json:"top7"`
	Top8      string `json:"top8"`
	Top9      string `json:"top9"`
	Top10     string `json:"top10"`
	Top11     string `json:"top11"`
	Top12     string `json:"top12"`
	Top13     string `json:"top13"`
	Top14     string `json:"top14"`
	Top15     string `json:"top15"`
	Top16     string `json:"top16"`
	Top17     string `json:"top17"`
	Top18     string `json:"top18"`
	Top19     string `json:"top19"`
	Top20     string `json:"top20"`
	Top21     string `json:"top21"`
	Top22     string `json:"top22"`
	Top23     string `json:"top23"`
	Top24     string `json:"top24"`
	Top25     string `json:"top25"`
}

type IndustryLatest struct {
	Industry string `json:"industry"`
	Rank1    string `json:"rank1"`
}
