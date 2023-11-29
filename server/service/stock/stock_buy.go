package stock

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stoRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
)

type StockBuyServer struct{}

func (e *StockBuyServer) GetBuySummary() (list interface{}, err error) {
	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select a.*,b.*,c.industry from (select sum(is_holder) as is_holder, ts_code as symbol,max(name) as name,\ncount(1) as buy_num," +
		" max(sale_date) as last_sale,\nGROUP_CONCAT(buy_rate ORDER BY buy_date asc) as buy_his_rate," +
		" GROUP_CONCAT(sale_rate ORDER BY sale_date asc) as sale_his_rate\n\n from buy_sale_summary " +
		" GROUP BY ts_code) a\n left join (\nselect ts_code,open,high,low,`close`,pct_chg from " +
		"(select max(trade_date) as trade_date from stock_date)c_t left join stock_daily USING(trade_date)) b " +
		"on a.symbol = left(b.ts_code,6) \nleft join stock_basic c on a.symbol = c.symbol ORDER BY is_holder desc\n;"
	var result []stoRes.BuySaleSummaryRes
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	err = db.Raw(sql).Find(&result).Error
	fmt.Println(sql)
	redisSetList(getMethodName(), result)
	return result, err
}

// 明细，几笔清仓
func (e *StockBuyServer) GetBuyList(code string) (list []stock.BuySale, err error) {
	var result []stock.BuySale
	val, _ := json.Marshal(code)
	key := fmt.Sprintf("%s", val)
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}

	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select * from `buy_sale_summary` "
	if code != "" {
		sql = sql + " where ts_code = \"" + code + "\""
	}
	sql = sql + " order by is_holder desc, sale_date desc"

	err = db.Raw(sql).Find(&result).Error
	redisSetList(getMethodName(), result)
	return result, err
}

func (e *StockBuyServer) GetBuyDetail(summmary stoRes.BuySaleSummaryRes) (list []stock.BuySaleDate, err error) {
	var result []stock.BuySaleDate

	db := global.MustGetGlobalDBByDBName("stock").Model(&stock.BuySaleDate{})
	//sql := "select * from `buy_sale_date` "
	if summmary.Symbol != "" {
		db.Where("symbol = ?", summmary.Symbol)
		//sql = sql + " where ts_code = \"" + summmary.Symbol + "\""
	}
	if summmary.BuyDate != "" {
		db.Where("trade_date >= ?", summmary.BuyDate)
		//sql = sql + " trade_date >= \"" + summmary.Symbol + "\""
	}
	if summmary.SaleDate != "" {
		db.Where("trade_date <= ?", summmary.SaleDate)
		//sql = sql + " trade_date <= \"" + summmary.SaleDate + "\""
	}
	db.Order(" trade_date asc")
	//sql = sql + " order by trade_date asc"

	err = db.Find(&result).Error

	return result, err
}

func (e *StockBuyServer) GetIntention() (list interface{}, err error) {
	var result []stoRes.StoIntentionRes
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}

	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select * from ((select ts_code as symbol, name,trade_date,buy_rate,0 as record from buy_intention ORDER BY trade_date desc )"
	sql = sql + "union all select ts_code as symbol, name,buy_date as trade_date, buy_rate,1 as record from buy_sale_summary where id in (SELECT max(id) FROM `buy_sale_summary` GROUP by ts_code ORDER BY buy_date asc) ) a "
	sql = sql + " left join (select ts_code, close_max_2023,close_max_2022,close_max_2021," +
		"close_min_2023,close_min_2022,close_min_2021 from close_max_min) b on a.symbol = left(b.ts_code,6) "
	sql = sql + " left join (select close,ts_code,pre_close,pct_chg from " +
		"(select max(trade_date) as trade_date from stock_date) c_t " +
		"left join stock_daily USING(trade_date)) c on a.symbol = left(c.ts_code, 6)"
	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	redisSetList(getMethodName(), result)
	return result, err
}

func (e *StockBuyServer) GetSection(code string, codeType string) (list []stock.Section, err error) {
	var result []stock.Section
	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select * from section where ts_code = " + "\"" + code + "\" "
	sql = sql + " and type = " + "\"" + codeType + "\""
	sql = sql + " order by order_by asc;"
	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	return result, err
}

func (e *StockBuyServer) GetBands(bands stock.Bands) (list []stock.Bands, err error) {
	var result []stock.Bands
	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select * from (select name, industry, talib_daily.ts_code,talib_daily.lower,latest.value,"
	sql = sql + " stock_daily.close,talib_daily.trade_date,pct_chg,szzs.`close` as szzs_c,"
	sql = sql + " ROW_NUMBER() over (PARTITION by ts_code ORDER BY trade_date desc) as num"
	sql = sql + " from talib_daily join stock_daily on talib_daily.ts_code = stock_daily.ts_code and talib_daily.trade_date = stock_daily.trade_date "
	sql = sql + " left join szzs on talib_daily.trade_date = szzs.trade_date "
	sql = sql + " left join stock_hua_basic on stock_hua_basic.ts_code_name = stock_daily.ts_code "
	sql = sql + " left join latest on stock_hua_basic.ts_code_name = latest.ts_code "

	sql = sql + " where stock_daily.trade_date > DATE_SUB(CURDATE(), INTERVAL 2 month)\n "

	if bands.TsCode != "" {
		sql = sql + " and talib_daily.ts_code = '" + bands.TsCode + "'"
	}
	if bands.Scope == 0.0 {
		bands.Scope = 1
	}
	sql = sql + " and stock_daily.`close` < " + fmt.Sprintf("%f", bands.Scope) + " * talib_daily.lower"
	sql = sql + " ) a where num = 1"
	sql = sql + " ORDER BY trade_date desc limit 1000;\n "
	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	return result, err
}

func (e *StockBuyServer) GetBandsNum(bands stock.Bands) (list []stock.Bands, err error) {
	var result []stock.Bands
	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select talib_daily.ts_code,name,industry,sum(if(close < talib_daily.lower, 1, 0)) as num "
	sql = sql + " from talib_daily \nleft join stock_daily USING(trade_date,ts_code) \n"
	sql = sql + " left join stock_hua_basic on talib_daily.ts_code = stock_hua_basic.ts_code_name\n"
	sql = sql + " where trade_date > '2023-09-01'\n"
	sql = sql + " GROUP BY talib_daily.ts_code having num >= 2 ORDER BY num desc;"
	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	return result, err
}

func (e *StockBuyServer) GetBandsList(bands stock.Bands) (list []stock.Bands, err error) {
	var result []stock.Bands
	db := global.MustGetGlobalDBByDBName("stock")
	sql := "select * from (select name, industry, talib_daily.ts_code,talib_daily.lower,"
	sql = sql + " stock_daily.close,talib_daily.trade_date,pct_chg,szzs.`close` as szzs_c,"
	sql = sql + " ROW_NUMBER() over (PARTITION by ts_code ORDER BY trade_date desc) as num"
	sql = sql + " from talib_daily join stock_daily on talib_daily.ts_code = stock_daily.ts_code and talib_daily.trade_date = stock_daily.trade_date "
	sql = sql + " left join szzs on talib_daily.trade_date = szzs.trade_date "
	sql = sql + " left join stock_hua_basic  on stock_hua_basic.ts_code_name = stock_daily.ts_code "

	sql = sql + " where stock_daily.trade_date > DATE_SUB(CURDATE(), INTERVAL 4 year)\n "
	tsCode := bands.TsCode
	if bands.TsCode == "" {
		tsCode = "000001.SZ"
	}
	sql = sql + " and talib_daily.ts_code = '" + tsCode + "'"

	if bands.Scope == 0.0 {
		bands.Scope = 1
	}
	sql = sql + " and stock_daily.`close` < " + fmt.Sprintf("%f", bands.Scope) + " * talib_daily.lower"
	sql = sql + " ) a"
	//sql = sql + " where num = 1"
	sql = sql + " ORDER BY trade_date desc limit 1000;\n "
	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	return result, err
}

func (e *StockBuyServer) GetBabAndDetails(bands stock.Bands) (list []stock.Bands, err error) {
	var result []stock.Bands
	db := global.MustGetGlobalDBByDBName("stock")

	sql := "select talib_daily.ts_code,stock_daily.trade_date,stock_daily.`close`,stock_daily.pct_chg,talib_daily.lower,szzs.`close` as szzs_c "
	sql = sql + " from stock_daily "
	sql = sql + " left join talib_daily on talib_daily.ts_code = stock_daily.ts_code and talib_daily.trade_date = stock_daily.trade_date\n"
	sql = sql + " left join szzs on stock_daily.trade_date = szzs.trade_date "
	if bands.TsCode == "" {
		bands.TsCode = "000001.SZ"
	}
	sql = sql + " where talib_daily.ts_code = '" + bands.TsCode + "'"
	sql = sql + " and stock_daily.trade_date > DATE_SUB(CURDATE(), INTERVAL 4 year)\n "
	sql = sql + "ORDER BY talib_daily.trade_date desc ;"

	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	return result, err
}
