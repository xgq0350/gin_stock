package stock

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	"github.com/gin-gonic/gin"
)

type StockMessageApi struct{}

func (e *StockMessageApi) GetHolderByIndustry(c *gin.Context) {
	var industry stock.Basic
	err := c.ShouldBindQuery(&industry)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, err := stockHolderServer.GetHolderByIndustry(industry)

	response.OkWithData(response.PageResult{
		List: list,
	}, c)
}
