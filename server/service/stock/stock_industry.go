package stock

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	"strconv"
)

type StockIndustryServer struct{}

func (e *StockIndustryServer) GetGainianStock() (list interface{}, err error) {
	var db = global.MustGetGlobalDBByDBName("stock")
	var result []stock.StockBasic
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	db.Raw("SELECT * FROM `stock_hua_basic` where gainian like \"%新能源%\" \nand industry in (select industry from daily_industry)\nand name not like \"ST%\" ORDER BY industry asc;").Scan(&result)
	redisSetList(getMethodName()+key, result)
	return result, nil
}

func (e *StockIndustryServer) GetStockRank(basic stock.StockRank) (list interface{}, err error) {
	var db = global.MustGetGlobalDBByDBName("stock")
	var result []stock.StockRank
	val, _ := json.Marshal(basic)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}

	sql := "select industry, name,GROUP_CONCAt(rank1 ORDER BY trade_date desc) as rank1 from stock_rank"
	if basic.Industry != "" {
		sql = sql + " where industry = \"" + basic.Industry + "\""
	} else {
		sql = sql + " where industry = \"" + "证券" + "\""
	}

	if basic.Name != "" {
		sql = sql + " and name = \"" + basic.Name + "\""
	}
	if basic.TradeDate == "" {
		basic.TradeDate = "2023-08-01"
	} else {

	}
	sql = sql + " and trade_date > \"" + basic.TradeDate + "\""
	sql = sql + " GROUP BY industry, name ORDER BY industry,name limit 100"
	fmt.Println(sql)
	err = db.Raw(sql).Find(&result).Error
	if err != nil {

	}
	redisSetList(key, result)
	return result, nil
}

func (e *StockIndustryServer) GetIndustryRank() (list interface{}, err error) {
	var result []stock.IndustryRank
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	var db = global.MustGetGlobalDBByDBName("stock")
	sql := "select industry, GROUP_CONCAT(top ORDER BY trade_date desc) as rank from daily_industry GROUP BY industry;"

	db.Raw(sql).Scan(&result)
	redisSetList(getMethodName()+key, result)
	return result, nil
}

func (e *StockIndustryServer) GetIndustryLatest() (list interface{}, err error) {
	var result []stock.IndustryRank
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	var db = global.MustGetGlobalDBByDBName("stock")
	sql := "SELECT latest.ts_code,value,close,round((value-close) * 100/close,2) as today_fudu,"
	sql = sql + "round((close-b.money) * 100/b.money,2) as two_year_fudu,\n b.money,\n"
	sql = sql + " round((close/c.money) * 100,2) as two_year_great_fudu,\n\tc.money as two_year_great,\n"
	sql = sql + " name,industry,caiwufenxi,stock_hua_basic.fenlei,gainian "
	sql = sql + "\nFROM `latest` left join stock_hua_basic on latest.ts_code = stock_hua_basic.ts_code_name\nleft join (select * from stock_daily  where trade_date in (select max(trade_date) from stock_date)) a  using(ts_code)\nleft join (select * from section where type = 'm24' and order_by = 0) b using(ts_code)\nleft join (select * from section where type = 'm24' and order_by = 9) c using(ts_code)\n\n"
	sql = sql + "where industry = '证券';"

	db.Raw(sql).Scan(&result)
	//redisSetListTime(getMethodName()+key, result, time.Minute*2)
	return result, nil
}

func (e *StockIndustryServer) GetIndustryDetail(industry string) (list interface{}, err error) {
	var result []stock.IndustryDetail
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	var db = global.MustGetGlobalDBByDBName("stock")

	sql := "select * from daily_industry " + "where industry = '" + industry + "' ORDER BY trade_date desc;\n"
	db.Raw(sql).Scan(&result)
	redisSetList(getMethodName()+key, result)

	return result, nil
}

// 顶部行业
func (e *StockIndustryServer) GetIndustryTop(page request.PageInfo) (list interface{}, total int64, err error) {
	var result []stock.IndustryDaily
	val, _ := json.Marshal(page)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	totalStr, _ := redisGet(key + "count")
	totals, _ := strconv.Atoi(totalStr)
	if errKR == nil {
		total = int64(totals)
		return result, total, nil
	}
	var db = global.MustGetGlobalDBByDBName("stock")

	sql := "select trade_date,\nmax(case when `rank_top` = 'Top1' then industry end) as top1,\nmax(case when `rank_top` = 'Top2' then industry end) as top2,\nmax(case when `rank_top` = 'Top3' then industry end) as top3,\nmax(case when `rank_top` = 'Top4' then industry end) as top4,\nmax(case when `rank_top` = 'Top5' then industry end) as top5,\nmax(case when `rank_top` = 'Top6' then industry end) as top6,\nmax(case when `rank_top` = 'Top7' then industry end) as top7,\nmax(case when `rank_top` = 'Top8' then industry end) as top8,\nmax(case when `rank_top` = 'Top9' then industry end) as top9,\nmax(case when `rank_top` = 'Top10' then industry end) as top10,\nmax(case when `rank_top` = 'Top11' then industry end) as top11,\nmax(case when `rank_top` = 'Top12' then industry end) as top12,\nmax(case when `rank_top` = 'Top13' then industry end) as top13,\nmax(case when `rank_top` = 'Top14' then industry end) as top14,\nmax(case when `rank_top` = 'Top15' then industry end) as top15,\nmax(case when `rank_top` = 'Top16' then industry end) as top16,\nmax(case when `rank_top` = 'Top17' then industry end) as top17,\nmax(case when `rank_top` = 'Top18' then industry end) as top18,\nmax(case when `rank_top` = 'Top19' then industry end) as top19,\nmax(case when `rank_top` = 'Top20' then industry end) as top20,\nmax(case when `rank_top` = 'Top21' then industry end) as top21,\nmax(case when `rank_top` = 'Top22' then industry end) as top22,\nmax(case when `rank_top` = 'Top23' then industry end) as top23,\nmax(case when `rank_top` = 'Top24' then industry end) as top24,\nmax(case when `rank_top` = 'Top25' then industry end) as top25,\nmax(case when `rank_top` = 'Top26' then industry end) as top26,\nmax(case when `rank_top` = 'Top27' then industry end) as top27,\nmax(case when `rank_top` = 'Top28' then industry end) as top28,\nmax(case when `rank_top` = 'Top29' then industry end) as top29,\nmax(case when `rank_top` = 'Top30' then industry end) as top30,\nmax(case when `rank_top` = 'Top31' then industry end) as top31,\nmax(case when `rank_top` = 'Top32' then industry end) as top32,\nmax(case when `rank_top` = 'Top33' then industry end) as top33,\nmax(case when `rank_top` = 'Top34' then industry end) as top34,\nmax(case when `rank_top` = 'Top35' then industry end) as top35,\nmax(case when `rank_top` = 'Top36' then industry end) as top36,\nmax(case when `rank_top` = 'Top37' then industry end) as top37,\nmax(case when `rank_top` = 'Top38' then industry end) as top38,\nmax(case when `rank_top` = 'Top39' then industry end) as top39,\nmax(case when `rank_top` = 'Top40' then industry end) as top40,\nmax(case when `rank_top` = 'Top41' then industry end) as top41,\nmax(case when `rank_top` = 'Top42' then industry end) as top42,\nmax(case when `rank_top` = 'Top43' then industry end) as top43,\nmax(case when `rank_top` = 'Top44' then industry end) as top44,\nmax(case when `rank_top` = 'Top45' then industry end) as top45,\nmax(case when `rank_top` = 'Top46' then industry end) as top46,\nmax(case when `rank_top` = 'Top47' then industry end) as top47,\nmax(case when `rank_top` = 'Top48' then industry end) as top48,\nmax(case when `rank_top` = 'Top49' then industry end) as top49,\nmax(case when `rank_top` = 'Top50' then industry end) as top50\nfrom daily_industry \nGROUP BY trade_date\nORDER BY trade_date desc "
	sql = sql + " limit " + fmt.Sprintf("%d", page.PageSize) + " offset " + fmt.Sprintf("%d", (page.Page-1)*page.PageSize)
	countSql := "select count(1) from daily_industry GROUP BY trade_date;"
	var count int64
	db.Raw(countSql).Count(&count)
	db.Raw(sql).Scan(&result)
	redisSetList(key, result)
	redisSetList(key+"count", count)

	return result, count, nil
}

func (e *StockIndustryServer) GetStockRankMonth(industry string, year int, info request.PageInfo) (list interface{}, err error) {
	var result []stock.StockRankMonth

	indus, _ := json.Marshal(industry)
	infos, _ := json.Marshal(info)
	key := getMethodName() + fmt.Sprintf("%s", indus) + fmt.Sprintf("%s", infos) + fmt.Sprintf("%d", year)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}
	var db = global.MustGetGlobalDBByDBName("stock")

	sql := "SELECT industry, ts_code, max(name) as name, year_c,"
	sql = sql + "max(if(month_c = 1, rank1, '')) as month_1,"
	sql = sql + "max(if(month_c = 2, rank1, '')) as month_2,"
	sql = sql + "max(if(month_c = 3, rank1, '')) as month_3,"
	sql = sql + "max(if(month_c = 4, rank1, '')) as month_4,"
	sql = sql + "max(if(month_c = 5, rank1, '')) as month_5,"
	sql = sql + "max(if(month_c = 6, rank1, '')) as month_6,"
	sql = sql + "max(if(month_c = 7, rank1, '')) as month_7,"
	sql = sql + "max(if(month_c = 8, rank1, '')) as month_8,"
	sql = sql + "max(if(month_c = 9, rank1, '')) as month_9,"
	sql = sql + "max(if(month_c = 10, rank1, '')) as month_10,"
	sql = sql + "max(if(month_c = 11, rank1, '')) as month_11,"
	sql = sql + "max(if(month_c = 12, rank1, '')) as month_12"
	sql = sql + " FROM `stock_rank_month` where industry != \"\" "
	if industry != "" {
		sql = sql + " and industry = \"" + industry + "\""
	}

	if year > 2011 {
		sql = sql + " and year_c >= \"" + fmt.Sprintf("%d", year) + "\""
	}
	//if len(years) > 0 {
	//	sql = sql + "where industry in (" + years + ")"
	//}

	sql = sql + " GROUP BY industry, ts_code, year_c order by industry,ts_code,year_c "
	if industry == "" {
		offset := fmt.Sprintf("%d", ((info.Page - 1) * info.PageSize))
		sql = sql + " limit " + fmt.Sprintf("%d", info.PageSize) + " offset " + offset
	}
	fmt.Println(sql)

	db.Raw(sql).Scan(&result)
	redisSetList(key, result)

	return result, nil
}

func (e *StockIndustryServer) GetYearClose() (list interface{}, err error) {
	var db = global.MustGetGlobalDBByDBName("stock")
	var result []stock.YearClose
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	db.Raw("select * from  (select max(trade_date) as trade_date from stock_date) a \nleft join close_max_day on a.trade_date = close_max_day.trade_date\nleft join stock_basic using(ts_code)\n ORDER BY year_chg asc,ts_code limit 500;").Scan(&result)
	redisSetList(getMethodName()+key, result)
	return result, nil
}

func (e *StockIndustryServer) GetPctSum() (list interface{}, err error) {
	sql := "select * from \n(\nselect ts_code\n, sum(if(d_pct=60, 1,0)) as chg_60\n, sum(if(d_pct=30, 1,0)) as chg_30\n\n from (\nselect ts_code,pct_chg,60 as d_pct from stock_daily\nleft join stock_basic using(ts_code)\nwhere trade_date > '20230715' and abs(pct_chg) > 4\nunion all \nselect ts_code,pct_chg,30 as d_pct from stock_daily\nleft join stock_basic using(ts_code)\nwhere trade_date > '20230815' and abs(pct_chg) > 4) a\nGROUP BY ts_code) b\nleft join stock_basic on stock_basic.ts_code = b.ts_code\n"
	var db = global.MustGetGlobalDBByDBName("stock")
	var result []stock.PctSum
	key := ""
	errKR := redisGetList(getMethodName()+key, &result)
	if errKR == nil {
		return result, nil
	}
	db.Raw(sql).Scan(&result)
	redisSetList(getMethodName()+key, result)
	return result, nil
}
