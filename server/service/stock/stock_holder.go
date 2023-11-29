package stock

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stoRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
)

type StockHolderServer struct{}

func (e *StockHolderServer) GetHolderByIndustry(basic stock.Basic) (list interface{}, err error) {
	var result []stoRes.StoHolder
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}

	var db = global.MustGetGlobalDBByDBName("stock")
	sql := "select ts_code,name, industry,max(if( year_code = '2021',holders,0)) as holder_2021,\n\t\tmax(if( year_code = '2022',holders,0)) as holder_2022,\n\t\tmax(if( year_code = '2023',holders,0)) as holder_2023\n\t\tfrom (select ts_code,year(trade_date) as year_code,GROUP_CONCAT(holders ORDER BY trade_date asc) as holders from hua_holders GROUP BY ts_code,year(trade_date)) a "
	sql = sql + "left join  stock_basic using(ts_code) "

	if basic.Industry != "" {
		sql = sql + " where industry = '" + basic.Industry + "'"
	}
	sql = sql + " GROUP BY ts_code,industry"
	fmt.Println(sql)

	err = db.Raw(sql).Find(&result).Error
	if err != nil {

	}
	redisSetList(getMethodName()+key, result)
	return result, nil
}
