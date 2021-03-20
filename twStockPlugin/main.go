package main

import (
	"fmt"
	"time"
	"twStockPlugin/stockInfo"
	"twStockPlugin/stockInstitutionalInvestors"
	"twStockPlugin/stockTrading"
	"twStockPlugin/utils"
)

func main () {
	// testing date utils
	dateUtils := utils.DateUtils{}
	fmt.Println(dateUtils.AcToRepublicEra("20200318"))
	fmt.Println(dateUtils.AcToRepublicEra("2020/03/18"))
	fmt.Println(dateUtils.AcToRepublicEra("2020-03-18"))
	fmt.Println(dateUtils.RepublicEraToAc("1100318"))
	fmt.Println(dateUtils.RepublicEraToAc("110/03/18"))
	fmt.Println(dateUtils.RepublicEraToAc("110-03-18"))

	// testing fetch stock info
	stockInfo_ := stockInfo.Fetch()
	fmt.Println(stockInfo_["2330"])

	// testing fetch stock trading
	date_ := time.Date(2021, 3, 15, 0, 0, 0, 0, time.UTC)
	stockTrading_ := stockTrading.Fetch(date_)
	fmt.Println(stockTrading_["1101"])
	fmt.Println(stockTrading_["9962"])

	time.Sleep(3 * time.Second)
	// testing fetch stock institutional investors
	stockInstitutionalInvestors_ := stockInstitutionalInvestors.Fetch(date_)
	fmt.Println(stockInstitutionalInvestors_["2353"])
	fmt.Println(stockInstitutionalInvestors_["3540"])
}
