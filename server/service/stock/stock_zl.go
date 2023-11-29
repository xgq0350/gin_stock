package stock

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	strRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
	"strconv"
)

type StockZlServer struct{}

func (e *StockZlServer) GetZlLowByDate(profit stock.StockZlList, date stock.StockDate) (list interface{}, total int64, err error) {
	var result []stock.StockZlTotal
	var key = date.TradeDate
	var keyC = key + "count"

	errKR := redisGetList(getMethodName()+key, &result)
	totalStr, errCR := redisGet(getMethodName() + keyC)
	if errKR == nil && errCR == nil {
		fmt.Println("use redis")
		totals, _ := strconv.Atoi(totalStr)
		total = int64(totals)
		return result, total, nil
	}

	var db = global.MustGetGlobalDBByDBName("stock")

	selectSql := "select * from (select zlshenqi.ts_code as symbol,count(1) as total,name,fullname, industry "
	countSql := "select * from (select ts_code as symbol, count(1) as total "
	sql := " from zlshenqi left join stock_basic using(ts_code)"

	if date.TradeDate == "" {
		date.TradeDate = "20230801"
	}
	sql = sql + "\nwhere trade_date >'" + date.TradeDate + "'"

	pro := "2.0"
	if profit.ProfitPercent != 0.0 {
		pro = fmt.Sprintf("%f", profit.ProfitPercent)
	}
	sql = sql + " and profit_percent < " + pro

	sql = sql + " GROUP BY ts_code ORDER BY total desc) a"
	sql = sql + " left join (select ts_code,close_max_2023,close_max_2022,close_max_2021," +
		"close_min_2023,close_min_2022,close_min_2021 from close_max_min) b on a.symbol = b.ts_code "
	sql = sql + " left join (select close,ts_code,pre_close,pct_chg from " +
		"(select max(trade_date) as trade_date from stock_date) c_t " +
		"left join stock_daily USING(trade_date)) c on a.symbol = c.ts_code"
	countSql = countSql + sql
	selectSql = selectSql + sql

	db.Raw(selectSql).Scan(&result)
	db.Raw(countSql).Count(&total)
	redisSetList(getMethodName()+key, result)
	redisSetTotal(getMethodName()+keyC, int(total))

	return result, total, nil
}

func (e *StockZlServer) GetWeekAvgPro(zlList stock.StockZlList) (list interface{}, err error) {
	var db = global.MustGetGlobalDBByDBName("stock")
	var result []strRes.StoZlDaily
	val, _ := json.Marshal(zlList)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}

	sql := "select year(zlshenqi.trade_date) as year,week(zlshenqi.trade_date) as week,AVG(profit_percent) as profit_percent, round(avg(close),2) as avg_close,"
	sql = sql + "round(round(avg(close),4) * 0.99,2) as buy1,round(round(avg(close),4) * 1.01,2) as sale1,"
	sql = sql + "round(round(avg(close),4) * 0.98,2) as buy2,round(round(avg(close),4) * 1.02,2) as sale2,"
	sql = sql + "GROUP_CONCAT(close ORDER BY zlshenqi.trade_date asc) as concat_close "
	sql = sql + " from zlshenqi  left join stock_daily on stock_daily.ts_code = zlshenqi.ts_code and stock_daily.trade_date = zlshenqi.trade_date"

	var tscode = zlList.TsCode
	var tradeDate = zlList.TradeDate
	if zlList.TsCode == "" {
		tscode = "000001.SZ"
	}
	sql = sql + " where zlshenqi.ts_code = \"" + tscode + "\" and "
	if zlList.TradeDate == "" {
		tradeDate = "20210101"
	}
	sql = sql + "zlshenqi.trade_date > '" + tradeDate + "' GROUP BY WEEK(trade_date),year(zlshenqi.trade_date) order by year(zlshenqi.trade_date) desc, week(trade_date) desc\n "
	fmt.Println(sql)
	db.Raw(sql).Scan(&result)
	redisSetList(key, result)

	return result, nil
}
