package stockInstitutionalInvestors

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"goTwStock/constant"
	"goTwStock/schema"
	"goTwStock/utils"
)

var requestHandler = utils.RequestHandler{}

func fetchTwse(date_ time.Time, results map[string]*schema.StockInstitutionalInvestors) {
	/*
		上市三大法人資訊
	*/
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("%s/fund/T86", constant.TwseDomain)
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
		foreignDealersExcludeBuy := strings.ReplaceAll(data[2], ",", "")
		foreignDealersExcludeBuy_, _ := strconv.ParseFloat(foreignDealersExcludeBuy, 64)
		foreignDealersExcludeSell := strings.ReplaceAll(data[3], ",", "")
		foreignDealersExcludeSell_, _ := strconv.ParseFloat(foreignDealersExcludeSell, 64)
		foreignDealersBuy := strings.ReplaceAll(data[5], ",", "")
		foreignDealersBuy_, _ := strconv.ParseFloat(foreignDealersBuy, 64)
		foreignDealersSell := strings.ReplaceAll(data[6], ",", "")
		foreignDealersSell_, _ := strconv.ParseFloat(foreignDealersSell, 64)
		trustBuy := strings.ReplaceAll(data[8], ",", "")
		trustBuy_, _ := strconv.ParseFloat(trustBuy, 64)
		trustSell := strings.ReplaceAll(data[9], ",", "")
		trustSell_, _ := strconv.ParseFloat(trustSell, 64)
		dealersProprietaryBuy := strings.ReplaceAll(data[12], ",", "")
		dealersProprietaryBuy_, _ := strconv.ParseFloat(dealersProprietaryBuy, 64)
		dealersProprietarySell := strings.ReplaceAll(data[13], ",", "")
		dealersProprietarySell_, _ := strconv.ParseFloat(dealersProprietarySell, 64)
		dealersHedgeBuy := strings.ReplaceAll(data[15], ",", "")
		dealersHedgeBuy_, _ := strconv.ParseFloat(dealersHedgeBuy, 64)
		dealersHedgeSell := strings.ReplaceAll(data[16], ",", "")
		dealersHedgeSell_, _ := strconv.ParseFloat(dealersHedgeSell, 64)
		total := strings.ReplaceAll(data[18], ",", "")
		total_, _ := strconv.ParseFloat(total, 64)
		stockInstitutionalInvestors := &schema.StockInstitutionalInvestors{
			Code:                      data[0],
			ForeignDealersExcludeBuy:  foreignDealersExcludeBuy_,
			ForeignDealersExcludeSell: foreignDealersExcludeSell_,
			ForeignDealersBuy:         foreignDealersBuy_,
			ForeignDealersSell:        foreignDealersSell_,
			TrustBuy:                  trustBuy_,
			TrustSell:                 trustSell_,
			DealersProprietaryBuy:     dealersProprietaryBuy_,
			DealersProprietarySell:    dealersProprietarySell_,
			DealersHedgeBuy:           dealersHedgeBuy_,
			DealersHedgeSell:          dealersHedgeSell_,
			Total:                     total_,
		}
		results[stockInstitutionalInvestors.Code] = stockInstitutionalInvestors
	}
}

func fetchTpex(date_ time.Time, results map[string]*schema.StockInstitutionalInvestors) {
	/*
		上櫃三大法人資訊
	*/
	dateUtils := utils.DateUtils{}
	date__ := date_.Format("2006/01/02")
	date__ = dateUtils.AcToRepublicEra(date__)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	url := fmt.Sprintf("%s/web/stock/3insti/daily_trade/3itrade_hedge_result.php", constant.TpexDomain)
	params := map[string]string{
		"l":  "zh-tw",
		"se": "EW",
		"t":  "D",
		"d": date__,
		"_":  timestamp,
	}
	body := requestHandler.GetRequest(url, params)
	data := make(map[string][][]string)
	json.Unmarshal(body, &data)
	dataList := data["aaData"]
	for _, data := range dataList {
		foreignDealersExcludeBuy := strings.ReplaceAll(data[2], ",", "")
		foreignDealersExcludeBuy_, _ := strconv.ParseFloat(foreignDealersExcludeBuy, 64)
		foreignDealersExcludeSell := strings.ReplaceAll(data[3], ",", "")
		foreignDealersExcludeSell_, _ := strconv.ParseFloat(foreignDealersExcludeSell, 64)
		foreignDealersBuy := strings.ReplaceAll(data[5], ",", "")
		foreignDealersBuy_, _ := strconv.ParseFloat(foreignDealersBuy, 64)
		foreignDealersSell := strings.ReplaceAll(data[6], ",", "")
		foreignDealersSell_, _ := strconv.ParseFloat(foreignDealersSell, 64)
		trustBuy := strings.ReplaceAll(data[11], ",", "")
		trustBuy_, _ := strconv.ParseFloat(trustBuy, 64)
		trustSell := strings.ReplaceAll(data[12], ",", "")
		trustSell_, _ := strconv.ParseFloat(trustSell, 64)
		dealersProprietaryBuy := strings.ReplaceAll(data[14], ",", "")
		dealersProprietaryBuy_, _ := strconv.ParseFloat(dealersProprietaryBuy, 64)
		dealersProprietarySell := strings.ReplaceAll(data[15], ",", "")
		dealersProprietarySell_, _ := strconv.ParseFloat(dealersProprietarySell, 64)
		dealersHedgeBuy := strings.ReplaceAll(data[17], ",", "")
		dealersHedgeBuy_, _ := strconv.ParseFloat(dealersHedgeBuy, 64)
		dealersHedgeSell := strings.ReplaceAll(data[18], ",", "")
		dealersHedgeSell_, _ := strconv.ParseFloat(dealersHedgeSell, 64)
		total := strings.ReplaceAll(data[23], ",", "")
		total_, _ := strconv.ParseFloat(total, 64)
		stockInstitutionalInvestors := &schema.StockInstitutionalInvestors{
			Code:                      data[0],
			ForeignDealersExcludeBuy:  foreignDealersExcludeBuy_,
			ForeignDealersExcludeSell: foreignDealersExcludeSell_,
			ForeignDealersBuy:         foreignDealersBuy_,
			ForeignDealersSell:        foreignDealersSell_,
			TrustBuy:                  trustBuy_,
			TrustSell:                 trustSell_,
			DealersProprietaryBuy:     dealersProprietaryBuy_,
			DealersProprietarySell:    dealersProprietarySell_,
			DealersHedgeBuy:           dealersHedgeBuy_,
			DealersHedgeSell:          dealersHedgeSell_,
			Total:                     total_,
		}
		results[stockInstitutionalInvestors.Code] = stockInstitutionalInvestors
	}
}

func Fetch(date_ time.Time) map[string]*schema.StockInstitutionalInvestors{
	results := map[string]*schema.StockInstitutionalInvestors{}
	fetchTwse(date_, results)
	fetchTpex(date_, results)
	return results
}

