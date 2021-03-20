package stockMarginTrading

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"twStockPlugin/constant"
	"twStockPlugin/schema"
	"twStockPlugin/utils"
)

var requestHandler = utils.RequestHandler{}

func fetchTwse(date_ time.Time, results map[string]*schema.StockMarginTrading) {
	/*
		上市融資融券餘額資訊
	*/
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("%s/exchangeReport/MI_MARGN", constant.TwseDomain)
	params := map[string]string{
		"response":   "json",
		"date":       date_.Format("20060102"),
		"selectType": "ALL",
		"_":          timestamp,
	}
	body := requestHandler.GetRequest(url, params)
	data := make(map[string][][]string)
	json.Unmarshal(body, &data)
	dataList := data["data"]
	for _, data := range dataList {
		marginPurchase := strings.ReplaceAll(data[2], ",", "")
		marginPurchase_, _ := strconv.ParseFloat(marginPurchase, 64)
		marginSale := strings.ReplaceAll(data[3], ",", "")
		marginSale_, _ := strconv.ParseFloat(marginSale, 64)
		cashRedemption := strings.ReplaceAll(data[4], ",", "")
		cashRedemption_, _ := strconv.ParseFloat(cashRedemption, 64)
		marginBalanceOfPreviousDay := strings.ReplaceAll(data[5], ",", "")
		marginBalanceOfPreviousDay_, _ := strconv.ParseFloat(marginBalanceOfPreviousDay, 64)
		marginBalanceOfTheDay := strings.ReplaceAll(data[6], ",", "")
		marginBalanceOfTheDay_, _ := strconv.ParseFloat(marginBalanceOfTheDay, 64)
		marginQuota := strings.ReplaceAll(data[7], ",", "")
		marginQuota_, _ := strconv.ParseFloat(marginQuota, 64)

		shortCovering := strings.ReplaceAll(data[8], ",", "")
		shortCovering_, _ := strconv.ParseFloat(shortCovering, 64)
		shortSale := strings.ReplaceAll(data[9], ",", "")
		shortSale_, _ := strconv.ParseFloat(shortSale, 64)
		stockRedemption := strings.ReplaceAll(data[10], ",", "")
		stockRedemption_, _ := strconv.ParseFloat(stockRedemption, 64)
		shortBalanceOfPreviousDay := strings.ReplaceAll(data[11], ",", "")
		shortBalanceOfPreviousDay_, _ := strconv.ParseFloat(shortBalanceOfPreviousDay, 64)
		shortBalanceOfTheDay := strings.ReplaceAll(data[12], ",", "")
		shortBalanceOfTheDay_, _ := strconv.ParseFloat(shortBalanceOfTheDay, 64)
		shortQuota := strings.ReplaceAll(data[13], ",", "")
		shortQuota_, _ := strconv.ParseFloat(shortQuota, 64)
		offSetting := strings.ReplaceAll(data[14], ",", "")
		offSetting_, _ := strconv.ParseFloat(offSetting, 64)

		stockMarginTrading := &schema.StockMarginTrading{
			Code:                       data[0],
			MarginPurchase:             marginPurchase_,
			MarginSale:                 marginSale_,
			CashRedemption:             cashRedemption_,
			MarginBalanceOfPreviousDay: marginBalanceOfPreviousDay_,
			MarginBalanceOfTheDay:      marginBalanceOfTheDay_,
			MarginQuota:                marginQuota_,
			ShortCovering:              shortCovering_,
			ShortSale:                  shortSale_,
			StockRedemption:            stockRedemption_,
			ShortBalanceOfPreviousDay:  shortBalanceOfPreviousDay_,
			ShortBalanceOfTheDay:       shortBalanceOfTheDay_,
			ShortQuota:                 shortQuota_,
			OffSetting:                 offSetting_,
			Note:                       data[15],
		}
		results[stockMarginTrading.Code] = stockMarginTrading
	}
}

func fetchTpex(date_ time.Time, results map[string]*schema.StockMarginTrading) {
	/*
		上櫃融資融券餘額資訊
	*/
	dateUtils := utils.DateUtils{}
	date__ := date_.Format("2006/01/02")
	date__ = dateUtils.AcToRepublicEra(date__)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("%s/web/stock/margin_trading/margin_balance/margin_bal_result.php", constant.TpexDomain)
	params := map[string]string{
		"l": "zh-tw",
		"o": "json",
		"d": date__,
		"_": timestamp,
	}

	body := requestHandler.GetRequest(url, params)
	data := make(map[string][][]string)
	json.Unmarshal(body, &data)
	dataList := data["aaData"]
	for _, data := range dataList {
		marginPurchase := strings.ReplaceAll(data[3], ",", "")
		marginPurchase_, _ := strconv.ParseFloat(marginPurchase, 64)
		marginSale := strings.ReplaceAll(data[4], ",", "")
		marginSale_, _ := strconv.ParseFloat(marginSale, 64)
		cashRedemption := strings.ReplaceAll(data[5], ",", "")
		cashRedemption_, _ := strconv.ParseFloat(cashRedemption, 64)
		marginBalanceOfPreviousDay := strings.ReplaceAll(data[2], ",", "")
		marginBalanceOfPreviousDay_, _ := strconv.ParseFloat(marginBalanceOfPreviousDay, 64)
		marginBalanceOfTheDay := strings.ReplaceAll(data[6], ",", "")
		marginBalanceOfTheDay_, _ := strconv.ParseFloat(marginBalanceOfTheDay, 64)
		marginQuota := strings.ReplaceAll(data[9], ",", "")
		marginQuota_, _ := strconv.ParseFloat(marginQuota, 64)

		shortCovering := strings.ReplaceAll(data[11], ",", "")
		shortCovering_, _ := strconv.ParseFloat(shortCovering, 64)
		shortSale := strings.ReplaceAll(data[12], ",", "")
		shortSale_, _ := strconv.ParseFloat(shortSale, 64)
		stockRedemption := strings.ReplaceAll(data[13], ",", "")
		stockRedemption_, _ := strconv.ParseFloat(stockRedemption, 64)
		shortBalanceOfPreviousDay := strings.ReplaceAll(data[10], ",", "")
		shortBalanceOfPreviousDay_, _ := strconv.ParseFloat(shortBalanceOfPreviousDay, 64)
		shortBalanceOfTheDay := strings.ReplaceAll(data[14], ",", "")
		shortBalanceOfTheDay_, _ := strconv.ParseFloat(shortBalanceOfTheDay, 64)
		shortQuota := strings.ReplaceAll(data[17], ",", "")
		shortQuota_, _ := strconv.ParseFloat(shortQuota, 64)
		offSetting := strings.ReplaceAll(data[18], ",", "")
		offSetting_, _ := strconv.ParseFloat(offSetting, 64)

		stockMarginTrading := &schema.StockMarginTrading{
			Code:                       data[0],
			MarginPurchase:             marginPurchase_,
			MarginSale:                 marginSale_,
			CashRedemption:             cashRedemption_,
			MarginBalanceOfPreviousDay: marginBalanceOfPreviousDay_,
			MarginBalanceOfTheDay:      marginBalanceOfTheDay_,
			MarginQuota:                marginQuota_,
			ShortCovering:              shortCovering_,
			ShortSale:                  shortSale_,
			StockRedemption:            stockRedemption_,
			ShortBalanceOfPreviousDay:  shortBalanceOfPreviousDay_,
			ShortBalanceOfTheDay:       shortBalanceOfTheDay_,
			ShortQuota:                 shortQuota_,
			OffSetting:                 offSetting_,
			Note:                       data[19],
		}
		results[stockMarginTrading.Code] = stockMarginTrading
	}
}

func Fetch(date_ time.Time) map[string]*schema.StockMarginTrading {
	results := map[string]*schema.StockMarginTrading{}
	fetchTwse(date_, results)
	fetchTpex(date_, results)
	return results
}
