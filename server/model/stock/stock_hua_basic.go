package stock

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type StockHuaBasic struct {
	global.GVA_MO
	Liangdian   string `json:"liangdian" form:"liangdian" gorm:"comment:亮点"`
	Shenwan     string `json:"shenwan" form:"shenwan" gorm:"comment:申万"`
	Caiwufenxi  string `json:"caiwufenxi" form:"caiwufenxi" gorm:"comment:财务分类"`
	Gainian     string `json:"gainian" form:"gainian" gorm:"comment:概念"`
	Fenlei      string `json:"fenlei" form:"fenlei" gorm:"comment:分类"`
	Zhuying     string `json:"zhuying" form:"zhuying" gorm:"comment:主营"`
	Shiyingdong string `json:"shiyingdong"`
	Shiyingjing string `json:"shiyingjing"`
}

func (StockHuaBasic) TableName() string {
	return "stock_hua_basic"
}
