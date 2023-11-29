package controller

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/stock"
	stoRes "github.com/flipped-aurora/gin-vue-admin/server/model/stock/response"
	stoUtil "github.com/flipped-aurora/gin-vue-admin/server/utils/stock"
	"strconv"
)

type StockController struct {
}

func (e *StockController) GetBuyDetail(summmary stoRes.BuySaleSummaryRes) (list interface{}, err error) {
	var result []stock.BuySaleDate
	val, _ := json.Marshal(summmary)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}

	result, _ = stockBuyServer.GetBuyDetail(summmary)
	origin := 1.0
	leiji := 0.0
	num := 0.0
	nn := 0
	for key, value := range result {
		if key == 0 {
			origin = value.Money
			result[key].Zhang = "-"
		} else {
			result[key].Zhang = fmt.Sprintf("%.1f", (value.Money-origin)/origin*100) + "%"
		}
		util := stoUtil.Tax{}
		if value.Type == "1" {
			leiji = leiji + util.CalcTax(value.Type, value.Vol)
			nn, _ = strconv.Atoi(value.Num)
			num = num + float64(nn)
		} else if value.Type == "2" {
			leiji = leiji - util.CalcTax(value.Type, value.Vol)
			nn, _ = strconv.Atoi(value.Num)
			num = num - float64(nn)
		}

		result[key].Leiji = fmt.Sprintf("%.2f", leiji)
		result[key].Pingjun = fmt.Sprintf("%.3f", leiji/num)
		low, _ := strconv.ParseFloat(value.Low, 64)
		high, _ := strconv.ParseFloat(value.High, 64)
		result[key].High = fmt.Sprintf("%.2f", high)
		result[key].Low = fmt.Sprintf("%.2f", low) +
			"/" + fmt.Sprintf("%.1f", (low-value.Money)*100/value.Money) + "%"

	}
	redisSetList(key, result)
	return result, nil
}

func (e *StockController) GetBuyList(code string) (list interface{}, err error) {
	var result []stock.BuySale
	val, _ := json.Marshal(code)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}

	result, _ = stockBuyServer.GetBuyList(code)
	redisSetList(key, result)
	return result, nil
}

func (e *StockController) GetSection(section stock.Section) (list interface{}, err error) {
	var result []stock.Section
	val, _ := json.Marshal(section)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}
	result, _ = stockBuyServer.GetSection(section.Tscode, section.Type)
	//redisSetList(key, result)
	return result, err
}

func (e *StockController) GetBands(bands stock.Bands) (list interface{}, err error) {
	var result []stock.Bands
	val, _ := json.Marshal(bands)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}
	if bands.Type == 2 {
		result, _ = stockBuyServer.GetBandsNum(bands)
	} else {
		result, _ = stockBuyServer.GetBands(bands)
	}
	redisSetList(key, result)
	return result, err
}

func (e *StockController) GetBandsList(bands stock.Bands) (list interface{}, err error) {
	var result []stock.Bands
	val, _ := json.Marshal(bands)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}
	result, _ = stockBuyServer.GetBandsList(bands)
	redisSetList(key, result)
	return result, err
}

func (e *StockController) GetBabAndDetails(bands stock.Bands) (list interface{}, err error) {
	var result []stock.Bands
	val, _ := json.Marshal(bands)
	key := getMethodName() + fmt.Sprintf("%s", val)
	errKR := redisGetList(key, &result)
	if errKR == nil {
		return result, nil
	}
	result, _ = stockBuyServer.GetBabAndDetails(bands)
	redisSetList(key, result)
	return result, err
}
