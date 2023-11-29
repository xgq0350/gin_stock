package stock

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stockReq "github.com/flipped-aurora/gin-vue-admin/server/model/stock/request"
	stockRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"strconv"
)

type StockBasicServer struct{}

func (exa *StockBasicServer) GetStockBasicList(sysUserAuthorityID uint, info request.PageInfo, stockBasic stockReq.SearchStoBasic) (list interface{}, total int64, err error) {

	limit := info.PageSize
	basic := stockBasic.StockBasic
	offset := info.PageSize * (info.Page - 1)

	db := global.MustGetGlobalDBByDBName("stock").Model(&stock.StockBasic{})
	var a system.SysAuthority
	a.AuthorityId = sysUserAuthorityID
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataId []uint
	for _, v := range auth.DataAuthorityId {
		dataId = append(dataId, v.AuthorityId)
	}

	var stockList []stock.StockBasic

	//拼装查询数据
	ddb := db.Joins("left join stock_hua_basic on stock_basic.symbol = stock_hua_basic.symbol")

	ddb = ddb.Joins("left join (select profit_percent,ts_code as code,trade_date,lowerest_price,latest_price from zlshenqi join (select max(trade_date) as trade_date from stock_date ) a using(trade_date)\nleft join stock_basic using(ts_code)) a on stock_basic.ts_code = a.code")
	ddb = ddb.Joins("left join stock_daily on stock_basic.ts_code = stock_daily.ts_code and a.trade_date = stock_daily.trade_date ")
	if stockBasic.Industry != "" {
		ddb = ddb.Where("stock_basic.industry = ?", stockBasic.Industry)
	}
	if stockBasic.Name != "" {
		ddb = ddb.Where("stock_basic.name = ?", stockBasic.Name)
	}
	if stockBasic.Symbol != "" {
		ddb = ddb.Where("stock_basic.symbol = ?", stockBasic.Symbol)
	}

	if stockBasic.TsCode != "" {
		ddb = ddb.Where("stock_basic.ts_code = ?", stockBasic.TsCode)
	}
	if basic.Caiwufenxi != "" {
		ddb = ddb.Where("stock_hua_basic.caiwufenxi = ?", basic.Caiwufenxi)
	}
	if basic.Gainian != "" {
		ddb = ddb.Where("stock_hua_basic.gainian like ?", "%"+basic.Gainian+"%")
	}
	if basic.Fenlei != "" {
		ddb = ddb.Where("stock_hua_basic.fenlei = ? ", basic.Fenlei)
	}
	market := []string{"主板", "中小板"}
	ddb = ddb.Where("market in (?)", market).Where("list_status = ?", "L")

	sqlSelect := "stock_basic.id, stock_basic.name,stock_basic.ts_code,stock_basic.area," +
		"stock_basic.industry,stock_basic.fullname,stock_basic.symbol,shiyingjing,shiyingdong," +
		"stock_hua_basic.liangdian,zhuying,shenwan,caiwufenxi,gainian,fenlei," +
		"a.lowerest_price,a.profit_percent,stock_daily.close, latest_price, " +
		"round((latest_price-lowerest_price)/lowerest_price * 100,1) as lat_low"

	if stockBasic.OrderKey != "" {
		order := stockBasic.OrderKey
		if stockBasic.Desc == false {
			order = order + " desc"
		} else {

		}
		ddb = ddb.Order(order)
	}
	valBasic, _ := json.Marshal(stockBasic)
	valPage, _ := json.Marshal(info)

	var key = getMethodName() + fmt.Sprintf("%s", valPage) + fmt.Sprintf("%s", valBasic)
	var keyC = key + "count"
	errl := ddb.Count(&total).Error

	errKR := redisGetList(key, &stockList)
	totalStr, errCR := redisGet(keyC)
	if errKR == nil && errCR == nil {
		totals, _ := strconv.Atoi(totalStr)
		total = int64(totals)
		return stockList, total, nil
	}

	if errl != nil {
		return stockList, total, err
	} else {
		ddb = ddb.Select(sqlSelect)
		err = ddb.Limit(limit).Offset(offset).Find(&stockList).Error
	}
	redisSetList(key, stockList)
	redisSetTotal(keyC, int(total))
	return stockList, total, err

}

func (exa *StockBasicServer) GetStockIndustry() (list interface{}, err error) {
	db := global.MustGetGlobalDBByDBName("stock").Model(&stock.StockBasic{})
	var stockList []stock.StockBasic

	key := getMethodName()
	errKR := redisGetList(key, &stockList)
	if errKR == nil {
		return stockList, nil
	}

	err = db.Group("industry").Select("industry").Find(&stockList).Error
	if err != nil {
		return stockList, err
	}
	redisSetList(key, stockList)
	return stockList, err
}

func (exa *StockBasicServer) GetMingci() (list interface{}, err error) {
	db := global.MustGetGlobalDBByDBName("stock").Model(&stock.StockDingyi{})
	var dingyiList []stock.StockDingyi
	key := getMethodName()
	errKR := redisGetList(key, &dingyiList)
	if errKR == nil {
		return dingyiList, nil
	}

	err = db.Find(&dingyiList).Error
	if err != nil {
		return dingyiList, err
	}
	redisSetList(key, dingyiList)
	return dingyiList, err
}

func (exa *StockBasicServer) GetZhiShuCode(pageOrder stockReq.PageOrder) (list interface{}, err error) {
	db := global.MustGetGlobalDBByDBName("stock")
	var zhishuList []stockRes.ZhishuMaxClose
	var key = ""
	if pageOrder.OrderKey != "" {
		key = pageOrder.OrderKey + strconv.FormatBool(pageOrder.Desc)
	}
	key = getMethodName() + key
	errKR := redisGetList(key, &zhishuList)
	if errKR == nil {
		return zhishuList, nil
	}
	sql := "select * from (SELECT code_name, code_name as ts_code, zhishu.name as name,zidian.name as zs_name,industry,caiwufenxi FROM `zhishu` \nleft join zidian using(type) \nleft join stock_hua_basic on stock_hua_basic.symbol = left(zhishu.code_name,6)) a "
	sql = sql + " left join (select ts_code,close_max_2023,close_max_2022,close_max_2021," +
		"close_min_2023,close_min_2022,close_min_2021 from close_max_min) b on a.ts_code = b.ts_code "
	sql = sql + " left join (select close,ts_code,pre_close,pct_chg from " +
		"(select max(trade_date) as trade_date from stock_date) c_t " +
		"left join stock_daily USING(trade_date)) c on a.ts_code = c.ts_code"

	if pageOrder.OrderKey != "" {
		fmt.Println(pageOrder.OrderKey)
		order := pageOrder.OrderKey
		if pageOrder.Desc == false {
			order = order + " desc"
		} else {

		}
		sql = sql + " order by " + order
	}
	fmt.Println(sql)
	db = db.Raw(sql)
	err = db.Find(&zhishuList).Error
	if err != nil {
		return zhishuList, err
	}
	redisSetList(key, zhishuList)

	return zhishuList, err
}

func (exa *StockBasicServer) GetLatest(basic stock.Basic) (info interface{}, err error) {
	db := global.MustGetGlobalDBByDBName("stock")
	var key = getMethodName() + basic.TsCode
	if basic.TsCode == "" {
		basic.TsCode = "000001.SZ"
	}
	var basicInfo stock.BasicInfo

	errKR := redisGetList(key, &basicInfo)
	if errKR == nil {
		return basicInfo, nil
	}

	sql := "select * from stock_hua_basic\nleft join \n(SELECT *,ROW_NUMBER() over(ORDER BY trade_date desc) as num FROM `close_max_day` where ts_code = "
	sql = sql + "'" + basic.TsCode + "' "
	sql = sql + ") a on stock_hua_basic.symbol = left(a.ts_code,6)\nleft join close_max_min b on stock_hua_basic.symbol = left(b.ts_code,6)\nwhere num =1 and b.ts_code = "
	sql = sql + "\"" + basic.TsCode + "\";\n\n"
	err = db.Raw(sql).Find(&basicInfo).Error
	if err != nil {
		return info, err
	}
	redisSetList(key, basicInfo)

	return basicInfo, err
}
