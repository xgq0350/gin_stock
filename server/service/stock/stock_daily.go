package stock

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
	stockRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
)

type StockDailyServer struct {
}

func (e *StockDailyServer) GetDailyLimit(basic stockReq.StockDaily) (list interface{}, err error) {
	var result []stockReq.StockDaily
	var key = ""
	val, _ := json.Marshal(basic)
	key = getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}

	basicDate := ""
	if basic.TradeDate == "" {
		basicDate = " join (select max(trade_date) as max_date from stock_date) a on stock_daily.trade_date = a.max_date "
	} else {
		basicDate = " where trade_date = '" + basic.TradeDate + "'"
	}

	industry := ""
	if basic.Industry != "" {
		industry = " and industry = '" + basic.Industry + "' "
	}

	shenwan := ""
	if basic.Shenwan != "" {
		industry = " and shenwan = '" + basic.Shenwan + "' "
	}

	sql := "select stock_daily.ts_code,name,pct_chg,`close`,industry,liangdian,shenwan,caiwufenxi,gainian,fenlei,shiyingdong,\nclose_max_2023,close_max_2022,close_max_2021,\nclose_min_2023,close_min_2022,close_min_2021\n\n from stock_daily \n"
	sql = sql + " left join stock_hua_basic on stock_daily.ts_code = stock_hua_basic.ts_code_name\n"
	sql = sql + " left join close_max_min USING(ts_code)\n"
	sql = sql + basicDate + industry + shenwan
	db := global.MustGetGlobalDBByDBName("stock")

	order := ""
	if basic.OrderKey == "" {
		sql = sql + " order by pct_chg desc "
	}

	if basic.OrderKey != "" {
		order = basic.OrderKey
		if order == "pct_chg" {

		} else {
			if basic.Desc == false {
				order = "," + order + " desc "
			} else {
				order = "," + order + " asc "
			}
			sql = sql + order
		}
	}
	sql = sql + " limit 300;\n"
	db = db.Raw(sql)

	fmt.Println(sql)
	err = db.Find(&result).Error
	redisSetList(key, result)
	return result, err
}

func (e *StockDailyServer) GetDailyPctChg(basic stockRes.StockMonthPct) (list interface{}, err error) {
	var result []stockRes.StockMonthPct
	var key = ""
	val, _ := json.Marshal(basic)
	key = getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}

	if basic.TsCode == "" {
		basic.TsCode = "000001.SZ"
	}
	sql := "SELECT ts_code, year,month,GROUP_CONCAT(round(pct_chg,1)  ORDER BY trade_date desc  separator \" /\\\\ \") as pct_chg, \n"
	sql = sql + " sum(if(abs(pct_chg) > 4, 1, 0)) as pct_popular4,\nsum(if(pct_chg > 4, 1, 0)) as pct_pos4,\nsum(if(abs(pct_chg) > 2, 1, 0)) as pct_popular2,\nsum(if(pct_chg > 2, 1, 0)) as pct_pos2\n"
	sql = sql + "\nFROM `stock_daily` "
	sql = sql + "\nwhere ts_code ='" + basic.TsCode + "'"
	sql = sql + "\nGROUP BY ts_code, year,month ORDER BY `year` desc,`month` desc;"

	db := global.MustGetGlobalDBByDBName("stock")

	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	redisSetList(key, result)
	return result, err
}
