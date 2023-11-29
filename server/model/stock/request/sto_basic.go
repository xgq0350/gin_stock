package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
)

type SearchStoBasic struct {
	stock.StockBasic
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}

type MaxDay struct {
	Inter int `json:"inter"  form:"inter" ` //间隔天数
}

type PageOrder struct {
	request.PageInfo
	Order
}

type Order struct {
	OrderKey string `json:"orderKey" form:"orderKey"` // 排序
	Desc     bool   `json:"desc" form:"desc"`         // 排序方式:升序false(默认)|降序true
}
