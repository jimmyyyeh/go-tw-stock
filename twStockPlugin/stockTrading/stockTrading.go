package stockTrading

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"goTwStock/constant"
	"goTwStock/regex"
	"goTwStock/schema"
	"goTwStock/utils"
)

var requestHandler = utils.RequestHandler{}

func fetchTwse(date_ time.Time, results map[string]*schema.StockTrading) {
	/*
		上市收盤資訊
	*/
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("%s/exchangeReport/MI_INDEX", constant.TwseDomain)
	params := map[string]string{
		"response": "json",
		"date":     date_.Format("20060102"),
		"type":     "ALL",
		"_":        timestamp,
	}
	body := requestHandler.GetRequest(url, params)
	data := make(map[string][][]string)
	json.Unmarshal(body, &data)
	dataList := data["data9"]
	for _, data := range dataList {
		tradingVolume := strings.ReplaceAll(data[2], ",", "")
		tradingVolume_, _ := strconv.ParseFloat(tradingVolume, 64)
		transaction := strings.ReplaceAll(data[3], ",", "")
		transaction_, _ := strconv.ParseFloat(transaction, 64)
		tradeValue := strings.ReplaceAll(data[4], ",", "")
		tradeValue_, _ := strconv.ParseFloat(tradeValue, 64)
		open := strings.ReplaceAll(data[5], ",", "")
		open_, _ := strconv.ParseFloat(open, 64)
		high := strings.ReplaceAll(data[6], ",", "")
		high_, _ := strconv.ParseFloat(high, 64)
		low := strings.ReplaceAll(data[7], ",", "")
		low_, _ := strconv.ParseFloat(low, 64)
		close_ := strings.ReplaceAll(data[8], ",", "")
		close__, _ := strconv.ParseFloat(close_, 64)
		changeSymbol := regex.ChangeSymbol.FindString(data[9])
		changeValue := fmt.Sprintf("%s%s", changeSymbol, data[10])
		changeValue_, _ := strconv.ParseFloat(changeValue, 64)
		stockTrading := &schema.StockTrading{
			Code:        data[0],
			TradeVolume: tradingVolume_,
			Transaction: transaction_,
			TradeValue:  tradeValue_,
			Open:        open_,
			High:        high_,
			Low:         low_,
			Close:       close__,
			Change:      changeValue_,
		}
		results[stockTrading.Code] = stockTrading
	}
}

func fetchTpex(date_ time.Time, results map[string]*schema.StockTrading) {
	/*
		上櫃收盤資訊
	*/
	dateUtils := utils.DateUtils{}
	date__ := date_.Format("2006/01/02")
	date__ = dateUtils.AcToRepublicEra(date__)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("%s%s", constant.TpexDomain, "/web/stock/aftertrading/otc_quotes_no1430/stk_wn1430_result.php")
	params := map[string]string{
		"l": "zh-tw",
		"se" :"AL",
		"d": date__,
		"_": timestamp,
	}
	body := requestHandler.GetRequest(url, params)
	data := make(map[string][][]string)
	json.Unmarshal(body, &data)
	dataList := data["aaData"]
	for _, data := range dataList {
		tradingVolume := strings.ReplaceAll(data[7], ",", "")
		tradingVolume_, _ := strconv.ParseFloat(tradingVolume, 64)
		transaction := strings.ReplaceAll(data[9], ",", "")
		transaction_, _ := strconv.ParseFloat(transaction, 64)
		tradeValue := strings.ReplaceAll(data[8], ",", "")
		tradeValue_, _ := strconv.ParseFloat(tradeValue, 64)
		open := strings.ReplaceAll(data[4], ",", "")
		open_, _ := strconv.ParseFloat(open, 64)
		high := strings.ReplaceAll(data[5], ",", "")
		high_, _ := strconv.ParseFloat(high, 64)
		low := strings.ReplaceAll(data[6], ",", "")
		low_, _ := strconv.ParseFloat(low, 64)
		close_ := strings.ReplaceAll(data[2], ",", "")
		close__, _ := strconv.ParseFloat(close_, 64)
		changeValue := strings.ReplaceAll(data[3], ",", "")
		changeValue_, _ := strconv.ParseFloat(changeValue, 64)
		stockTrading := &schema.StockTrading{
			Code:        data[0],
			TradeVolume: tradingVolume_,
			Transaction: transaction_,
			TradeValue:  tradeValue_,
			Open:        open_,
			High:        high_,
			Low:         low_,
			Close:       close__,
			Change:      changeValue_,
		}
		results[stockTrading.Code] = stockTrading
	}
}

func Fetch(date_ time.Time) map[string]*schema.StockTrading{
	results := map[string]*schema.StockTrading{}
	fetchTwse(date_, results)
	fetchTpex(date_, results)
	return results
}
