package main

import (
	"fmt"
	"time"
	"twStockPlugin/stockInfo"
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
	fmt.Println(stockInfo_["1101"])

	// testing fetch stock trading
	date_ := time.Date(2021, 3, 15, 0, 0, 0, 0, time.UTC)
	stockTrading_ := stockTrading.Fetch(date_)
	fmt.Println(stockTrading_["1101"])
	fmt.Println(stockTrading_["9962"])
}
