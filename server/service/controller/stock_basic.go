package controller

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
)

type StockBasic struct {
}

func (exa *StockBasic) GetStockBasicList(sysUserAuthorityID uint, info request.PageInfo, stockBasic stockReq.SearchStoBasic) (list interface{}, total int64, err error) {
	return list, total, nil
}
